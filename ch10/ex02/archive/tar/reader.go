package tar

import (
	"archive/tar"
	"os"

	"github.com/masahiro331/colang/ch10/ex02/archive"
)

type reader struct {
	r *tar.Reader
}

func (r *reader) Next() (*archive.Header, error) {
	hdr, err := r.r.Next()
	if err != nil {
		return nil, err
	}
	return &archive.Header{Name: hdr.Name, FileInfo: hdr.FileInfo()}, nil
}

func (r *reader) Read(b []byte) (n int, err error) {
	return r.r.Read(b)
}

func read(f *os.File) (archive.Reader, error) {
	return &reader{tar.NewReader(f)}, nil
}
func init() {
	archive.RegisterFormat("tar", "\x75\x73\x74\x61\x72", 0x101, read)
}
