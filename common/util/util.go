package util

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"io"
	mrand "math/rand"
	"strconv"
	"time"
)

//random
func RandNumber(max int) int {
	mrand.Seed(time.Now().UnixNano())
	return mrand.Intn(max)
}

func RandSMSCode() string {
	return RandNumericString(6)
}

func RandNumericString(size int) string {
	mrand.Seed(time.Now().UnixNano())
	var buf bytes.Buffer
	for i := 0; i < size; i++ {
		buf.WriteString(strconv.FormatInt(int64(mrand.Intn(10)), 10))
	}

	return buf.String()
}

const chars = "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringv2(size int) string {
	mrand.Seed(time.Now().UnixNano())
	var buf bytes.Buffer
	for i := 0; i < size; i++ {
		buf.WriteByte(chars[mrand.Intn(len(chars))])
	}

	return buf.String()
}

func RandString(size int) (string, error) {
	bl := size / 2
	if size%2 != 0 {
		bl++
	}

	buf := make([]byte, bl)
	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(buf), nil
}

//40 bytes string(sha1 hash
func RandSalt() (string, error) {
	baseStr, err := RandString(10)
	if err != nil {
		return "", err
	}

	str := baseStr + strconv.FormatInt(time.Now().UnixNano(), 10)
	hash := sha1.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func CryptoPassword(salt, password string) string {
	if salt == "" {
		// FIXME : 目前是为了测试兼容而做的临时解决方案。
		return password
	}

	var buf bytes.Buffer
	buf.WriteByte('^')
	buf.WriteString(password)
	buf.WriteByte(':')
	buf.WriteString(salt)
	buf.WriteByte('$')

	data := sha1.Sum(buf.Bytes())
	return hex.EncodeToString(data[:])
}
