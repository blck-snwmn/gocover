package crypto

import (
	"reflect"
	"testing"
)

func Test_crypto(t *testing.T) {
	want := []byte{0x10, 0x21, 0x32}
	if got := crypto([]byte{0x00, 0x11, 0x22}); !reflect.DeepEqual(got, want) {
		t.Errorf("crypto() = %v, want %v", got, want)
	}
}
