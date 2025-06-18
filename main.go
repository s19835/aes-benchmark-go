package main

import (
	"fmt"
	"os"
	"time"

	"github.com/s19835/aes-benchmark-go/benchmarks"
)

func main() {
	inputsize := []string{"1KB", "1MB", "100MB"}
	methods := []string{"standard", "cryptopasta"}

	for _, method := range methods {
		for _, size := range inputsize {
			data, err := os.ReadFile("benchmarks/data/" + size + ".txt")
			if err != nil {
				panic(err)
			}

			start := time.Now()
			var enc []byte

			switch method {
			case "standard":
				enc = benchmarks.EncryptStandard(data)
			case "cryptopasta":
				enc = benchmarks.EncryptCryptopasta(data)
			}

			elapsed := time.Since(start)
			fmt.Printf("Method: %s | Size: %s | Time: %s | Encrypted Length: %d\n", method, size, elapsed, len(enc))
		}
	}
}
