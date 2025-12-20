# glua

## 简介

glua 基于 [gopher-lua](https://github.com/r0kyi/gopher-lua) 二次开发，在原版基础上新增了许多模块

## 安装

```bash
go get github.com/r0kyi/glua
```

## 使用

直接安装二进制版本

```bash
go install github.com/r0kyi/glua/cmd/glua@latest
```

在 go 代码中使用

```lua
package main

import (
	"os"

	"github.com/r0kyi/glua"
)

func main() {
	L := glua.NewState()
	defer L.Close()

	args := os.Args
	if len(args) < 2 {
		println("usage: glua <lua file>")
		return
	}

	if err := L.DoFile(args[1]); err != nil {
		println(err.Error())
	}
}

```

## 模块列表

| 模块名称                     | 模块介绍     |
| ---------------------------- | ------------ |
| [base](docs/base.md)         | base 编解码  |
| [cron](docs/cron.md)         | 计划任务     |
| [crypto](docs/crypto.md)     | 加解密       |
| [database](docs/database.md) | 数据库连接   |
| [format](docs/format.md)     | 格式化字符串 |
| [http](docs/http.md)         | http 请求    |
| [json](docs/json.md)         | json 解析    |
| [jwt](docs/jwt.md)           | jwt 解析     |
| [re](docs/re.md)             | 正则         |
| [time](docs/time.md)         | 时间         |
| [web](docs/web.md)           | web 服务     |
| [xml](docs/xml.md)           | xml 解析     |
| [yaml](docs/yaml.md)         | yaml 解析    |

## 特别鸣谢

- [gopher-lua](https://github.com/yuin/gopher-lua)