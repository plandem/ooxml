package ooxml

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
)

//StreamWriterCallback is type that defines a callback that will be called to get postponed updates
type StreamWriterCallback func(encoder *xml.Encoder)

//StreamWriter is a generic stream writer
type StreamWriter struct {
	*xml.Encoder
}

//StreamFileWriter is stream writer for *zip.File
type StreamFileWriter struct {
	*StreamWriter
	buff     *bytes.Buffer
	cb       StreamWriterCallback
	fileName string
}

//NewStreamFileWriter returns a StreamFileWriter for fileName f
func NewStreamFileWriter(f string, cb StreamWriterCallback) *StreamFileWriter {
	buff := &bytes.Buffer{}
	enc := xml.NewEncoder(buff)

	return &StreamFileWriter{
		&StreamWriter{enc},
		buff,
		cb,
		f,
	}
}

//Save saves current state of stream and postponed changes to *zip.Writer
func (s *StreamFileWriter) Save(to *zip.Writer) error {
	writer, err := to.Create(s.fileName)
	if err != nil {
		return err
	}

	//save current state of stream
	writer.Write(s.buff.Bytes())

	//if there is a callback to get postponed changes, then stream it directly to zip writer
	if s.cb != nil {
		finalizeEnc := xml.NewEncoder(writer)
		s.cb(finalizeEnc)
	}

	return nil
}
