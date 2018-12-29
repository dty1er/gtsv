package gtsv

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestInt(t *testing.T) {
	tests := []struct {
		name   string
		tsv    string
		row    int
		col    int
		result [][]int
	}{
		{
			name: "int",
			tsv: "1\t2\t3\n" +
				"4\t5\t6\n",
			row:    2,
			col:    3,
			result: [][]int{[]int{1, 2, 3}, []int{4, 5, 6}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gr := New(bytes.NewBufferString(tt.tsv))
			var rowCnt int
			var ret [][]int

			for gr.Next() {
				rowCnt++
				var line []int
				for i := 0; i < tt.col; i++ {
					line = append(line, gr.Int())
				}
				ret = append(ret, line)
			}

			if err := gr.Error(); err != nil && err != io.EOF {
				t.Fatalf("unexpected error: %s", err)
			}

			if tt.row != rowCnt {
				t.Fatalf("row check failed expected: %d, actual: %d", tt.row, rowCnt)
			}

			if !reflect.DeepEqual(tt.result, ret) {
				t.Fatalf("returned value check failed expected: %v, actual: %v", tt.result, ret)
			}
		})
	}

}
