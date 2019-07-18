// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ooxml

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"io"
	"io/ioutil"
	"os"
	"path"
)

//StreamFileWriterFinalizer is callback that will be called to do final processing before closing stream
type StreamFileWriterFinalizer func() error

//StreamFileWriter is stream writer for *zip.File
type StreamFileWriter struct {
	*xml.Encoder
	target    interface{}
	writer    *zip.Writer
	finalizer StreamFileWriterFinalizer
}

var _ io.Closer = (*StreamFileWriter)(nil)

//NewStreamFileWriter returns a StreamFileWriter for fileName
func NewStreamFileWriter(f string, memory bool, finalizer StreamFileWriterFinalizer) (*StreamFileWriter, error) {
	var writer *zip.Writer
	var target interface{}

	if memory {
		//stream to memory
		buf := bytes.NewBuffer(nil)
		writer = zip.NewWriter(buf)
		target = buf
	} else {
		//stream to disk
		tmpFile, err := ioutil.TempFile("", path.Base(f))
		if err != nil {
			return nil, err
		}

		writer = zip.NewWriter(tmpFile)
		target = tmpFile
	}

	zipFile, err := writer.Create(f)
	if err != nil {
		return nil, err
	}

	_, err = zipFile.Write([]byte(xml.Header))
	if err != nil {
		return nil, err
	}

	enc := xml.NewEncoder(zipFile)
	return &StreamFileWriter{
		enc,
		target,
		writer,
		finalizer,
	}, nil
}

//Close previously allocated resources for writing
func (s *StreamFileWriter) Close() error {
	if s.writer != nil {
		//call finalizer, if required
		if s.finalizer != nil {
			if err := s.finalizer(); err != nil {
				return err
			}
		}

		//flush zipper
		if err := s.Flush(); err != nil {
			return err
		}

		//close writer
		var writer io.Closer
		writer, s.writer = s.writer, nil
		return writer.Close()
	}

	return nil
}

//Save current state of stream to *zip.Writer
func (s *StreamFileWriter) Save(to *zip.Writer) error {
	if err := s.Close(); err != nil {
		return err
	}

	if buf, ok := s.target.(*bytes.Buffer); ok {
		//stored in memory
		readerAt := bytes.NewReader(buf.Bytes())
		if zipReader, err := zip.NewReader(readerAt, int64(readerAt.Len())); err != nil {
			return err
		} else {
			zipFile := zipReader.File[0]
			if err := CopyZipFile(zipFile, to); err != nil {
				return err
			}
		}
	} else if file, ok := s.target.(*os.File); ok {
		//stored at disk
		if zipReader, err := zip.OpenReader(file.Name()); err != nil {
			return err
		} else {
			zipFile := zipReader.File[0]
			if err := CopyZipFile(zipFile, to); err != nil {
				return err
			}

			if err = zipReader.Close(); err != nil {
				return err
			}

			os.Remove(file.Name())
		}
	}

	s.target = nil
	return nil
}
