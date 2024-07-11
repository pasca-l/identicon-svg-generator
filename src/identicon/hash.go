package identicon

import (
	"crypto/md5"
	"fmt"
)

type Hash struct {
	hash []byte // 4-bit data array
}

func generateMd5Hash(input string) Hash {
	// get MD5 hash value from input string
	md5Hash := md5.Sum([]byte(input))
	fmt.Printf("Hash value of \"%s\": %x\n", input, md5Hash)

	nibbles := make([]byte, 0, 4*len(md5Hash))
	for _, B := range md5Hash {
		nibbles = append(nibbles, (B>>4)&0x0f, B&0x0f)
	}

	return Hash{
		hash: nibbles,
	}
}
