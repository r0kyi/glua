package glua

import (
	"github.com/r0kyi/glua/plugin/base"
	"github.com/r0kyi/glua/plugin/cron"
	"github.com/r0kyi/glua/plugin/crypto"
	"github.com/r0kyi/glua/plugin/format"
	"github.com/r0kyi/glua/plugin/http"
	"github.com/r0kyi/glua/plugin/json"
	"github.com/r0kyi/glua/plugin/re"
	"github.com/r0kyi/glua/plugin/time"
	"github.com/r0kyi/glua/plugin/web"
	"github.com/r0kyi/glua/plugin/xml"
	"github.com/r0kyi/glua/plugin/yaml"
	lua "github.com/yuin/gopher-lua"
)

func NewState() *lua.LState {
	L := lua.NewState()
	table := L.NewTable()

	table.RawSetString("base", base.Preload(L))
	table.RawSetString("cron", cron.Preload(L))
	table.RawSetString("crypto", crypto.Preload(L))
	table.RawSetString("format", format.Preload(L))
	table.RawSetString("http", http.Preload(L))
	table.RawSetString("json", json.Preload(L))
	table.RawSetString("re", re.Preload(L))
	table.RawSetString("time", time.Preload(L))
	table.RawSetString("web", web.Preload(L))
	table.RawSetString("xml", xml.Preload(L))
	table.RawSetString("yaml", yaml.Preload(L))

	L.SetGlobal("glua", table)
	return L
}
