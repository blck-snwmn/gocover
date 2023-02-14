package decrypt

import (
	"reflect"
	"testing"
)

func Test_decrypt(t *testing.T) {
	want := []byte{0xF0, 0x01, 0x12}
	if got := decrypt([]byte{0x00, 0x11, 0x22}); !reflect.DeepEqual(got, want) {
		t.Errorf("crypto() = %v, want %v", got, want)
	}
}
