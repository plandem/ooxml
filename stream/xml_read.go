package stream

import (
	"encoding/xml"
	"io"
)

//ReadStreamIteratorCallback is type that defines a callback that will be called for each iteration of ReadStreamIterator
type ReadStreamIteratorCallback func(decoder *xml.Decoder, start *xml.StartElement) bool

//ReadStreamIterator is type that defines a read stream iterator for all xml.StartElement elements
type ReadStreamIterator func(cb ReadStreamIteratorCallback) bool

//ReadStream is wrapper around xml.Decoder with additional functionality for reading in stream mode
type ReadStream struct {
	*xml.Decoder
}

//NewReadStream return a ReadStream for reader
func NewReadStream(reader io.Reader) *ReadStream {
	return &ReadStream{
		xml.NewDecoder(reader),
	}
}

//seekForStart returns next start element
func (rs *ReadStream) seekForStart() *xml.StartElement {
	for {
		tok, err := rs.Decoder.Token()
		if err != nil {
			break
		}

		if se, ok := tok.(xml.StartElement); ok {
			return &se
		}
	}

	return nil
}

//ElementIterator returns iterator for all start elements
func (rs *ReadStream) ElementIterator(start *xml.StartElement) (ReadStreamIterator, bool) {
	prev := start

	//get next start, if required
	if prev == nil {
		prev = rs.seekForStart()
	}

	iterator := func(cb ReadStreamIteratorCallback) bool {
		//if callback returns false, then stop
		if proceed := cb(rs.Decoder, prev); !proceed {
			prev = nil
			return false
		}

		//if there is no any start elem, then stop
		prev = rs.seekForStart()
		return prev != nil
	}

	return iterator, prev != nil
}
