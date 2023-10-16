package shortener

import (
	"crypto/sha256"
	"fmt"
	"github.com/itchyny/base58-go"
	"math/big"
	"os"
)

func sha2560f(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encoded(input []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(input)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

func GenerateShortUrl(initialUrl string, userId string) string {
	urlHashedBytes := sha2560f(initialUrl + userId)
	generatedNumber := new(big.Int).SetBytes(urlHashedBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}
