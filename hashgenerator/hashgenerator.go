//Package hashgenerator generates the hash of a given string
package hashgenerator

import "crypto/md5"

//GenerateHash is a global function used to generate hash of a given string
func GenerateHash(name string) [16]byte {
	data := []byte(name)
	x := md5.Sum(data)
	return x
}
