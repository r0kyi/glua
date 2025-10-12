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

type Hash struct {
	plaintext  string
	ciphertext string
	key        string
}

func (h *Hash) md4() {
	hasher := md4.New()
	hasher.Write(core.S2B(h.plaintext))
	h.ciphertext = hex.EncodeToString(hasher.Sum(nil))
}

func (h *Hash) md5() {
	hasher := md5.New()
	hasher.Write(core.S2B(h.plaintext))
	h.ciphertext = hex.EncodeToString(hasher.Sum(nil))
}

func (h *Hash) ripemd160() {
	hasher := ripemd160.New()
	hasher.Write(core.S2B(h.plaintext))
	h.ciphertext = hex.EncodeToString(hasher.Sum(nil))
}

func (h *Hash) sha1() {
	hasher := sha1.New()
	hasher.Write(core.S2B(h.plaintext))
	h.ciphertext = hex.EncodeToString(hasher.Sum(nil))
}

func (h *Hash) sha3224() {
	hasher := sha3.New224()
	hasher.Write(core.S2B(h.plaintext))
	h.ciphertext = hex.EncodeToString(hasher.Sum(nil))
}

func (h *Hash) sha3256() {
	hasher := sha3.New256()
	hasher.Write(core.S2B(h.plaintext))
	h.ciphertext = hex.EncodeToString(hasher.Sum(nil))
}

func (h *Hash) sha3384() {
	hasher := sha3.New384()
	hasher.Write(core.S2B(h.plaintext))
	h.ciphertext = hex.EncodeToString(hasher.Sum(nil))
}

func (h *Hash) sha3512() {
	hasher := sha3.New512()
	hasher.Write(core.S2B(h.plaintext))
	h.ciphertext = hex.EncodeToString(hasher.Sum(nil))
}

func (h *Hash) sha224() {
	hasher := sha256.New224()
	hasher.Write(core.S2B(h.plaintext))
	h.ciphertext = hex.EncodeToString(hasher.Sum(nil))
}

func (h *Hash) sha256() {
	hasher := sha256.New()
	hasher.Write(core.S2B(h.plaintext))
	h.ciphertext = hex.EncodeToString(hasher.Sum(nil))
}

func (h *Hash) sha384() {
	hasher := sha512.New384()
	hasher.Write(core.S2B(h.plaintext))
	h.ciphertext = hex.EncodeToString(hasher.Sum(nil))
}

func (h *Hash) sha512() {
	hasher := sha512.New()
	hasher.Write(core.S2B(h.plaintext))
	h.ciphertext = hex.EncodeToString(hasher.Sum(nil))
}

func (h *Hash) sha512224() {
	hasher := sha512.New512_224()
	hasher.Write(core.S2B(h.plaintext))
	h.ciphertext = hex.EncodeToString(hasher.Sum(nil))
}

func (h *Hash) sha512256() {
	hasher := sha512.New512_256()
	hasher.Write(core.S2B(h.plaintext))
	h.ciphertext = hex.EncodeToString(hasher.Sum(nil))
}

func (h *Hash) blake2s128() error {
	hasher, err := blake2s.New128(core.S2B(h.key))
	if err != nil {
		return err
	}

	hasher.Write(core.S2B(h.plaintext))
	h.ciphertext = hex.EncodeToString(hasher.Sum(nil))

	return nil
}

func (h *Hash) blake2s256() error {
	hasher, err := blake2s.New256(core.S2B(h.key))
	if err != nil {
		return err
	}

	hasher.Write(core.S2B(h.plaintext))
	h.ciphertext = hex.EncodeToString(hasher.Sum(nil))

	return nil
}

func (h *Hash) blake2b256() error {
	hasher, err := blake2b.New256(core.S2B(h.key))
	if err != nil {
		return err
	}

	hasher.Write(core.S2B(h.plaintext))
	h.ciphertext = hex.EncodeToString(hasher.Sum(nil))

	return nil
}

func (h *Hash) blake2b384() error {
	hasher, err := blake2b.New384(core.S2B(h.key))
	if err != nil {
		return err
	}

	hasher.Write(core.S2B(h.plaintext))
	h.ciphertext = hex.EncodeToString(hasher.Sum(nil))

	return nil
}

func (h *Hash) blake2b512() error {
	hasher, err := blake2b.New512(core.S2B(h.key))
	if err != nil {
		return err
	}

	hasher.Write(core.S2B(h.plaintext))
	h.ciphertext = hex.EncodeToString(hasher.Sum(nil))

	return nil
}
