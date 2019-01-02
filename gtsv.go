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
	reader       io.Reader
	readBuff     []byte // temporary buffer which stores line
	colBuff      []byte // buffer which stores current column
	reservedBuff []byte // basically won't used. if `buff` is not enough to store line, copy readBuff into this for backup.
	readErr      error
	col          int
	row          int
	err          error
	needUnescape bool

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

// HasNextColumn returns client called Next() even row still has unread column
func (gr *Reader) HasNextColumn() bool {
	return gr.colBuff != nil && gr.readBuff != nil
}

// Next returns true when next row exists.
func (gr *Reader) Next() bool {
	if gr.err != nil {
		return false
	}

	if gr.HasNextColumn() {
		gr.col++ // gtsverror.col will be unread column position number
		gr.err = gr.newError()
		return false
	}

	gr.col = 0
	gr.row++
	for {
		if len(gr.readBuff) <= 0 {
			if gr.readErr != nil {
				gr.err = gr.readErr
				if gr.err != io.EOF {
					gr.err = gr.newError()

				} else if len(gr.reservedBuff) > 0 {
					gr.err = gr.newError()

				} else {
					gr.err = nil
				}
				return false
			}
			n, err := gr.reader.Read(gr.buff[:]) // first, read and get some bytes and store to buffer
			gr.readBuff = gr.buff[:n]
			if err == io.EOF {
				gr.readErr = err
			} else if err != nil {
				gr.readErr = gr.newError()
			}
			gr.needUnescape = (bytes.IndexByte(gr.readBuff, '\\') >= 0)

		}

		n := bytes.IndexByte(gr.readBuff, '\n') // read from buffer
		if n >= 0 {
			// next row found
			read := gr.readBuff[:n]
			gr.readBuff = gr.readBuff[n+1:]

			// append reservedBuff
			if len(gr.reservedBuff) > 0 {
				gr.reservedBuff = append(gr.reservedBuff, read...)
				read = gr.reservedBuff
				gr.reservedBuff = gr.reservedBuff[:0] // make empty
			}
			gr.colBuff = read
			return true
		}
		gr.reservedBuff = append(gr.reservedBuff, gr.readBuff...)
		gr.readBuff = nil
	}
}

// Int returns next int column
func (gr *Reader) Int() int {
	if gr.err != nil {
		return 0
	}
	b, err := gr.nextColumn()
	if err != nil {
		gr.err = gr.newError()
		return 0
	}

	s := bytesToString(b)
	n, err := strconv.Atoi(s)
	if err == nil {
		return n
	}
	gr.err = gr.newError()
	return 0
}

// Uint returns next uint column
func (gr *Reader) Uint() uint {
	if gr.err != nil {
		return 0
	}
	b, err := gr.nextColumn()
	if err != nil {
		gr.err = gr.newError()
		return 0
	}

	s := bytesToString(b)
	n, err := strconv.Atoi(s)
	if err == nil && 0 <= n {
		return uint(n)
	}

	gr.err = gr.newError()
	return 0
}

// Int8 returns next int8 column
func (gr *Reader) Int8() int8 {
	if gr.err != nil {
		return 0
	}
	b, err := gr.nextColumn()
	if err != nil {
		gr.err = gr.newError()
		return 0
	}

	s := bytesToString(b)
	n, err := strconv.Atoi(s)
	if err == nil && math.MinInt8 <= n && n <= math.MaxInt8 {
		return int8(n)
	}
	gr.err = gr.newError()
	return 0
}

// Uint8 returns next uint8 column
func (gr *Reader) Uint8() uint8 {
	if gr.err != nil {
		return 0
	}
	b, err := gr.nextColumn()
	if err != nil {
		gr.err = gr.newError()
		return 0
	}

	s := bytesToString(b)
	n, err := strconv.Atoi(s)
	if err == nil && 0 <= n && n <= math.MaxUint8 {
		return uint8(n)
	}
	gr.err = gr.newError()
	return 0
}

// Int16 returns next int16 column
func (gr *Reader) Int16() int16 {
	if gr.err != nil {
		return 0
	}
	b, err := gr.nextColumn()
	if err != nil {
		gr.err = gr.newError()
		return 0
	}

	s := bytesToString(b)
	n, err := strconv.Atoi(s)
	if err == nil && math.MinInt16 <= n && n <= math.MaxInt16 {
		return int16(n)
	}
	gr.err = gr.newError()
	return 0
}

// Uint16 returns next uint16 column
func (gr *Reader) Uint16() uint16 {
	if gr.err != nil {
		return 0
	}
	b, err := gr.nextColumn()
	if err != nil {
		gr.err = gr.newError()
		return 0
	}

	s := bytesToString(b)
	n, err := strconv.Atoi(s)
	if err == nil && 0 <= n && n <= math.MaxUint16 {
		return uint16(n)
	}
	gr.err = gr.newError()
	return 0
}

// Int32 returns next int32 column
func (gr *Reader) Int32() int32 {
	if gr.err != nil {
		return 0
	}
	b, err := gr.nextColumn()
	if err != nil {
		gr.err = gr.newError()
		return 0
	}

	s := bytesToString(b)
	n, err := strconv.Atoi(s)
	if err == nil && math.MinInt32 <= n && n <= math.MaxInt32 {
		return int32(n)
	}
	gr.err = gr.newError()
	return 0
}

// Uint32 returns next uint32 column
func (gr *Reader) Uint32() uint32 {
	if gr.err != nil {
		return 0
	}
	b, err := gr.nextColumn()
	if err != nil {
		gr.err = gr.newError()
		return 0
	}

	s := bytesToString(b)
	n, err := strconv.Atoi(s)
	if err == nil && 0 <= n && n <= math.MaxUint32 {
		return uint32(n)
	}
	gr.err = gr.newError()
	return 0
}

// Int64 returns next int64 column
func (gr *Reader) Int64() int64 {
	if gr.err != nil {
		return 0
	}
	b, err := gr.nextColumn()
	if err != nil {
		gr.err = gr.newError()
		return 0
	}

	s := bytesToString(b)
	n, err := strconv.ParseInt(s, 10, 64)
	if err == nil && math.MinInt64 <= n && n <= math.MaxInt64 {
		return int64(n)
	}
	gr.err = gr.newError()
	return 0
}

// Uint64 returns next uint64 column
func (gr *Reader) Uint64() uint64 {
	if gr.err != nil {
		return 0
	}
	b, err := gr.nextColumn()
	if err != nil {
		gr.err = gr.newError()
		return 0
	}

	s := bytesToString(b)
	n, err := strconv.ParseUint(s, 10, 64)

	if err == nil && 0 <= n && n <= math.MaxUint64 {
		return uint64(n)
	}
	gr.err = gr.newError()
	return 0
}

// Float32 returns next float32 column
func (gr *Reader) Float32() float32 {
	if gr.err != nil {
		return 0
	}
	b, err := gr.nextColumn()
	if err != nil {
		gr.err = gr.newError()
		return 0
	}

	s := bytesToString(b)
	n, err := strconv.ParseFloat(s, 32)

	if err == nil {
		return float32(n)
	}
	gr.err = gr.newError()
	return 0
}

// Float64 returns next float64 column
func (gr *Reader) Float64() float64 {
	if gr.err != nil {
		return 0
	}
	b, err := gr.nextColumn()
	if err != nil {
		gr.err = gr.newError()
		return 0
	}

	s := bytesToString(b)
	n, err := strconv.ParseFloat(s, 64)

	if err == nil {
		return n
	}
	gr.err = gr.newError()
	return 0
}

// Bytes returns next []byte column
func (gr *Reader) Bytes() []byte {
	if gr.err != nil {
		return nil
	}
	b, err := gr.nextColumn()
	if err != nil {
		gr.err = gr.newError()
		return nil
	}

	if !gr.needUnescape {
		return b
	}

	// need to unescape
	n := bytes.IndexByte(b, '\\')
	if n < 0 {
		// Actually there was no '\'. Basically won't reach here
		return b
	}

	n++
	// for example: https://play.golang.org/p/UXkXBXLgsP_O
	d := b[:n]
	b = b[n:]
	for len(b) > 0 {
		switch b[0] {
		case 'b':
			d[len(d)-1] = '\b'
		case 'f':
			d[len(d)-1] = '\f'
		case 'r':
			d[len(d)-1] = '\r'
		case 'n':
			d[len(d)-1] = '\n'
		case 't':
			d[len(d)-1] = '\t'
		case '0':
			d[len(d)-1] = 0
		case '\'':
			d[len(d)-1] = '\''
		case '\\':
			d[len(d)-1] = '\\'
		default:
			d[len(d)-1] = b[0]
		}

		b = b[1:]
		n = bytes.IndexByte(b, '\\')
		if n < 0 {
			d = append(d, b...)
			break
		}
		n++
		d = append(d, b[:n]...)
		b = b[n:]
	}
	return d
}

// String returns next string column
func (gr *Reader) String() string {
	return string(gr.Bytes())
}

// Bool returns next bool column
func (gr *Reader) Bool() bool {
	if gr.err != nil {
		return false
	}
	b, err := gr.nextColumn()
	if err != nil {
		gr.err = gr.newError()
		return false
	}

	s := bytesToString(b)
	n, err := strconv.ParseBool(s)
	if err == nil {
		return n
	}
	gr.err = gr.newError()
	return false
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

func (gr *Reader) newError() *gtsverror {
	return &gtsverror{row: gr.row, col: gr.col}
}

func bytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b)) // faster than string(b)
}
