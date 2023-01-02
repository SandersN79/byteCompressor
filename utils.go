package ByteCompressor

import (
	"bytes"
	"io"
	"os"
)

type ClosingBuffer struct {
	*bytes.Buffer
}

func (cb *ClosingBuffer) Close() (err error) {
	return
}

func getTestFileBuffer(filename string) (io.ReadCloser, error) {
	folderPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	folderPath += "/test/"
	dat, err := os.ReadFile(folderPath + filename)
	if err != nil {
		return nil, err
	}
	return &ClosingBuffer{bytes.NewBuffer(dat)}, nil
}

func streamToBytes(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}
