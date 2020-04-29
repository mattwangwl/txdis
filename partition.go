package txdis

import (
	"bytes"
	"crypto/sha512"
	"encoding/binary"
)

func newPartition(salt []byte, size int) iPartition {
	return &partition{
		salt: salt,
		size: uint64(size),
	}
}

type iPartition interface {
	Calculate(id int64) int
}

type partition struct {
	salt []byte
	size uint64
}

// 區間計算
func (p partition) Calculate(id int64) int {
	hashID := p.hash(id)
	newID := p.parse(hashID)

	result := newID % p.size

	return int(result)
}

// 加鹽雜揍
func (p partition) hash(numb int64) [64]byte {
	data := make([]byte, 8)
	binary.LittleEndian.PutUint64(data, uint64(numb))

	var buf bytes.Buffer
	buf.Write(data)
	buf.Write(p.salt)

	return sha512.Sum512(buf.Bytes())
}

// 返回成uint64
func (partition) parse(b [64]byte) uint64 {
	return binary.LittleEndian.Uint64(b[:])
}
