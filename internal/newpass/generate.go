package newpass

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func Generate(length, count int)  error{
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := 0; i < count; i++ {
		for i := range result {
			num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
			if err != nil {
				return  err
			}
			result[i] = charset[num.Int64()]
		}
		fmt.Println(string(result))
	}

	return nil
}