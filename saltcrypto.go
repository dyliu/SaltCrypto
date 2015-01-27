package SaltCrypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

/**
 * 带盐值加密
 */
func Crypto(password, salt string) string {
	prf := hmac.New(sha256.New, []byte(password))
	hashLen := prf.Size()
	numBlocks := (32 + hashLen - 1) / hashLen
	var buf [4]byte
	dk := make([]byte, 0, numBlocks*hashLen)
	U := make([]byte, hashLen)
	for block := 1; block <= numBlocks; block++ {
		prf.Reset()
		prf.Write([]byte(salt))
		buf[0] = byte(block >> 24)
		buf[1] = byte(block >> 16)
		buf[2] = byte(block >> 8)
		buf[3] = byte(block)
		prf.Write(buf[:4])
		dk = prf.Sum(dk)
		T := dk[len(dk)-hashLen:]
		copy(U, T)
		for n := 2; n <= 10000; n++ {
			prf.Reset()
			prf.Write(U)
			U = U[:0]
			U = prf.Sum(U)
			// fmt.Println(T)
			for x := range U {
				T[x] ^= U[x]
			}
			// fmt.Println(T)
		}
	}
	return fmt.Sprintf("%x", dk[:32])
}
