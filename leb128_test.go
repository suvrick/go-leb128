package leb128

import (
	"bytes"
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {

	buf := new(bytes.Buffer)
	err := Write(buf, "www")
	fmt.Printf("%x\nerr: %v\n", buf.Bytes(), err)

	v, err := Read(buf, 'S')
	fmt.Printf("%v\nerr: %v\n", v.(string), err)
}

func TestWrite(t *testing.T) {

	type args struct {
		value int64
		want  []uint8
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"zero case",
			args{
				0,
				[]byte{0},
			},
		},
		{
			"255 case",
			args{
				255,
				[]byte{0xff, 0x01},
			},
		},
		{
			"-9019283812387 case",
			args{
				-9019283812387,
				[]byte{0xdd, 0x9f, 0xab, 0xc6, 0xc0, 0xf9, 0x7d},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			Write(buf, tt.args.value)
			r := buf.Bytes()
			if !bytes.Equal(r, tt.args.want) {
				t.Errorf("%s, want: %x, got: %x\n", tt.name, tt.args.want, r)
			}
		})
	}
}
