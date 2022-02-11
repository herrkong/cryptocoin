package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func Hash(Input string) (OutPut string, err error) {
	sha256Bytes := sha256.Sum256([]byte(Input))
	OutPut = hex.EncodeToString(sha256Bytes[:])
	fmt.Println(OutPut,len(OutPut))
	return OutPut, nil
}

