package byteCompressor

import (
	"testing"
)

func Test_BytesCompressor(t *testing.T) {
	// Defining our test slice. Each unit test should have the following properties:
	tests := []struct {
		name     string // The name of the test
		want     string // What out instance we want our function to return.
		wantErr  bool   // whether we want an error.
		testFile string // The input of the test
	}{
		// Here we're declaring each unit test input and output data as defined before
		{
			"success",
			"test1.txt.gz",
			false,
			"test1.txt",
		},
	}
	// Iterating over the previous test slice
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tBuffer, err := getTestFileBuffer(tt.testFile)
			defer tBuffer.Close()
			tCompressor := NewBytesCompressor(tBuffer)
			cBytes := streamToBytes(tCompressor)
			if (err != nil) != tt.wantErr { // Checking the error
				t.Errorf("NewBytesCompressor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			oBuffer, _ := getTestFileBuffer(tt.want)
			oBytes := streamToBytes(oBuffer)
			if len(oBytes) != len(cBytes) {
				t.Errorf("NewBytesCompressor() = %v, want %v", len(cBytes), len(oBytes))
			}
		})
	}
}
