package crypto

const offset = 0x10

func crypto(b []byte) []byte {
	for i := range b {
		b[i] += offset
	}
	return b
}

type P struct{}

func (P) crypto(b []byte) []byte {
	return crypto(b)
}
