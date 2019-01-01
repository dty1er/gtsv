package gtsv

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"strconv"
	"unsafe"
)

// Reader reads TSV.
type Reader struct {
	reader   io.Reader
	readBuff []byte // temporary buffer which stores line
	colBuff  []byte // buffer which stores current column
	readErr  error
	col      int
	row      int
	err      error

	buff [6 << 10]byte // large enough
}

// New returnds new TSV reader from io.Reader
func New(r io.Reader) *Reader {
	return &Reader{reader: r, err: nil}
}

// Error returns error
func (gr *Reader) Error() error {
	return gr.err
}

// Next returns true when next row exists.
func (gr *Reader) Next() bool {
	if gr.err != nil {
		return false
	}

	gr.row++
	for {
		if len(gr.readBuff) <= 0 {
			if gr.readErr != nil {
				gr.err = gr.readErr
				if gr.err != io.EOF {
					gr.err = fmt.Errorf("cannot read row #%d: %s", gr.row, gr.err)
				}
				// if EOF, gr.err is io.EOF now
				return false
			}
			n, err := gr.reader.Read(gr.buff[:]) // first, read and get some bytes and store to buffer
			gr.readBuff = gr.buff[:n]
			gr.readErr = err
		}

		n := bytes.IndexByte(gr.readBuff, '\n') // read from buffer
		if n >= 0 {
			// next row found
			read := gr.readBuff[:n]
			gr.readBuff = gr.readBuff[n+1:]
			gr.colBuff = read
			return true
		}
	}
	// implement if cannot find \n
}

// Int returns next int column
func (gr *Reader) Int() int {
	if gr.err != nil {
		return 0
	}
	b, err := gr.nextColumn()
	if err != nil {
		gr.err = fmt.Errorf("cannot read `int`: %s", err)
		return 0
	}

	s := bytesToString(b)
	n, err := strconv.Atoi(s)
	if err == nil {
		return n
	}
	gr.err = fmt.Errorf("cannot parse %s as `int`: %s", s, err)
	return 0
}

// Uint returns next uint column
func (gr *Reader) Uint() uint {
	if gr.err != nil {
		return 0
	}
	b, err := gr.nextColumn()
	if err != nil {
		gr.err = fmt.Errorf("cannot read `uint`: %s", err)
		return 0
	}

	s := bytesToString(b)
	n, err := strconv.Atoi(s)
	if err == nil && 0 <= n {
		return uint(n)
	}

	gr.err = fmt.Errorf("cannot parse %s as `uint`: %s", s, err)
	return 0
}

// Int8 returns next int8 column
func (gr *Reader) Int8() int8 {
	if gr.err != nil {
		return 0
	}
	b, err := gr.nextColumn()
	if err != nil {
		gr.err = fmt.Errorf("cannot read `int8`: %s", err)
		return 0
	}

	s := bytesToString(b)
	n, err := strconv.Atoi(s)
	if err == nil && math.MinInt8 <= n && n <= math.MaxInt8 {
		return int8(n)
	}
	gr.err = fmt.Errorf("cannot parse %s as `int8`: %s", s, err)
	return 0
}

// Uint8 returns next uint8 column
func (gr *Reader) Uint8() uint8 {
	if gr.err != nil {
		return 0
	}
	b, err := gr.nextColumn()
	if err != nil {
		gr.err = fmt.Errorf("cannot read `uint8`: %s", err)
		return 0
	}

	s := bytesToString(b)
	n, err := strconv.Atoi(s)
	if err == nil && 0 <= n && n <= math.MaxUint8 {
		return uint8(n)
	}
	gr.err = fmt.Errorf("cannot parse %s as `uint8`: %s", s, err)
	return 0
}

// Int16 returns next int16 column
func (gr *Reader) Int16() int16 {
	if gr.err != nil {
		return 0
	}
	b, err := gr.nextColumn()
	if err != nil {
		gr.err = fmt.Errorf("cannot read `int16`: %s", err)
		return 0
	}

	s := bytesToString(b)
	n, err := strconv.Atoi(s)
	if err == nil && math.MinInt16 <= n && n <= math.MaxInt16 {
		return int16(n)
	}
	gr.err = fmt.Errorf("cannot parse %s as `int16`: %s", s, err)
	return 0
}

// Uint16 returns next uint16 column
func (gr *Reader) Uint16() uint16 {
	if gr.err != nil {
		return 0
	}
	b, err := gr.nextColumn()
	if err != nil {
		gr.err = fmt.Errorf("cannot read `uint16`: %s", err)
		return 0
	}

	s := bytesToString(b)
	n, err := strconv.Atoi(s)
	if err == nil && 0 <= n && n <= math.MaxUint16 {
		return uint16(n)
	}
	gr.err = fmt.Errorf("cannot parse %s as `uint16`: %s", s, err)
	return 0
}

// Int32 returns next int32 column
func (gr *Reader) Int32() int32 {
	if gr.err != nil {
		return 0
	}
	b, err := gr.nextColumn()
	if err != nil {
		gr.err = fmt.Errorf("cannot read `int32`: %s", err)
		return 0
	}

	s := bytesToString(b)
	n, err := strconv.Atoi(s)
	if err == nil && math.MinInt32 <= n && n <= math.MaxInt32 {
		return int32(n)
	}

	gr.err = fmt.Errorf("cannot parse %s as `int32`: %s", s, err)
	return 0
}

// Uint32 returns next uint32 column
func (gr *Reader) Uint32() uint32 {
	if gr.err != nil {
		return 0
	}
	b, err := gr.nextColumn()
	if err != nil {
		gr.err = fmt.Errorf("cannot read `uint32`: %s", err)
		return 0
	}

	s := bytesToString(b)
	n, err := strconv.Atoi(s)
	if err == nil && 0 <= n && n <= math.MaxUint32 {
		return uint32(n)
	}
	gr.err = fmt.Errorf("cannot parse %s as `uint32`: %s", s, err)
	return 0
}

// Int64 returns next int64 column
func (gr *Reader) Int64() int64 {
	if gr.err != nil {
		return 0
	}
	b, err := gr.nextColumn()
	if err != nil {
		gr.err = fmt.Errorf("cannot read `int64`: %s", err)
		return 0
	}

	s := bytesToString(b)
	n, err := strconv.ParseInt(s, 10, 64)
	if err == nil && math.MinInt64 <= n && n <= math.MaxInt64 {
		return int64(n)
	}
	gr.err = fmt.Errorf("cannot parse %s as `int64`: %s", s, err)
	return 0
}

// Uint64 returns next uint64 column
func (gr *Reader) Uint64() uint64 {
	if gr.err != nil {
		return 0
	}
	b, err := gr.nextColumn()
	if err != nil {
		gr.err = fmt.Errorf("cannot read `uint64`: %s", err)
		return 0
	}

	s := bytesToString(b)
	n, err := strconv.ParseUint(s, 10, 64)

	if err == nil && 0 <= n && n <= math.MaxUint64 {
		return uint64(n)
	}
	gr.err = fmt.Errorf("cannot parse %s as `uint64`: %s", s, err)
	return 0
}

// String returns next string column
func (gr *Reader) String() string {
	if gr.err != nil {
		return ""
	}
	b, err := gr.nextColumn()
	if err != nil {
		gr.err = fmt.Errorf("cannot read `string`: %s", err)
		return ""
	}
	return string(b)
}

func (gr *Reader) nextColumn() ([]byte, error) {
	gr.col++
	if gr.readBuff == nil {
		return nil, fmt.Errorf("no more columns")
	}

	n := bytes.IndexByte(gr.colBuff, '\t') // look for tab
	if n < 0 {
		// tab is not found, the most right column
		read := gr.colBuff
		gr.colBuff = nil
		return read, nil
	}
	read := gr.colBuff[:n]
	gr.colBuff = gr.colBuff[n+1:]
	return read, nil
}

func bytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b)) // faster than string(b)
}
