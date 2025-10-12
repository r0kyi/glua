# crypto

加密模块，提供基础的加密功能

## aes

**初始化**

| 参数名称 | 参数类型 | 备注                                                         |
| -------- | -------- | ------------------------------------------------------------ |
| key      | string   | 密钥，必须                                                   |
| iv       | string   | 初始化向量，必须                                             |
| mode     | string   | 加密模式，必须，可选值：<br />cbc<br />cfb<br />ofb<br />ctr<br />gcm<br />ecb |

### encrypt

aes 加密

**参数值**

| 参数名称  | 参数类型 | 备注       |
| --------- | -------- | ---------- |
| plaintext | string   | 明文，必须 |

**返回值**

| 参数名称   | 参数类型 | 备注                  |
| ---------- | -------- | --------------------- |
| ciphertext | string   | 密文                  |
| tag        | string   | 只有 **gcm** 模式返回 |
| err        | string   | 错误返回值            |

### decrypt

aes 解密

**参数值**

| 参数名称   | 参数类型 | 备注                  |
| ---------- | -------- | --------------------- |
| ciphertext | string   | 明文，必须            |
| tag        | string   | 只有 **gcm** 模式需要 |

**返回值**

| 参数名称  | 参数类型 | 备注       |
| --------- | -------- | ---------- |
| plaintext | string   | 密文       |
| err       | string   | 错误返回值 |

**demo-cbc**

```lua
local aes = glua.crypto.aes

local a = aes{
    key = "1234567890123456",
    iv = "1234567890123456",
    mode = "cbc",
}

local c, err = a.encrypt("123456")
if err ~= nil then
    print(err)
end

local p, err = a.decrypt(c)
if err ~= nil then
    print(err)
end

print(string.format("%02x", c))
print(p)

-- d637735ae9e21ba50cb686b74fab8d2c
-- 123456
```

**demo-cfb**

```lua
local aes = glua.crypto.aes

local a = aes{
    key = "1234567890123456",
    iv = "1234567890123456",
    mode = "cfb",
}

local c, err = a.encrypt("123456")
if err ~= nil then
    print(err)
end

local p, err = a.decrypt(c)
if err ~= nil then
    print(err)
end

print(string.format("%02x", c))
print(p)

-- 444efe38e96a
-- 123456
```

**demo-ofb**

```lua
local aes = glua.crypto.aes

local a = aes{
    key = "1234567890123456",
    iv = "1234567890123456",
    mode = "ofb",
}

local c, err = a.encrypt("123456")
if err ~= nil then
    print(err)
end

local p, err = a.decrypt(c)
if err ~= nil then
    print(err)
end

print(string.format("%02x", c))
print(p)

-- 444efe38e96a
-- 123456
```

**demo-ctr**

```lua
local aes = glua.crypto.aes

local a = aes{
    key = "1234567890123456",
    iv = "1234567890123456",
    mode = "ctr",
}

local c, err = a.encrypt("123456")
if err ~= nil then
    print(err)
end

local p, err = a.decrypt(c)
if err ~= nil then
    print(err)
end

print(string.format("%02x", c))
print(p)

-- 444efe38e96a
-- 123456
```

**demo-gcm**

```lua
local aes = glua.crypto.aes

local a = aes{
    key = "1234567890123456",
    iv = "1234567890123456",
    mode = "gcm",
    aad = "123456",
}

local c, tag, err = a.encrypt("123456")
if err ~= nil then
    print(err)
end

local p, err = a.decrypt(c, tag)
if err ~= nil then
    print(err)
end

print(string.format("%02x", c))
print(string.format("%02x", tag))
print(p)

-- b796d688b25b
-- 88b01e32f7b2ea5a773e68f007056c5c
-- 123456
```

**demo-ecb**

```lua
local aes = glua.crypto.aes

local a = aes{
    key = "1234567890123456",
    iv = "1234567890123456",
    mode = "ecb",
}

local c, err = a.encrypt("123456")
if err ~= nil then
    print(err)
end

local p, err = a.decrypt(c)
if err ~= nil then
    print(err)
end

print(string.format("%02x", c))
print(p)

-- c97554911e393c5cf451fa5b0c1f3f7b
-- 123456
```

## md4

md4 哈希

**参数值**

| 参数名称  | 参数类型 | 备注       |
| --------- | -------- | ---------- |
| plaintext | string   | 明文，必须 |

**返回值**

| 参数名称   | 参数类型 | 备注 |
| ---------- | -------- | ---- |
| ciphertext | string   | 密文 |

**demo**

```lua
local md4 = glua.crypto.md4

print(md4("123456"))

-- 585028aa0f794af812ee3be8804eb14a
```

## md5

md5 哈希

**参数值**

| 参数名称  | 参数类型 | 备注       |
| --------- | -------- | ---------- |
| plaintext | string   | 明文，必须 |

**返回值**

| 参数名称   | 参数类型 | 备注 |
| ---------- | -------- | ---- |
| ciphertext | string   | 密文 |

**demo**

```lua
local md5 = glua.crypto.md5

print(md5("123456"))

-- e10adc3949ba59abbe56e057f20f883e
```

## ripemd160

ripemd160 哈希

**参数值**

| 参数名称  | 参数类型 | 备注       |
| --------- | -------- | ---------- |
| plaintext | string   | 明文，必须 |

**返回值**

| 参数名称   | 参数类型 | 备注 |
| ---------- | -------- | ---- |
| ciphertext | string   | 密文 |

**demo**

```lua
local ripemd160 = glua.crypto.ripemd160

print(ripemd160("123456"))

-- d8913df37b24c97f28f840114d05bd110dbb2e44
```

## sha1

sha1 哈希

**参数值**

| 参数名称  | 参数类型 | 备注       |
| --------- | -------- | ---------- |
| plaintext | string   | 明文，必须 |

**返回值**

| 参数名称   | 参数类型 | 备注 |
| ---------- | -------- | ---- |
| ciphertext | string   | 密文 |

**demo**

```lua
local sha1 = glua.crypto.sha1

print(sha1("123456"))

-- 7c4a8d09ca3762af61e59520943dc26494f8941b
```

## sha3_224

sha3_224 哈希

**参数值**

| 参数名称  | 参数类型 | 备注       |
| --------- | -------- | ---------- |
| plaintext | string   | 明文，必须 |

**返回值**

| 参数名称   | 参数类型 | 备注 |
| ---------- | -------- | ---- |
| ciphertext | string   | 密文 |

**demo**

```lua
local sha3_224 = glua.crypto.sha3_224

print(sha3_224("123456"))

-- 6be790258b73da9441099c4cb6aeec1f0c883152dd74e7581b70a648
```

## sha3_256

sha3_256 哈希

**参数值**

| 参数名称  | 参数类型 | 备注       |
| --------- | -------- | ---------- |
| plaintext | string   | 明文，必须 |

**返回值**

| 参数名称   | 参数类型 | 备注 |
| ---------- | -------- | ---- |
| ciphertext | string   | 密文 |

**demo**

```lua
local sha3_256 = glua.crypto.sha3_256

print(sha3_256("123456"))

-- d7190eb194ff9494625514b6d178c87f99c5973e28c398969d2233f2960a573e
```

## sha3_384

sha3_384 哈希

**参数值**

| 参数名称  | 参数类型 | 备注       |
| --------- | -------- | ---------- |
| plaintext | string   | 明文，必须 |

**返回值**

| 参数名称   | 参数类型 | 备注 |
| ---------- | -------- | ---- |
| ciphertext | string   | 密文 |

**demo**

```lua
local sha3_384 = glua.crypto.sha3_384

print(sha3_384("123456"))

-- 1fb0da774034ba308fbe02f3e90dc004191df7aec3758b6be8451d09f1ff7ec18765f96e71faff637925c6be1d65f1cd
```

## sha3_512

sha3_512 哈希

**参数值**

| 参数名称  | 参数类型 | 备注       |
| --------- | -------- | ---------- |
| plaintext | string   | 明文，必须 |

**返回值**

| 参数名称   | 参数类型 | 备注 |
| ---------- | -------- | ---- |
| ciphertext | string   | 密文 |

**demo**

```lua
local sha3_512 = glua.crypto.sha3_512

print(sha3_512("123456"))

-- 64d09d9930c8ecf79e513167a588cb75439b762ce8f9b22ea59765f32aa74ca19d2f1e97dc922a3d4954594a05062917fb24d1f8e72f2ed02a58ed7534f94d27
```

## sha224

sha224 哈希

**参数值**

| 参数名称  | 参数类型 | 备注       |
| --------- | -------- | ---------- |
| plaintext | string   | 明文，必须 |

**返回值**

| 参数名称   | 参数类型 | 备注 |
| ---------- | -------- | ---- |
| ciphertext | string   | 密文 |

**demo**

```lua
local sha224 = glua.crypto.sha224

print(sha224("123456"))

-- f8cdb04495ded47615258f9dc6a3f4707fd2405434fefc3cbf4ef4e6
```

## sha256

sha256 哈希

**参数值**

| 参数名称  | 参数类型 | 备注       |
| --------- | -------- | ---------- |
| plaintext | string   | 明文，必须 |

**返回值**

| 参数名称   | 参数类型 | 备注 |
| ---------- | -------- | ---- |
| ciphertext | string   | 密文 |

**demo**

```lua
local sha256 = glua.crypto.sha256

print(sha256("123456"))

-- 8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92
```

## sha384

sha384 哈希

**参数值**

| 参数名称  | 参数类型 | 备注       |
| --------- | -------- | ---------- |
| plaintext | string   | 明文，必须 |

**返回值**

| 参数名称   | 参数类型 | 备注 |
| ---------- | -------- | ---- |
| ciphertext | string   | 密文 |

**demo**

```lua
local sha384 = glua.crypto.sha384

print(sha384("123456"))

-- 0a989ebc4a77b56a6e2bb7b19d995d185ce44090c13e2984b7ecc6d446d4b61ea9991b76a4c2f04b1b4d244841449454
```

## sha512

sha512 哈希

**参数值**

| 参数名称  | 参数类型 | 备注       |
| --------- | -------- | ---------- |
| plaintext | string   | 明文，必须 |

**返回值**

| 参数名称   | 参数类型 | 备注 |
| ---------- | -------- | ---- |
| ciphertext | string   | 密文 |

**demo**

```lua
local sha512 = glua.crypto.sha512

print(sha512("123456"))

-- ba3253876aed6bc22d4a6ff53d8406c6ad864195ed144ab5c87621b6c233b548baeae6956df346ec8c17f5ea10f35ee3cbc514797ed7ddd3145464e2a0bab413
```

## sha512_224

sha512_224 哈希

**参数值**

| 参数名称  | 参数类型 | 备注       |
| --------- | -------- | ---------- |
| plaintext | string   | 明文，必须 |

**返回值**

| 参数名称   | 参数类型 | 备注 |
| ---------- | -------- | ---- |
| ciphertext | string   | 密文 |

**demo**

```lua
local sha512_224 = glua.crypto.sha512_224

print(sha512_224("123456"))

-- 007ca663c61310fbee4c1680a5bbe70071825079b23f092713383296
```

## sha512_256

sha512_256 哈希

**参数值**

| 参数名称  | 参数类型 | 备注       |
| --------- | -------- | ---------- |
| plaintext | string   | 明文，必须 |

**返回值**

| 参数名称   | 参数类型 | 备注 |
| ---------- | -------- | ---- |
| ciphertext | string   | 密文 |

**demo**

```lua
local sha512_256 = glua.crypto.sha512_256

print(sha512_256("123456"))

-- 184b5379d5b5a7ab42d3de1d0ca1fedc1f0ffb14a7673ebd026a6369745deb72
```

## blake2s_128

blake2s_128 哈希

**参数值**

| 参数名称  | 参数类型 | 备注       |
| --------- | -------- | ---------- |
| plaintext | string   | 明文，必须 |
| key       | string   | 密钥，必须 |

**返回值**

| 参数名称   | 参数类型 | 备注 |
| ---------- | -------- | ---- |
| ciphertext | string   | 密文 |

**demo**

```lua
local blake2s_128 = glua.crypto.blake2s_128

print(blake2s_128("123456", "123456"))

-- 721386900a7976c26ad7d5b964e4a518
```

## blake2s_256

blake2s_256 哈希

**参数值**

| 参数名称  | 参数类型 | 备注         |
| --------- | -------- | ------------ |
| plaintext | string   | 明文，必须   |
| key       | string   | 密钥，非必须 |

**返回值**

| 参数名称   | 参数类型 | 备注 |
| ---------- | -------- | ---- |
| ciphertext | string   | 密文 |

**demo**

```lua
local blake2s_256 = glua.crypto.blake2s_256

print(blake2s_256("123456"))
print(blake2s_256("123456", "123456"))

-- ba2649757ec72ed0b9bd8b3063687767946145f13abcb38e2718fdaad6c771e0
-- 1bc1a1a13702b26b63611d0852a67a462bde6193ebec7f58d73431499007076f
```

## blake2b_256

blake2b_256 哈希

**参数值**

| 参数名称  | 参数类型 | 备注         |
| --------- | -------- | ------------ |
| plaintext | string   | 明文，必须   |
| key       | string   | 密钥，非必须 |

**返回值**

| 参数名称   | 参数类型 | 备注 |
| ---------- | -------- | ---- |
| ciphertext | string   | 密文 |

**demo**

```lua
local blake2b_256 = glua.crypto.blake2b_256

print(blake2b_256("123456"))
print(blake2b_256("123456", "123456"))

-- 4361be62001d25deb2bd85fab3e46011afae57539026d8d37d57f45e29571271
-- 8c6850c8517c68a22e54a3e8ab540ad4ae4175755b9b84b1f6bccaa9a5ed9576
```

## blake2b_384

blake2b_384 哈希

**参数值**

| 参数名称  | 参数类型 | 备注         |
| --------- | -------- | ------------ |
| plaintext | string   | 明文，必须   |
| key       | string   | 密钥，非必须 |

**返回值**

| 参数名称   | 参数类型 | 备注 |
| ---------- | -------- | ---- |
| ciphertext | string   | 密文 |

**demo**

```lua
local blake2b_384 = glua.crypto.blake2b_384

print(blake2b_384("123456"))
print(blake2b_384("123456", "123456"))

-- 1967dc11fa8a65943f44a62d60df70e200b298346e1beafdaac499cd95705d6bd87a3c268a0eb5c9ca578cb56ecef40c
-- 8d82458aa906f0c10ce0b045a237709907a6c49b0d430bc156d6ca3c1372fa4edd08d55597a3beb9a06e364e98424620
```

## blake2b_512

blake2b_512 哈希

**参数值**

| 参数名称  | 参数类型 | 备注         |
| --------- | -------- | ------------ |
| plaintext | string   | 明文，必须   |
| key       | string   | 密钥，非必须 |

**返回值**

| 参数名称   | 参数类型 | 备注 |
| ---------- | -------- | ---- |
| ciphertext | string   | 密文 |

**demo**

```lua
local blake2b_512 = glua.crypto.blake2b_512

print(blake2b_512("123456"))
print(blake2b_512("123456", "123456"))

-- b3910b0f4b6f1aede44da90bb7705a868b265861b36e6f7f29dba7223f6f1ce7b10e0dd25e47deb70bd7f3b24f7da653409cd9014f8715e4013c15fee38ab418
-- ef69b2d9dbb2818e98351c9bc6620607b4cc5ab6453b0584defdeec5ff08b43f66f0fbf43a3893f6a6356ecda02c34c9e019dc3b033a7c7d66c6d8730b982665
```

