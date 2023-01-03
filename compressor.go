package byteCompressor

import (
	"compress/gzip"
	"io"
)

// BytesCompressor is an io.Reader that provides compressed bytes
type BytesCompressor interface {
	io.Reader
}

// NewBytesCompressor takes a source of uncompressed bytes and concurrently compresses into a byteCompressor
func NewBytesCompressor(rc io.ReadCloser) BytesCompressor {
	rp, wp := io.Pipe() //read process, write process
	go func() {         //go func is asynchronous, will continue past it while its running
		defer wp.Close()
		g, err := gzip.NewWriterLevel(wp, gzip.BestCompression) // Set compression level as needed, i.e. gzip.BestSpeed
		defer g.Close()
		if err != nil {
			wp.CloseWithError(err)
		}
		io.Copy(g, rc) //copies rc into g
	}()
	return rp
}
