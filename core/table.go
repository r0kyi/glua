package core

import (
	"fmt"
	"math"
	"reflect"

	lua "github.com/yuin/gopher-lua"
)

func LTableToStrut(tbl *lua.LTable, out interface{}) error {
	v := reflect.ValueOf(out)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("out must be a non-nil pointer to struct")
	}

	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("out must be a pointer to struct")
	}

	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("lua")
		if tag == "" {
			continue
		}

		lv := tbl.RawGetString(tag)
		if lv == lua.LNil {
			continue
		}

		fv := v.Field(i)
		if !fv.CanSet() {
			continue
		}

		if err := setValue(lv, fv); err != nil {
			return fmt.Errorf("field %s: %w", field.Name, err)
		}
	}

	return nil
}

func setValue(lv lua.LValue, fv reflect.Value) error {
	switch fv.Kind() {
	case reflect.String:
		if str, ok := lv.(lua.LString); ok {
			fv.SetString(string(str))
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if num, ok := lv.(lua.LNumber); ok {
			fv.SetInt(int64(num))
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if num, ok := lv.(lua.LNumber); ok {
			fv.SetUint(uint64(num))
		}
	case reflect.Float32, reflect.Float64:
		if num, ok := lv.(lua.LNumber); ok {
			fv.SetFloat(float64(num))
		}
	case reflect.Bool:
		if b, ok := lv.(lua.LBool); ok {
			fv.SetBool(bool(b))
		}
	case reflect.Struct:
		if tbl, ok := lv.(*lua.LTable); ok {
			return LTableToStrut(tbl, fv.Addr().Interface())
		}
	case reflect.Slice:
		if tbl, ok := lv.(*lua.LTable); ok {
			elemType := fv.Type().Elem()
			slice := reflect.MakeSlice(fv.Type(), 0, tbl.Len())

			index := lua.LNil
			for {
				nextIndex, val := tbl.Next(index)
				if nextIndex == lua.LNil {
					break
				}

				elem := reflect.New(elemType).Elem()
				if err := setValue(val, elem); err != nil {
					return err
				}
				slice = reflect.Append(slice, elem)
				index = nextIndex
			}

			fv.Set(slice)
			return nil
		}
		return fmt.Errorf("expected table for slice, got %T", lv)

	case reflect.Map:
		if tbl, ok := lv.(*lua.LTable); ok {
			if fv.IsNil() {
				fv.Set(reflect.MakeMap(fv.Type()))
			}

			keyType := fv.Type().Key()
			elemType := fv.Type().Elem()

			key := lua.LNil
			for {
				nextKey, nextVal := tbl.Next(key)
				if nextKey == lua.LNil {
					break
				}

				goKey := reflect.New(keyType).Elem()
				if err := setValue(nextKey, goKey); err != nil {
					return err
				}

				goVal := reflect.New(elemType).Elem()

				if elemType.Kind() == reflect.Slice && elemType.Elem().Kind() == reflect.String {
					switch v := nextVal.(type) {
					case lua.LString:
						goVal.Set(reflect.ValueOf([]string{string(v)}))
					case *lua.LTable:
						slice := reflect.MakeSlice(elemType, 0, v.Len())
						idx := lua.LNil
						for {
							ni, val := v.Next(idx)
							if ni == lua.LNil {
								break
							}
							if s, ok := val.(lua.LString); ok {
								slice = reflect.Append(slice, reflect.ValueOf(string(s)))
							}
							idx = ni
						}
						goVal.Set(slice)
					default:
						return fmt.Errorf("expected string or table for map[string][]string, got %T", nextVal)
					}
				} else {
					if err := setValue(nextVal, goVal); err != nil {
						return err
					}
				}

				fv.SetMapIndex(goKey, goVal)
				key = nextKey
			}
			return nil
		}
		return fmt.Errorf("expected table for map, got %T", lv)
	default:
		return fmt.Errorf("unsupported kind: %s", fv.Kind())
	}

	return nil
}

func LTableToMap(tbl *lua.LTable) map[string]any {
	result := make(map[string]any)
	tbl.ForEach(func(key lua.LValue, value lua.LValue) {
		k := key.String()
		result[k] = lValueToGo(value)
	})

	return result
}

func lValueToGo(val lua.LValue) any {
	switch v := val.(type) {
	case lua.LBool:
		return bool(v)
	case lua.LNumber:
		return float64(v)
	case lua.LString:
		return string(v)
	case *lua.LTable:
		arr := make([]interface{}, 0)
		m := make(map[string]interface{})
		isArray := true

		v.ForEach(func(key, value lua.LValue) {
			if isArray {
				if key.Type() != lua.LTNumber {
					isArray = false
				} else {
					num := float64(key.(lua.LNumber))
					if num != math.Floor(num) || int(num) != len(arr)+1 {
						isArray = false
					}
				}
			}

			if !isArray {
				m[key.String()] = lValueToGo(value)
			} else {
				arr = append(arr, lValueToGo(value))
			}
		})

		if isArray {
			return arr
		}
		return m
	case *lua.LFunction:
		return v.String()
	case *lua.LUserData:
		return v.Value
	default:
		return v.String()
	}
}
