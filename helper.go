package jsontostruct

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"io"
	"unicode/utf16"
)

var (
	magicNum   uint32 = 0xBAADF00D
	passphrase string
)

const sql_v2 = 0x02

func SetPasshrase(p string) {
	passphrase = p
}

func passphraseToKey() []byte {
	var runes []rune
	for _, l := range passphrase {
		runes = append(runes, l)
	}
	windowsPhrase := utf16.Encode(runes)
	var keybuf bytes.Buffer
	binary.Write(&keybuf, binary.LittleEndian, windowsPhrase)

	key := sha256.Sum256(keybuf.Bytes())
	return key[:]
}

func DecryptByPassphrase(cyphertext []byte) (string, error) {
	// Only support V2 (AES)
	if len(cyphertext) > 0 {
		if cyphertext[0] != sql_v2 {
			return "", errors.New("required sql v2")
		}
	} else {
		return "", errors.New("chyphertext is empty")
	}

	key := passphraseToKey()

	block, e := aes.NewCipher(key)
	if e != nil {
		return "", e
	}

	iv := make([]byte, aes.BlockSize)
	io.ReadFull(rand.Reader, iv)

	mode := cipher.NewCBCDecrypter(block, cyphertext[4:20])

	dst := make([]byte, len(cyphertext)-20)
	mode.CryptBlocks(dst, cyphertext[20:])

	if binary.LittleEndian.Uint32(dst[:4]) != magicNum {
		return "", errors.New("magic bytes failed")
	}
	if binary.LittleEndian.Uint16(dst[4:6]) != 0 {
		return "", errors.New("authenticator unsupported")
	}

	ptLen := binary.LittleEndian.Uint16(dst[6:8])

	return string(dst[8 : 8+ptLen]), nil
}
