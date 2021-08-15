package xlsx_token

import (
	"archive/zip"
	"io"
)

type XlsxReader struct {
	zipReader *zip.Reader
}

// Open returns XlsxReader from *os.File
func Open(name string) (*XlsxReader, error) {
	file, err := zip.OpenReader(name)
	if err != nil {
		return nil, err
	}

	return &XlsxReader{
		zipReader: &file.Reader,
	}, nil
}

// NewReader returns XlsxReader from *io.ReaderAt interface
func NewReader(r io.ReaderAt, size int64) (*XlsxReader, error) {
	reader, err := zip.NewReader(r, size)
	if err != nil {
		return nil, err
	}

	return &XlsxReader{
		zipReader: reader,
	}, nil
}
