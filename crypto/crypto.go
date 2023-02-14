package crypto

const offset = 0x10

func crypto(b []byte) []byte {
	for i := range b {
		b[i] += offset
	}
	return b
}

type P struct{}

// Crypto crypto
func (P) Crypto(b []byte) []byte {
	return crypto(b)
}
