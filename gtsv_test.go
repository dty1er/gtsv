package gtsv

import (
	"bytes"
	"io"
	"testing"
)

func Test(t *testing.T) {
	bs := bytes.NewBufferString(
		"1\t2\t3\n" +
			"4\t5\t6\n")

	gr := New(bs)

	for gr.Next() {
		c1 := gr.Int()
		c2 := gr.Int()
		c3 := gr.Int()
		t.Logf("c1=%d, c2=%d, c3=%d\n", c1, c2, c3)
	}
	if err := gr.Error(); err != nil && err != io.EOF {
		t.Errorf("unexpected error: %s", err)
	}
}
