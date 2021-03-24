package main

import (
	"fmt"
	"os"

	"audio-wav/wav"
)

func main() {
	if len(os.Args) != 2 && len(os.Args) != 3 {
		fmt.Printf("Usage: %s <in.wav> [out.wav]\n", os.Args[0])
		os.Exit(-1)
	}

	infile, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer infile.Close()

	wavReader := wav.NewWavReader(infile)
	_, err = wavReader.ReadRIFF()
	if err != nil {
		panic(err)
	}
	format, err := wavReader.ReadFormat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("AudioFormat:\t%v\n", format.AudioFormat)
	fmt.Printf("NumChannels:\t%v\n", format.NumChannels)
	fmt.Printf("SampleRate:\t%v\n", format.SampleRate)
	fmt.Printf("ByteRate:\t%v\n", format.ByteRate)
	fmt.Printf("BlockAlign:\t%v\n", format.BlockAlign)
	fmt.Printf("BitsPerSample:\t%v\n", format.BitsPerSample)

	data, err := wavReader.ReadChunk("data")
	if err != nil {
		panic(err)
	}

	if len(os.Args) < 3 {
		os.Exit(0)
	}

	outfile, err := os.Create(os.Args[2])
	defer outfile.Close()
	if err != nil {
		panic(err)
	}
	formatNew := &wav.WavFormat{
		AudioFormat:   1,
		NumChannels:   1,
		SampleRate:    16000,
		BitsPerSample: 16,
	}
	numSamples := len(data) / 2
	wavWriter := wav.NewWavWriter(outfile, formatNew, numSamples)
	wavWriter.WriteRIFF()
	wavWriter.WriteFormat()
	wavWriter.WriteData(data)
}
