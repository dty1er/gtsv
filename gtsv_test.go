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

			err := gr.Error()
			if (err != nil) != tt.hasError && err != io.EOF {
				t.Fatalf("error is not io.EOF but %s", err)
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

func TestUint(t *testing.T) {
	tests := []struct {
		name     string
		tsv      string
		row      int
		col      int
		result   [][]uint
		hasError bool
	}{
		{
			name: "uint",
			tsv: "1\t2\t3\n" +
				"4\t5\t6\n",
			row:    2,
			col:    3,
			result: [][]uint{[]uint{1, 2, 3}, []uint{4, 5, 6}},
		},
		{
			name: "contains negative number",
			tsv: "1\t2\t3\n" +
				"4\t5\t-1\n",
			row:      2,
			col:      3,
			result:   [][]uint{[]uint{1, 2, 3}, []uint{4, 5, 0}},
			hasError: true,
		},
		{
			name: "contains string",
			tsv: "1\t2\t3\n" +
				"4\t5\ta\n",
			row:      2,
			col:      3,
			result:   [][]uint{[]uint{1, 2, 3}, []uint{4, 5, 0}},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gr := New(bytes.NewBufferString(tt.tsv))
			var rowCnt int
			var ret [][]uint

			for gr.Next() {
				rowCnt++
				var line []uint
				for i := 0; i < tt.col; i++ {
					line = append(line, gr.Uint())
				}
				ret = append(ret, line)
			}
			err := gr.Error()
			if (err != nil) != tt.hasError && err != io.EOF {
				t.Fatalf("error is not io.EOF but %s", err)
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

func TestInt8(t *testing.T) {
	tests := []struct {
		name     string
		tsv      string
		row      int
		col      int
		result   [][]int8
		hasError bool
	}{
		{
			name: "int8",
			tsv: "1\t2\t3\n" +
				"4\t5\t6\n",
			row:    2,
			col:    3,
			result: [][]int8{[]int8{1, 2, 3}, []int8{4, 5, 6}},
		},
		{
			name: "contains too large number",
			tsv: "1\t2\t3\n" +
				"4\t5\t128\n",
			row:      2,
			col:      3,
			result:   [][]int8{[]int8{1, 2, 3}, []int8{4, 5, 0}},
			hasError: true,
		},
		{
			name: "contains string",
			tsv: "1\t2\t3\n" +
				"4\t5\ta\n",
			row:      2,
			col:      3,
			result:   [][]int8{[]int8{1, 2, 3}, []int8{4, 5, 0}},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gr := New(bytes.NewBufferString(tt.tsv))
			var rowCnt int
			var ret [][]int8

			for gr.Next() {
				rowCnt++
				var line []int8
				for i := 0; i < tt.col; i++ {
					line = append(line, gr.Int8())
				}
				ret = append(ret, line)
			}
			err := gr.Error()
			if (err != nil) != tt.hasError && err != io.EOF {
				t.Fatalf("error is not io.EOF but %s", err)
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

func TestUint8(t *testing.T) {
	tests := []struct {
		name     string
		tsv      string
		row      int
		col      int
		result   [][]uint8
		hasError bool
	}{
		{
			name: "uint8",
			tsv: "1\t2\t3\n" +
				"4\t5\t6\n",
			row:    2,
			col:    3,
			result: [][]uint8{[]uint8{1, 2, 3}, []uint8{4, 5, 6}},
		},
		{
			name: "contains negative number",
			tsv: "1\t2\t3\n" +
				"4\t5\t-1\n",
			row:      2,
			col:      3,
			result:   [][]uint8{[]uint8{1, 2, 3}, []uint8{4, 5, 0}},
			hasError: true,
		},
		{
			name: "contains too large number",
			tsv: "1\t2\t3\n" +
				"4\t5\t256\n",
			row:      2,
			col:      3,
			result:   [][]uint8{[]uint8{1, 2, 3}, []uint8{4, 5, 0}},
			hasError: true,
		},
		{
			name: "contains string",
			tsv: "1\t2\t3\n" +
				"4\t5\ta\n",
			row:      2,
			col:      3,
			result:   [][]uint8{[]uint8{1, 2, 3}, []uint8{4, 5, 0}},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gr := New(bytes.NewBufferString(tt.tsv))
			var rowCnt int
			var ret [][]uint8

			for gr.Next() {
				rowCnt++
				var line []uint8
				for i := 0; i < tt.col; i++ {
					line = append(line, gr.Uint8())
				}
				ret = append(ret, line)
			}
			err := gr.Error()
			if (err != nil) != tt.hasError && err != io.EOF {
				t.Fatalf("error is not io.EOF but %s", err)
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

func TestInt16(t *testing.T) {
	tests := []struct {
		name     string
		tsv      string
		row      int
		col      int
		result   [][]int16
		hasError bool
	}{
		{
			name: "int16",
			tsv: "1\t2\t3\n" +
				"4\t5\t6\n",
			row:    2,
			col:    3,
			result: [][]int16{[]int16{1, 2, 3}, []int16{4, 5, 6}},
		},
		{
			name: "contains too large number",
			tsv: "1\t2\t3\n" +
				"4\t5\t32768\n",
			row:      2,
			col:      3,
			result:   [][]int16{[]int16{1, 2, 3}, []int16{4, 5, 0}},
			hasError: true,
		},
		{
			name: "contains string",
			tsv: "1\t2\t3\n" +
				"4\t5\t65536\n",
			row:      2,
			col:      3,
			result:   [][]int16{[]int16{1, 2, 3}, []int16{4, 5, 0}},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gr := New(bytes.NewBufferString(tt.tsv))
			var rowCnt int
			var ret [][]int16

			for gr.Next() {
				rowCnt++
				var line []int16
				for i := 0; i < tt.col; i++ {
					line = append(line, gr.Int16())
				}
				ret = append(ret, line)
			}
			err := gr.Error()
			if (err != nil) != tt.hasError && err != io.EOF {
				t.Fatalf("error is not io.EOF but %s", err)
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

func TestUint16(t *testing.T) {
	tests := []struct {
		name     string
		tsv      string
		row      int
		col      int
		result   [][]uint16
		hasError bool
	}{
		{
			name: "uint16",
			tsv: "1\t2\t3\n" +
				"4\t5\t6\n",
			row:    2,
			col:    3,
			result: [][]uint16{[]uint16{1, 2, 3}, []uint16{4, 5, 6}},
		},
		{
			name: "contains negative number",
			tsv: "1\t2\t3\n" +
				"4\t5\t-1\n",
			row:      2,
			col:      3,
			result:   [][]uint16{[]uint16{1, 2, 3}, []uint16{4, 5, 0}},
			hasError: true,
		},
		{
			name: "contains too large number",
			tsv: "1\t2\t3\n" +
				"4\t5\t65536\n",
			row:      2,
			col:      3,
			result:   [][]uint16{[]uint16{1, 2, 3}, []uint16{4, 5, 0}},
			hasError: true,
		},
		{
			name: "contains string",
			tsv: "1\t2\t3\n" +
				"4\t5\ta\n",
			row:      2,
			col:      3,
			result:   [][]uint16{[]uint16{1, 2, 3}, []uint16{4, 5, 0}},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gr := New(bytes.NewBufferString(tt.tsv))
			var rowCnt int
			var ret [][]uint16

			for gr.Next() {
				rowCnt++
				var line []uint16
				for i := 0; i < tt.col; i++ {
					line = append(line, gr.Uint16())
				}
				ret = append(ret, line)
			}
			err := gr.Error()
			if (err != nil) != tt.hasError && err != io.EOF {
				t.Fatalf("error is not io.EOF but %s", err)
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

func TestInt32(t *testing.T) {
	tests := []struct {
		name     string
		tsv      string
		row      int
		col      int
		result   [][]int32
		hasError bool
	}{
		{
			name: "int32",
			tsv: "1\t2\t3\n" +
				"4\t5\t6\n",
			row:    2,
			col:    3,
			result: [][]int32{[]int32{1, 2, 3}, []int32{4, 5, 6}},
		},
		{
			name: "contains too large number",
			tsv: "1\t2\t3\n" +
				"4\t5\t2147483648\n",
			row:      2,
			col:      3,
			result:   [][]int32{[]int32{1, 2, 3}, []int32{4, 5, 0}},
			hasError: true,
		},
		{
			name: "contains string",
			tsv: "1\t2\t3\n" +
				"4\t5\ta\n",
			row:      2,
			col:      3,
			result:   [][]int32{[]int32{1, 2, 3}, []int32{4, 5, 0}},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gr := New(bytes.NewBufferString(tt.tsv))
			var rowCnt int
			var ret [][]int32

			for gr.Next() {
				rowCnt++
				var line []int32
				for i := 0; i < tt.col; i++ {
					line = append(line, gr.Int32())
				}
				ret = append(ret, line)
			}
			err := gr.Error()
			if (err != nil) != tt.hasError && err != io.EOF {
				t.Fatalf("error is not io.EOF but %s", err)
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

func TestUint32(t *testing.T) {
	tests := []struct {
		name     string
		tsv      string
		row      int
		col      int
		result   [][]uint32
		hasError bool
	}{
		{
			name: "uint32",
			tsv: "1\t2\t3\n" +
				"4\t5\t6\n",
			row:    2,
			col:    3,
			result: [][]uint32{[]uint32{1, 2, 3}, []uint32{4, 5, 6}},
		},
		{
			name: "contains negative number",
			tsv: "1\t2\t3\n" +
				"4\t5\t-1\n",
			row:      2,
			col:      3,
			result:   [][]uint32{[]uint32{1, 2, 3}, []uint32{4, 5, 0}},
			hasError: true,
		},
		{
			name: "contains too large number",
			tsv: "1\t2\t3\n" +
				"4\t5\t4294967296\n",
			row:      2,
			col:      3,
			result:   [][]uint32{[]uint32{1, 2, 3}, []uint32{4, 5, 0}},
			hasError: true,
		},
		{
			name: "contains string",
			tsv: "1\t2\t3\n" +
				"4\t5\ta\n",
			row:      2,
			col:      3,
			result:   [][]uint32{[]uint32{1, 2, 3}, []uint32{4, 5, 0}},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gr := New(bytes.NewBufferString(tt.tsv))
			var rowCnt int
			var ret [][]uint32

			for gr.Next() {
				rowCnt++
				var line []uint32
				for i := 0; i < tt.col; i++ {
					line = append(line, gr.Uint32())
				}
				ret = append(ret, line)
			}
			err := gr.Error()
			if (err != nil) != tt.hasError && err != io.EOF {
				t.Fatalf("error is not io.EOF but %s", err)
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

func TestInt64(t *testing.T) {
	tests := []struct {
		name     string
		tsv      string
		row      int
		col      int
		result   [][]int64
		hasError bool
	}{
		{
			name: "int64",
			tsv: "1\t2\t3\n" +
				"4\t5\t6\n",
			row:    2,
			col:    3,
			result: [][]int64{[]int64{1, 2, 3}, []int64{4, 5, 6}},
		},
		{
			name: "contains too large number",
			tsv: "1\t2\t3\n" +
				"4\t5\t9223372036854775808\n",
			row:      2,
			col:      3,
			result:   [][]int64{[]int64{1, 2, 3}, []int64{4, 5, 0}},
			hasError: true,
		},
		{
			name: "contains string",
			tsv: "1\t2\t3\n" +
				"4\t5\ta\n",
			row:      2,
			col:      3,
			result:   [][]int64{[]int64{1, 2, 3}, []int64{4, 5, 0}},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gr := New(bytes.NewBufferString(tt.tsv))
			var rowCnt int
			var ret [][]int64

			for gr.Next() {
				rowCnt++
				var line []int64
				for i := 0; i < tt.col; i++ {
					line = append(line, gr.Int64())
				}
				ret = append(ret, line)
			}
			err := gr.Error()
			if (err != nil) != tt.hasError && err != io.EOF {
				t.Fatalf("error is not io.EOF but %s", err)
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

func TestUint64(t *testing.T) {
	tests := []struct {
		name     string
		tsv      string
		row      int
		col      int
		result   [][]uint64
		hasError bool
	}{
		{
			name: "uint64",
			tsv: "1\t2\t3\n" +
				"4\t5\t6\n",
			row:    2,
			col:    3,
			result: [][]uint64{[]uint64{1, 2, 3}, []uint64{4, 5, 6}},
		},
		{
			name: "contains negative number",
			tsv: "1\t2\t3\n" +
				"4\t5\t-1\n",
			row:      2,
			col:      3,
			result:   [][]uint64{[]uint64{1, 2, 3}, []uint64{4, 5, 0}},
			hasError: true,
		},
		{
			name: "contains too large number",
			tsv: "1\t2\t3\n" +
				"4\t5\t18446744073709551616\n",
			row:      2,
			col:      3,
			result:   [][]uint64{[]uint64{1, 2, 3}, []uint64{4, 5, 0}},
			hasError: true,
		},
		{
			name: "contains string",
			tsv: "1\t2\t3\n" +
				"4\t5\ta\n",
			row:      2,
			col:      3,
			result:   [][]uint64{[]uint64{1, 2, 3}, []uint64{4, 5, 0}},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gr := New(bytes.NewBufferString(tt.tsv))
			var rowCnt int
			var ret [][]uint64

			for gr.Next() {
				rowCnt++
				var line []uint64
				for i := 0; i < tt.col; i++ {
					line = append(line, gr.Uint64())
				}
				ret = append(ret, line)
			}
			err := gr.Error()
			if (err != nil) != tt.hasError && err != io.EOF {
				t.Fatalf("error is not io.EOF but %s", err)
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

func TestString(t *testing.T) {
	tests := []struct {
		name   string
		tsv    string
		row    int
		col    int
		result [][]string
	}{
		{
			name: "string",
			tsv: "aaa\tbbb\tccc\n" +
				"ddd\teee\tfff\n",
			row:    2,
			col:    3,
			result: [][]string{[]string{"aaa", "bbb", "ccc"}, []string{"ddd", "eee", "fff"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gr := New(bytes.NewBufferString(tt.tsv))
			var rowCnt int
			var ret [][]string

			for gr.Next() {
				rowCnt++
				var line []string
				for i := 0; i < tt.col; i++ {
					line = append(line, gr.String())
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
