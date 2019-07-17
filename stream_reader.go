// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ooxml

import (
	"archive/zip"
	"encoding/xml"
	"io"
)

//StreamFileReader is stream reader for *zip.File
type StreamFileReader struct {
	*xml.Decoder
	source *zip.File
	reader io.ReadCloser
}

var _ io.Closer = (*StreamFileReader)(nil)

//NewStreamFileReader returns a StreamFileReader for *zip.File f
func NewStreamFileReader(f *zip.File) (*StreamFileReader, error) {
	reader, err := f.Open()
	if err != nil {
		return nil, err
	}

	return &StreamFileReader{
		xml.NewDecoder(reader),
		f,
		reader,
	}, nil
}

//Close closes previously opened file for reading
func (s *StreamFileReader) Close() error {
	if s.reader != nil {
		var reader io.ReadCloser
		reader, s.reader = s.reader, nil
		return reader.Close()
	}

	return nil
}
