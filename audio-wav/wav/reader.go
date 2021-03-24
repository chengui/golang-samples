package wav

import (
	"io"
	"fmt"
	"bytes"
	"encoding/binary"
)

type WavReader struct {
	reader io.Reader
	*RIFF
}

func NewWavReader(reader io.Reader) *WavReader {
	return &WavReader{
		reader: reader,
	}
}

func (w *WavReader) ReadRIFF() (*RIFF, error) {
	riff := RIFF{}
	binary.Read(w.reader, binary.BigEndian, &riff.Id)
	if string(riff.Id[:]) != "RIFF" {
		return nil, fmt.Errorf("invalid Riff format")
	}
	binary.Read(w.reader, binary.LittleEndian, &riff.Size)
	binary.Read(w.reader, binary.BigEndian, &riff.Type)
	if string(riff.Type[:]) != "WAVE" {
		return nil, fmt.Errorf("invalid wave format")
	}
	offset := uint32(4)
	for offset < riff.Size {
		chunk := Chunk{}
		binary.Read(w.reader, binary.BigEndian, &chunk.Id)
		binary.Read(w.reader, binary.LittleEndian, &chunk.Size)
		chunk.Data = make([]byte, chunk.Size)
		binary.Read(w.reader, binary.LittleEndian, &chunk.Data)
		riff.Chunks = append(riff.Chunks, &chunk)
		offset += chunk.Size + 8
	}
	w.RIFF = &riff
	return &riff, nil
}

func (w *WavReader) ReadFormat() (*WavFormat, error) {
	if w.RIFF == nil {
		RIFF, err := w.ReadRIFF()
		if err != nil {
			return nil, err
		}
		w.RIFF = RIFF
	}

	chunk := new(Chunk)
	for _, ch := range w.RIFF.Chunks {
		if string(ch.Id[:]) == string("fmt ") {
			chunk = ch
			break
		}
	}

	frmt := &WavFormat{}
	buf := bytes.NewReader(chunk.Data)
	binary.Read(buf, binary.LittleEndian, frmt)
	return frmt, nil
}

func (w *WavReader) ReadChunk(id string) ([]byte, error) {
	if w.RIFF == nil {
		RIFF, err := w.ReadRIFF()
		if err != nil {
			return nil, err
		}
		w.RIFF = RIFF
	}

	chunk := new(Chunk)
	for _, ch := range w.RIFF.Chunks {
		if string(ch.Id[:]) == id {
			chunk = ch
			break
		}
	}
	return chunk.Data, nil
}
