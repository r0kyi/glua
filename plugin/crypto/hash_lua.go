package crypto

import lua "github.com/yuin/gopher-lua"

func (h *Hash) String() string {
	return "glua.crypto.hash"
}

func (h *Hash) AssertFunction() lua.LGFunction {
	return nil
}

func (h *Hash) MetatableName() string {
	return "lua.table.hash"
}

func (h *Hash) md4L(L *lua.LState) int {
	if L.GetTop() != 1 {
		L.Push(lua.LNil)
		return 1
	}

	h.plaintext = L.CheckString(1)
	h.md4()
	L.Push(lua.LString(h.plaintext))

	return 1
}

func (h *Hash) md5L(L *lua.LState) int {
	if L.GetTop() != 1 {
		L.Push(lua.LNil)
		return 1
	}

	h.plaintext = L.CheckString(1)
	h.md5()
	L.Push(lua.LString(h.plaintext))

	return 1
}

func (h *Hash) ripemd160L(L *lua.LState) int {
	if L.GetTop() != 1 {
		L.Push(lua.LNil)
		return 1
	}

	h.plaintext = L.CheckString(1)
	h.ripemd160()
	L.Push(lua.LString(h.plaintext))

	return 1
}

func (h *Hash) sha1L(L *lua.LState) int {
	if L.GetTop() != 1 {
		L.Push(lua.LNil)
		return 1
	}

	h.plaintext = L.CheckString(1)
	h.sha1()
	L.Push(lua.LString(h.plaintext))

	return 1
}

func (h *Hash) sha3224L(L *lua.LState) int {
	if L.GetTop() != 1 {
		L.Push(lua.LNil)
		return 1
	}

	h.plaintext = L.CheckString(1)
	h.sha3224()
	L.Push(lua.LString(h.plaintext))

	return 1
}

func (h *Hash) sha3256L(L *lua.LState) int {
	if L.GetTop() != 1 {
		L.Push(lua.LNil)
		return 1
	}

	h.plaintext = L.CheckString(1)
	h.sha3256()
	L.Push(lua.LString(h.plaintext))

	return 1
}

func (h *Hash) sha3384L(L *lua.LState) int {
	if L.GetTop() != 1 {
		L.Push(lua.LNil)
		return 1
	}

	h.plaintext = L.CheckString(1)
	h.sha3384()
	L.Push(lua.LString(h.plaintext))

	return 1
}

func (h *Hash) sha3512L(L *lua.LState) int {
	if L.GetTop() != 1 {
		L.Push(lua.LNil)
		return 1
	}

	h.plaintext = L.CheckString(1)
	h.sha3512()
	L.Push(lua.LString(h.plaintext))

	return 1
}

func (h *Hash) sha224L(L *lua.LState) int {
	if L.GetTop() != 1 {
		L.Push(lua.LNil)
		return 1
	}

	h.plaintext = L.CheckString(1)
	h.sha224()
	L.Push(lua.LString(h.plaintext))

	return 1
}

func (h *Hash) sha256L(L *lua.LState) int {
	if L.GetTop() != 1 {
		L.Push(lua.LNil)
		return 1
	}

	h.plaintext = L.CheckString(1)
	h.sha256()
	L.Push(lua.LString(h.plaintext))

	return 1
}

func (h *Hash) sha384L(L *lua.LState) int {
	if L.GetTop() != 1 {
		L.Push(lua.LNil)
		return 1
	}

	h.plaintext = L.CheckString(1)
	h.sha384()
	L.Push(lua.LString(h.plaintext))

	return 1
}

func (h *Hash) sha512L(L *lua.LState) int {
	if L.GetTop() != 1 {
		L.Push(lua.LNil)
		return 1
	}

	h.plaintext = L.CheckString(1)
	h.sha512()
	L.Push(lua.LString(h.plaintext))

	return 1
}

func (h *Hash) sha512224L(L *lua.LState) int {
	if L.GetTop() != 1 {
		L.Push(lua.LNil)
		return 1
	}

	h.plaintext = L.CheckString(1)
	h.sha512224()
	L.Push(lua.LString(h.plaintext))

	return 1
}

func (h *Hash) sha512256L(L *lua.LState) int {
	if L.GetTop() != 1 {
		L.Push(lua.LNil)
		return 1
	}

	h.plaintext = L.CheckString(1)
	h.sha512256()
	L.Push(lua.LString(h.plaintext))

	return 1
}

func (h *Hash) blake2s128L(L *lua.LState) int {
	if L.GetTop() != 2 {
		L.Push(lua.LNil)
		return 1
	}

	h.plaintext = L.CheckString(1)
	h.key = L.CheckString(2)

	err := h.blake2s128()
	if err != nil {
		L.Push(lua.LString(err.Error()))
		return 1
	}

	L.Push(lua.LString(h.plaintext))

	return 1
}

func (h *Hash) blake2s256L(L *lua.LState) int {
	if L.GetTop() == 1 {
		h.plaintext = L.CheckString(1)

		err := h.blake2s256()
		if err != nil {
			L.Push(lua.LString(err.Error()))
			return 1
		}

		L.Push(lua.LString(h.plaintext))

		return 1
	} else if L.GetTop() == 2 {
		h.plaintext = L.CheckString(1)
		h.key = L.CheckString(2)

		err := h.blake2s256()
		if err != nil {
			L.Push(lua.LString(err.Error()))
			return 1
		}

		L.Push(lua.LString(h.plaintext))
		return 1
	} else {
		L.Push(lua.LNil)
		return 1
	}
}

func (h *Hash) blake2b256L(L *lua.LState) int {
	if L.GetTop() == 1 {
		h.plaintext = L.CheckString(1)

		err := h.blake2b256()
		if err != nil {
			L.Push(lua.LString(err.Error()))
			return 1
		}

		L.Push(lua.LString(h.plaintext))

		return 1
	} else if L.GetTop() == 2 {
		h.plaintext = L.CheckString(1)
		h.key = L.CheckString(2)

		err := h.blake2b256()
		if err != nil {
			L.Push(lua.LString(err.Error()))
			return 1
		}

		L.Push(lua.LString(h.plaintext))

		return 1
	} else {
		L.Push(lua.LNil)
		return 1
	}
}

func (h *Hash) blake2b384L(L *lua.LState) int {
	if L.GetTop() == 1 {
		h.plaintext = L.CheckString(1)

		err := h.blake2b384()
		if err != nil {
			L.Push(lua.LString(err.Error()))
			return 1
		}

		L.Push(lua.LString(h.plaintext))

		return 1
	} else if L.GetTop() == 2 {
		h.plaintext = L.CheckString(1)
		h.key = L.CheckString(2)

		err := h.blake2b384()
		if err != nil {
			L.Push(lua.LString(err.Error()))
			return 1
		}

		L.Push(lua.LString(h.plaintext))

		return 1
	} else {
		L.Push(lua.LNil)
		return 1
	}
}

func (h *Hash) blake2b512L(L *lua.LState) int {
	if L.GetTop() == 1 {
		h.plaintext = L.CheckString(1)

		err := h.blake2b512()
		if err != nil {
			L.Push(lua.LString(err.Error()))
			return 1
		}

		L.Push(lua.LString(h.plaintext))

		return 1
	} else if L.GetTop() == 2 {
		h.plaintext = L.CheckString(1)
		h.key = L.CheckString(2)

		err := h.blake2b512()
		if err != nil {
			L.Push(lua.LString(err.Error()))
			return 1
		}

		L.Push(lua.LString(h.plaintext))

		return 1
	} else {
		L.Push(lua.LNil)
		return 1
	}
}

func (h *Hash) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "md4":
		return L.NewFunction(h.md4L)
	case "md5":
		return L.NewFunction(h.md5L)
	case "ripemd160":
		return L.NewFunction(h.ripemd160L)
	case "sha1":
		return L.NewFunction(h.sha1L)
	case "sha3_224":
		return L.NewFunction(h.sha3224L)
	case "sha3_256":
		return L.NewFunction(h.sha3256L)
	case "sha3_384":
		return L.NewFunction(h.sha3384L)
	case "sha3_512":
		return L.NewFunction(h.sha3512L)
	case "sha224":
		return L.NewFunction(h.sha224L)
	case "sha256":
		return L.NewFunction(h.sha256L)
	case "sha384":
		return L.NewFunction(h.sha384L)
	case "sha512":
		return L.NewFunction(h.sha512L)
	case "sha512_224":
		return L.NewFunction(h.sha512224L)
	case "sha512_256":
		return L.NewFunction(h.sha512256L)
	case "blake2s_128":
		return L.NewFunction(h.blake2s128L)
	case "blake2s_256":
		return L.NewFunction(h.blake2s256L)
	case "blake2b_256":
		return L.NewFunction(h.blake2b256L)
	case "blake2b_384":
		return L.NewFunction(h.blake2b384L)
	case "blake2b_512":
		return L.NewFunction(h.blake2b512L)
	default:
		return lua.LNil
	}
}
