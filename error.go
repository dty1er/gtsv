package gtsv

import (
	"fmt"
)

// Error is the error interface.
// If `gt.Error()` returned non-nil,
// usually it implements this interface.
// So, Row() and Col() will return error position.
type Error interface {
	Row() int
	Col() int
}

// gtsverror contains row, col, type
type gtsverror struct {
	row int
	col int
}

// Row returns the row number error occurred
func (e *gtsverror) Row() int {
	return e.row
}

// Col returns the col number error occurred
func (e *gtsverror) Col() int {
	return e.col
}

// Error returns error message
func (e *gtsverror) Error() string {
	return fmt.Sprintf("Parse failed at row #%d, col #%d", e.row, e.col)
}
