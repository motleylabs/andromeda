package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcutil/base58"
	"github.com/kevinburke/nacl/sign"
)

func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

func ValidateMessage(pubkey, msg, nonce string) bool {
	signedMsgBytes := base58.Decode(msg)
	if nonce != string(signedMsgBytes[sign.SignatureSize:]) {
		fmt.Println("Nonce is not equal")
		return false
	}

	return sign.Verify(signedMsgBytes, base58.Decode(pubkey))
}
