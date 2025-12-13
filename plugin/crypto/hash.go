package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"

	"github.com/r0kyi/glua/core"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

type Hash struct{}

func (h *Hash) md4(plaintext string) string {
	hasher := md4.New()
	hasher.Write(core.S2B(plaintext))
	ciphertext := hex.EncodeToString(hasher.Sum(nil))

	return ciphertext
}

func (h *Hash) md5(plaintext string) string {
	hasher := md5.New()
	hasher.Write(core.S2B(plaintext))
	ciphertext := hex.EncodeToString(hasher.Sum(nil))

	return ciphertext
}

func (h *Hash) ripemd160(plaintext string) string {
	hasher := ripemd160.New()
	hasher.Write(core.S2B(plaintext))
	ciphertext := hex.EncodeToString(hasher.Sum(nil))

	return ciphertext
}

func (h *Hash) sha1(plaintext string) string {
	hasher := sha1.New()
	hasher.Write(core.S2B(plaintext))
	ciphertext := hex.EncodeToString(hasher.Sum(nil))

	return ciphertext
}

func (h *Hash) sha3224(plaintext string) string {
	hasher := sha3.New224()
	hasher.Write(core.S2B(plaintext))
	ciphertext := hex.EncodeToString(hasher.Sum(nil))

	return ciphertext
}

func (h *Hash) sha3256(plaintext string) string {
	hasher := sha3.New256()
	hasher.Write(core.S2B(plaintext))
	ciphertext := hex.EncodeToString(hasher.Sum(nil))

	return ciphertext
}

func (h *Hash) sha3384(plaintext string) string {
	hasher := sha3.New384()
	hasher.Write(core.S2B(plaintext))
	ciphertext := hex.EncodeToString(hasher.Sum(nil))

	return ciphertext
}

func (h *Hash) sha3512(plaintext string) string {
	hasher := sha3.New512()
	hasher.Write(core.S2B(plaintext))
	ciphertext := hex.EncodeToString(hasher.Sum(nil))

	return ciphertext
}

func (h *Hash) sha224(plaintext string) string {
	hasher := sha256.New224()
	hasher.Write(core.S2B(plaintext))
	ciphertext := hex.EncodeToString(hasher.Sum(nil))

	return ciphertext
}

func (h *Hash) sha256(plaintext string) string {
	hasher := sha256.New()
	hasher.Write(core.S2B(plaintext))
	ciphertext := hex.EncodeToString(hasher.Sum(nil))

	return ciphertext
}

func (h *Hash) sha384(plaintext string) string {
	hasher := sha512.New384()
	hasher.Write(core.S2B(plaintext))
	ciphertext := hex.EncodeToString(hasher.Sum(nil))

	return ciphertext
}

func (h *Hash) sha512(plaintext string) string {
	hasher := sha512.New()
	hasher.Write(core.S2B(plaintext))
	ciphertext := hex.EncodeToString(hasher.Sum(nil))

	return ciphertext
}

func (h *Hash) sha512224(plaintext string) string {
	hasher := sha512.New512_224()
	hasher.Write(core.S2B(plaintext))
	ciphertext := hex.EncodeToString(hasher.Sum(nil))

	return ciphertext
}

func (h *Hash) sha512256(plaintext string) string {
	hasher := sha512.New512_256()
	hasher.Write(core.S2B(plaintext))
	ciphertext := hex.EncodeToString(hasher.Sum(nil))

	return ciphertext
}

func (h *Hash) blake2s128(plaintext string, key string) (string, error) {
	hasher, err := blake2s.New128(core.S2B(key))
	if err != nil {
		return "", err
	}

	hasher.Write(core.S2B(plaintext))
	ciphertext := hex.EncodeToString(hasher.Sum(nil))

	return ciphertext, nil
}

func (h *Hash) blake2s256(plaintext string, key string) (string, error) {
	hasher, err := blake2s.New256(core.S2B(key))
	if err != nil {
		return "", err
	}

	hasher.Write(core.S2B(plaintext))
	ciphertext := hex.EncodeToString(hasher.Sum(nil))

	return ciphertext, nil
}

func (h *Hash) blake2b256(plaintext string, key string) (string, error) {
	hasher, err := blake2b.New256(core.S2B(key))
	if err != nil {
		return "", err
	}

	hasher.Write(core.S2B(plaintext))
	ciphertext := hex.EncodeToString(hasher.Sum(nil))

	return ciphertext, nil
}

func (h *Hash) blake2b384(plaintext string, key string) (string, error) {
	hasher, err := blake2b.New384(core.S2B(key))
	if err != nil {
		return "", err
	}

	hasher.Write(core.S2B(plaintext))
	ciphertext := hex.EncodeToString(hasher.Sum(nil))

	return ciphertext, nil
}

func (h *Hash) blake2b512(plaintext string, key string) (string, error) {
	hasher, err := blake2b.New512(core.S2B(key))
	if err != nil {
		return "", err
	}

	hasher.Write(core.S2B(plaintext))
	ciphertext := hex.EncodeToString(hasher.Sum(nil))

	return ciphertext, nil
}
