# aes-benchmark-go

This study benchmarks and compares the performance of mainstream Go crypto libraries that use AES-256 for encryption/decryption speed, memory, and simplicity of code. With controlled simulations on files of different sizes (1KB, 1MB, 100MB), we compare how well each library performs symmetric key encryption under realistic backend limitations.

```bash
aes-benchmark-go/
├── main.go                  # CLI test runner
├── benchmarks/
│   ├── crypto_standard.go   # crypto/aes implementation
│   ├── cryptopasta.go       # cryptopasta wrapper
│   └── data/                # test input files
├── results/
│   ├── benchmark_data.csv
│   └── charts.png
├── go.mod
└── README.md                # write-up starter
```
