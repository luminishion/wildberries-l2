package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

import (
	"bytes"
	"fmt"
)

type lzma struct {
}

func (l *lzma) compress(buf []byte) []byte {
	return buf
}

func (l *lzma) decompress(buf []byte) []byte {
	return buf
}

type aes struct {
	key []byte
}

func (c *aes) encrypt(buf []byte) []byte {
	return buf
}

func (c *aes) decrypt(buf []byte) []byte {
	return buf
}

type hmac struct {
	key []byte
}

func (h *hmac) get(buf []byte) []byte {
	return []byte("test")
}

func (h *hmac) size() int {
	return 4
}

type messageFacade struct {
	crypto     *aes
	compressor *lzma
	mac        *hmac
}

func NewMessageFacade(cryptoKey, hmacKey []byte) *messageFacade {
	return &messageFacade{
		&aes{cryptoKey},
		&lzma{},
		&hmac{hmacKey},
	}
}

func (m *messageFacade) Encode(buf []byte) []byte {
	compressed := m.compressor.compress(buf)
	encrypted := m.crypto.encrypt(compressed)

	macBuf := m.mac.get(encrypted)

	return append(macBuf, encrypted...)
}

func (m *messageFacade) Decode(buf []byte) []byte {
	hsize := m.mac.size()

	macBuf := buf[:hsize]
	encrypted := buf[hsize:]

	if bytes.Compare(macBuf, m.mac.get(encrypted)) != 0 {
		return nil
	}

	decrypted := m.crypto.decrypt(encrypted)
	decompressed := m.compressor.decompress(decrypted)

	return decompressed
}

func main() {
	key := []byte{1, 2, 3}

	message := NewMessageFacade(key, key)

	encoded := message.Encode([]byte{1, 1, 1})
	decoded := message.Decode(encoded)

	fmt.Println(decoded)
}
