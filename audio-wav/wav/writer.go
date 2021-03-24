package wav

import (
	"encoding/binary"
	"io"
)

type WavWriter struct {
	writer  io.Writer
	format  *WavFormat
	samples uint32
}

func NewWavWriter(writer io.Writer, format *WavFormat, numSamples int) *WavWriter {
	blockAlign := int(format.NumChannels * format.BitsPerSample / 8)
	byteRate := int(format.SampleRate) * blockAlign
	formatNew := &WavFormat{
		AudioFormat:   format.AudioFormat,
		NumChannels:   format.NumChannels,
		SampleRate:    format.SampleRate,
		ByteRate:      uint32(byteRate),
		BlockAlign:    uint16(blockAlign),
		BitsPerSample: format.BitsPerSample,
	}
	return &WavWriter{
		writer:  writer,
		format:  formatNew,
		samples: uint32(numSamples),
	}
}

func (w *WavWriter) WriteRIFF() error {
	dataSize := w.samples * uint32(w.format.BlockAlign)
	binary.Write(w.writer, binary.BigEndian, []byte("RIFF"))
	binary.Write(w.writer, binary.LittleEndian, uint32(36 + dataSize))
	binary.Write(w.writer, binary.BigEndian, []byte("WAVE"))
	return nil
}

func (w *WavWriter) WriteFormat() error {
	binary.Write(w.writer, binary.BigEndian, []byte("fmt "))
	binary.Write(w.writer, binary.LittleEndian, uint32(16))
	binary.Write(w.writer, binary.LittleEndian, w.format)
	return nil
}

func (w *WavWriter) WriteChunk(chunk *Chunk) error {
	binary.Write(w.writer, binary.BigEndian, chunk.Id)
	binary.Write(w.writer, binary.LittleEndian, chunk.Size)
	binary.Write(w.writer, binary.LittleEndian, chunk.Data)
	return nil
}

func (w *WavWriter) WriteData(data []byte) error {
	binary.Write(w.writer, binary.BigEndian, []byte("data"))
	binary.Write(w.writer, binary.LittleEndian, uint32(len(data)))
	binary.Write(w.writer, binary.LittleEndian, data)
	return nil
}
