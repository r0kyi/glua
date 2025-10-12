package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"

	"github.com/r0kyi/glua/core"
)

type Aes struct {
	plaintext  string
	ciphertext string
	Key        string `lua:"key"`
	Iv         string `lua:"iv"`
	Mode       string `lua:"mode"`
	tag        string
	AAD        string `lua:"aad"`
}

func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("invalid padding size")
	}

	padding := int(data[length-1])
	if padding > length {
		return nil, errors.New("invalid padding size")
	}

	return data[:length-padding], nil
}

func (a *Aes) cbcEncrypt() error {
	block, err := aes.NewCipher(core.S2B(a.Key))
	if err != nil {
		return err
	}

	iv := core.S2B(a.Iv)
	if len(iv) != block.BlockSize() {
		return errors.New("invalid iv size")
	}

	plainPadded := pkcs7Padding(core.S2B(a.plaintext), block.BlockSize())
	cipherBytes := make([]byte, len(plainPadded))
	cipher.NewCBCEncrypter(block, iv).CryptBlocks(cipherBytes, plainPadded)
	a.ciphertext = core.B2S(cipherBytes)

	return nil
}

func (a *Aes) cbcDecrypt() error {
	block, err := aes.NewCipher(core.S2B(a.Key))
	if err != nil {
		return err
	}

	cipherBytes := core.S2B(a.ciphertext)
	if len(cipherBytes)%block.BlockSize() != 0 {
		return errors.New("ciphertext is not a multiple of block size")
	}

	iv := core.S2B(a.Iv)
	if len(iv) != block.BlockSize() {
		return errors.New("invalid iv size")
	}

	plainBytes := make([]byte, len(cipherBytes))
	cipher.NewCBCDecrypter(block, iv).CryptBlocks(plainBytes, cipherBytes)

	plainBytes, err = pkcs7UnPadding(plainBytes)
	if err != nil {
		return err
	}

	a.plaintext = core.B2S(plainBytes)

	return nil
}

func (a *Aes) cfbEncrypt() error {
	block, err := aes.NewCipher(core.S2B(a.Key))
	if err != nil {
		return err
	}

	iv := core.S2B(a.Iv)
	if len(iv) != block.BlockSize() {
		return errors.New("invalid iv size")
	}

	plainBytes := core.S2B(a.plaintext)
	cipherBytes := make([]byte, len(plainBytes))
	cipher.NewCFBEncrypter(block, iv).XORKeyStream(cipherBytes, plainBytes)
	a.ciphertext = core.B2S(cipherBytes)

	return nil
}

func (a *Aes) cfbDecrypt() error {
	block, err := aes.NewCipher(core.S2B(a.Key))
	if err != nil {
		return err
	}

	iv := core.S2B(a.Iv)
	if len(iv) != block.BlockSize() {
		return errors.New("invalid iv size")
	}

	cipherBytes := core.S2B(a.ciphertext)
	plainBytes := make([]byte, len(cipherBytes))
	cipher.NewCFBDecrypter(block, iv).XORKeyStream(plainBytes, cipherBytes)
	a.plaintext = core.B2S(plainBytes)

	return nil
}

func (a *Aes) ofbEncrypt() error {
	block, err := aes.NewCipher(core.S2B(a.Key))
	if err != nil {
		return err
	}

	iv := core.S2B(a.Iv)
	if len(iv) != block.BlockSize() {
		return errors.New("invalid iv size")
	}

	plainBytes := core.S2B(a.plaintext)
	cipherBytes := make([]byte, len(plainBytes))
	cipher.NewOFB(block, iv).XORKeyStream(cipherBytes, plainBytes)
	a.ciphertext = core.B2S(cipherBytes)

	return nil
}

func (a *Aes) ofbDecrypt() error {
	block, err := aes.NewCipher(core.S2B(a.Key))
	if err != nil {
		return err
	}

	iv := core.S2B(a.Iv)
	if len(iv) != block.BlockSize() {
		return errors.New("invalid iv size")
	}

	cipherBytes := core.S2B(a.ciphertext)
	plainBytes := make([]byte, len(cipherBytes))
	cipher.NewOFB(block, iv).XORKeyStream(plainBytes, cipherBytes)
	a.plaintext = core.B2S(plainBytes)

	return nil
}

func (a *Aes) ctrEncrypt() error {
	block, err := aes.NewCipher(core.S2B(a.Key))
	if err != nil {
		return err
	}

	iv := core.S2B(a.Iv)
	if len(iv) != block.BlockSize() {
		return errors.New("invalid iv size")
	}

	plainBytes := core.S2B(a.plaintext)
	cipherBytes := make([]byte, len(plainBytes))
	cipher.NewCTR(block, iv).XORKeyStream(cipherBytes, plainBytes)
	a.ciphertext = core.B2S(cipherBytes)

	return nil
}

func (a *Aes) ctrDecrypt() error {
	block, err := aes.NewCipher(core.S2B(a.Key))
	if err != nil {
		return err
	}

	iv := core.S2B(a.Iv)
	if len(iv) != block.BlockSize() {
		return errors.New("invalid iv size")
	}

	cipherBytes := core.S2B(a.ciphertext)
	plainBytes := make([]byte, len(cipherBytes))
	cipher.NewCTR(block, iv).XORKeyStream(plainBytes, cipherBytes)
	a.plaintext = core.B2S(plainBytes)

	return nil
}

func (a *Aes) gcmEncrypt() error {
	block, err := aes.NewCipher(core.S2B(a.Key))
	if err != nil {
		return err
	}

	iv := core.S2B(a.Iv)
	gcm, err := cipher.NewGCMWithNonceSize(block, len(iv))
	if err != nil {
		return err
	}

	plainBytes := core.S2B(a.plaintext)
	cipherWithTag := gcm.Seal(nil, iv, plainBytes, core.S2B(a.AAD))

	tagLen := gcm.Overhead()
	if len(cipherWithTag) < tagLen {
		return errors.New("ciphertext too short")
	}

	a.ciphertext = core.B2S(cipherWithTag[:len(cipherWithTag)-tagLen])
	a.tag = core.B2S(cipherWithTag[len(cipherWithTag)-tagLen:])

	return nil
}

func (a *Aes) gcmDecrypt() error {
	block, err := aes.NewCipher(core.S2B(a.Key))
	if err != nil {
		return err
	}

	iv := core.S2B(a.Iv)

	gcm, err := cipher.NewGCMWithNonceSize(block, len(iv))
	if err != nil {
		return err
	}

	cipherWithTag := append(core.S2B(a.ciphertext), core.S2B(a.tag)...)

	plainBytes, err := gcm.Open(nil, iv, cipherWithTag, core.S2B(a.AAD))
	if err != nil {
		return err
	}

	a.plaintext = core.B2S(plainBytes)

	return nil
}

func (a *Aes) ecbEncrypt() error {
	block, err := aes.NewCipher(core.S2B(a.Key))
	if err != nil {
		return err
	}

	blockSize := block.BlockSize()
	plainPadded := pkcs7Padding(core.S2B(a.plaintext), blockSize)

	encrypted := make([]byte, len(plainPadded))
	for bs := 0; bs < len(plainPadded); bs += blockSize {
		block.Encrypt(encrypted[bs:bs+blockSize], plainPadded[bs:bs+blockSize])
	}

	a.ciphertext = core.B2S(encrypted)

	return nil
}

func (a *Aes) ecbDecrypt() error {
	block, err := aes.NewCipher(core.S2B(a.Key))
	if err != nil {
		return err
	}

	blockSize := block.BlockSize()

	cipherBytes := core.S2B(a.ciphertext)
	if len(cipherBytes)%blockSize != 0 {
		return errors.New("invalid ciphertext length")
	}

	decrypted := make([]byte, len(cipherBytes))
	for bs := 0; bs < len(cipherBytes); bs += blockSize {
		block.Decrypt(decrypted[bs:bs+blockSize], cipherBytes[bs:bs+blockSize])
	}

	plainBytes, err := pkcs7UnPadding(decrypted)
	if err != nil {
		return err
	}

	a.plaintext = core.B2S(plainBytes)

	return nil
}
