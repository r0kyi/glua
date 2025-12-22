package glua

import (
	"github.com/r0kyi/glua/plugin/base"
	"github.com/r0kyi/glua/plugin/cron"
	"github.com/r0kyi/glua/plugin/crypto"
	"github.com/r0kyi/glua/plugin/database"
	"github.com/r0kyi/glua/plugin/format"
	"github.com/r0kyi/glua/plugin/http"
	"github.com/r0kyi/glua/plugin/json"
	"github.com/r0kyi/glua/plugin/jwt"
	"github.com/r0kyi/glua/plugin/re"
	"github.com/r0kyi/glua/plugin/time"
	"github.com/r0kyi/glua/plugin/uuid"
	"github.com/r0kyi/glua/plugin/validator"
	"github.com/r0kyi/glua/plugin/web"
	"github.com/r0kyi/glua/plugin/xml"
	"github.com/r0kyi/glua/plugin/yaml"
	lua "github.com/r0kyi/gopher-lua"
)

func NewState() *lua.LState {
	L := lua.NewState()
	table := L.NewTable()

	table.RawSetString("base", base.Preload())
	table.RawSetString("cron", cron.Preload())
	table.RawSetString("crypto", crypto.Preload())
	table.RawSetString("database", database.Preload())
	table.RawSetString("format", format.Preload())
	table.RawSetString("http", http.Preload())
	table.RawSetString("json", json.Preload())
	table.RawSetString("jwt", jwt.Preload())
	table.RawSetString("re", re.Preload())
	table.RawSetString("time", time.Preload())
	table.RawSetString("uuid", uuid.Preload())
	table.RawSetString("validator", validator.Preload())
	table.RawSetString("web", web.Preload())
	table.RawSetString("xml", xml.Preload())
	table.RawSetString("yaml", yaml.Preload())

	L.SetGlobal("glua", table)
	return L
}
