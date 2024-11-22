package main

type OurByteBuffer struct {
	data      []byte
	readIndex int
}

func NewBuffer(data []byte) *OurByteBuffer {
	return &OurByteBuffer{
		data:      data,
		readIndex: 0,
	}
}

func (b *OurByteBuffer) Bytes() []byte {
	return b.data
}

func (b *OurByteBuffer) Write(data []byte) (n int) {
	b.data = append(b.data, data...)

	return len(data)
}

func (b *OurByteBuffer) Read(buf []byte) (n int, err error) {
	x := copy(buf, b.data[b.readIndex:])
	b.readIndex += x

	return x, nil
}
