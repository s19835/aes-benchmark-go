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

			f, _ := os.OpenFile("results/benchmark.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			defer f.Close()

			fmt.Fprintf(f, "%s,%s,%s,%d\n", method, size, elapsed, len(enc))
		}
	}
}
