package wav

type RIFF struct {
	Id             [4]byte
	Size           uint32
	Type           [4]byte
	Chunks         []*Chunk
}

type Chunk struct {
	Id             [4]byte
	Size           uint32
	Data           []byte
}

type WavFormat struct {
	AudioFormat    uint16
	NumChannels    uint16
	SampleRate     uint32
	ByteRate       uint32
	BlockAlign     uint16
	BitsPerSample  uint16
}

