package common

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
)

func HmacMD5(key []byte, data []byte) []byte {
	hmacMD5 := hmac.New(md5.New, key)
	hmacMD5.Write(data)
	return hmacMD5.Sum(nil)[:10]
}

func HmacSHA1(key []byte, data []byte) []byte {
	hmacSHA1 := hmac.New(sha1.New, key)
	hmacSHA1.Write(data)
	return hmacSHA1.Sum(nil)[:10]
}

func Md5Sum(d []byte) []byte {
	h := md5.New()
	h.Write(d)
	return h.Sum(nil)
}

func EVPBytesToKey(password string, keyLen int) (key []byte) {
	const md5Len = 16

	cnt := (keyLen-1)/md5Len + 1
	m := make([]byte, cnt * md5Len)
	copy(m, Md5Sum([]byte(password)))

	// Repeatedly call md5 until bytes generated is enough.
	// Each call to md5 uses data: prev md5 sum + password.
	d := make([]byte, md5Len + len(password))
	start := 0
	for i := 1; i < cnt; i++ {
		start += md5Len
		copy(d, m[start - md5Len:start])
		copy(d[md5Len:], password)
		copy(m[start:], Md5Sum(d))
	}
	return m[:keyLen]
}
