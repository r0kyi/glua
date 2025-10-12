# format

格式化字符串，与 go 用法一致

**参数值**

| 参数名称 | 参数类型                                 | 备注                   |
| -------- | ---------------------------------------- | ---------------------- |
| format   | string                                   | 格式化字符串，必须     |
| args     | nil<br />boolean<br />number<br />string | 任意数量的变量，非必须 |

**返回值**

| 参数名称 | 参数类型 | 备注               |
| -------- | -------- | ------------------ |
| result   | string   | 格式化之后的字符串 |

**demo**

```lua
local f = glua.format

local s = "hello"
local v = "world"
print(f("%s %v: %d", s, v, 123))

-- hello world: 123
```

