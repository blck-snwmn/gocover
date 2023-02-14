package decrypt

const offset = 0x10

func decrypt(b []byte) []byte {
	for i := range b {
		b[i] -= offset
	}
	return b
}

type P struct{}

func (P) Decrypt(b []byte) []byte {
	return decrypt(b)
}
