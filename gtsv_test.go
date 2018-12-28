package gtsv

import (
	"bytes"
	"testing"
)

func Test(t *testing.T) {
	bs := bytes.NewBufferString(
		"1\t2\t3\n" +
			"4\t5\t6\n")

	gr := New(bs)

	for gr.Next() {
		t.Log("0----")
		c1 := gr.Int()
		t.Log("1----")
		c2 := gr.Int()
		t.Log("2----")
		t.Logf("c1=%d, c2=%d\n", c1, c2)
	}
	if err := gr.Error(); err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}
