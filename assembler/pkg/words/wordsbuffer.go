package words

import "io"

type Int16SliceReader struct {
	data []int16
	pos  int
}

func NewInt16SliceReader(data []int16) *Int16SliceReader {
	return &Int16SliceReader{data: data, pos: 0}
}

func (r *Int16SliceReader) Read(p []byte) (n int, err error) {
	if r.pos >= len(r.data)*2 {
		return 0, io.EOF
	}
	for i := range p {
		if r.pos >= len(r.data)*2 {
			break
		}
		p[i] = byte(r.data[r.pos/2] >> ((1 - r.pos%2) * 8))
		r.pos++
		n++
	}
	return
}
