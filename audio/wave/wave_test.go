package wave

import (
	"fmt"
	"os"
)

func Example() {
	const tmpfile = "./tmp.wav"
	data := make([]byte, 1000)
	outfile, err := os.Create(tmpfile)
	if err != nil {
		panic(err)
	}

	formatNew := &WavFormat{
		AudioFormat:   1,
		NumChannels:   1,
		SampleRate:    16000,
		BitsPerSample: 16,
	}
	numSamples := len(data) / 2
	wavWriter := NewWavWriter(outfile, formatNew, numSamples)
	wavWriter.WriteRIFF()
	wavWriter.WriteFormat()
	wavWriter.WriteData(data)

	outfile.Close()

	infile, err := os.Open(tmpfile)
	if err != nil {
		panic(err)
	}
	wavReader := NewWavReader(infile)
	_, err = wavReader.ReadRIFF()
	if err != nil {
		panic(err)
	}
	format, err := wavReader.ReadFormat()
	if err != nil {
		panic(err)
	}

	fmt.Printf("AudioFormat: %v\n", format.AudioFormat)
	fmt.Printf("NumChannels: %v\n", format.NumChannels)
	fmt.Printf("SampleRate: %v\n", format.SampleRate)
	fmt.Printf("ByteRate: %v\n", format.ByteRate)
	fmt.Printf("BlockAlign: %v\n", format.BlockAlign)
	fmt.Printf("BitsPerSample: %v\n", format.BitsPerSample)

	chunk, err := wavReader.ReadChunk("data")
	if err != nil {
		panic(err)
	}
	infile.Close()
	os.Remove(tmpfile)
	fmt.Printf("DataLength: %v\n", len(chunk))
	// Output:
	// AudioFormat: 1
	// NumChannels: 1
	// SampleRate: 16000
	// ByteRate: 32000
	// BlockAlign: 2
	// BitsPerSample: 16
	// DataLength: 1000
}
