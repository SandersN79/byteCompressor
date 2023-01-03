# Bytes Compressor
Go Module developed to concurrently GZIP compress any io.ReadCloser in-memory.

## Setup
```bash
$ go get github.com/SandersN79/ByteCompressorSF
```

## Test
```bash
$ go test
```

## Usage
```go
package main

import (
	"github.com/SandersN79/ByteCompressorSF"
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	r := io.NopCloser(strings.NewReader("Hello, world!")) // r type is io.ReadCloser
	compressor := byteCompressor.NewBytesCompressor(r)
	buf := new(bytes.Buffer)
	buf.ReadFrom(compressor)
	fmt.Println("Output: ", buf.Bytes())
}
```