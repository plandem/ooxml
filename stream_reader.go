package ooxml

import (
	"archive/zip"
	"encoding/xml"
	"io"
)

//StreamReaderCallback is type that defines a callback that will be called for each iteration of StreamReaderIterator
type StreamReaderCallback func(decoder *xml.Decoder, start *xml.StartElement) bool

//StreamReaderIterator is type that defines a read stream iterator for all xml.StartElement elements
type StreamReaderIterator func(cb StreamReaderCallback) bool

//StreamReader is a generic stream reader
type StreamReader struct {
	*xml.Decoder
}

//StreamFileReader is stream reader for *zip.File
type StreamFileReader struct {
	*StreamReader
	f  *zip.File
	rc io.ReadCloser
}

var _ io.Closer = (*StreamFileReader)(nil)

//seekForStart returns next start element
func (sr *StreamReader) seekForStart() *xml.StartElement {
	for {
		tok, err := sr.Decoder.Token()
		if err != nil {
			break
		}

		if se, ok := tok.(xml.StartElement); ok {
			return &se
		}
	}

	return nil
}

//StartIterator returns iterator for all start elements
func (sr *StreamReader) StartIterator(start *xml.StartElement) (StreamReaderIterator, bool) {
	prev := start

	//get next start, if required
	if prev == nil {
		prev = sr.seekForStart()
	}

	iterator := func(cb StreamReaderCallback) bool {
		//if callback returns false, then stop
		if proceed := cb(sr.Decoder, prev); !proceed {
			prev = nil
			return false
		}

		//if there is no any start elem, then stop
		prev = sr.seekForStart()
		return prev != nil
	}

	return iterator, prev != nil
}

//NewStreamFileReader returns a StreamFileReader for *zip.File f
func NewStreamFileReader(f *zip.File) (*StreamFileReader, error) {
	rc, err := f.Open()
	if err != nil {
		return nil, err
	}

	return &StreamFileReader{
		&StreamReader{xml.NewDecoder(rc)},
		f,
		rc,
	}, nil
}

//Close closes previously opened file for reading
func (s *StreamFileReader) Close() error {
	if s.rc != nil {
		var rc io.ReadCloser
		rc, s.rc = s.rc, nil
		return rc.Close()
	}

	return nil
}
