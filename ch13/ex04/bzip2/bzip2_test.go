package bzip2

import (
	"bytes"
	"compress/bzip2"
	"io"
	"testing"
)

func TestBzip2(t *testing.T) {
	var compressed, uncompressed bytes.Buffer
	w, err := NewWriter(&compressed)
	if err != nil {
		t.Fatal(err)
	}

	tee := io.MultiWriter(w, &uncompressed)
	for i := 0; i < 1000000; i++ {
		io.WriteString(tee, "hogehogeho")
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	if got, want := compressed.Len(), 255; got != want {
		t.Errorf("hoge compressed to %d bytes, want %d", got, want)
	}

	var decompressed bytes.Buffer
	io.Copy(&decompressed, bzip2.NewReader(&compressed))
	if !bytes.Equal(uncompressed.Bytes(), decompressed.Bytes()) {
		t.Error("decompression different message")
	}
}
