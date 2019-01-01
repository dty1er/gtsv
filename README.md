# gtsv - Fast TSV Parser written in Go

TODO: 

* write this
* write godoc
* unescape
* handle non enough buffer

```
package main

import (
	"bytes"
	"fmt"

	"github.com/yagi5/gtsv"
)

func main() {
	tsv := "1\t2.1\ta\n" +
		"4\t5.2\tb\n" +
		"7\t8.3\tc\n"

	gt := gtsv.New(bytes.NewBufferString(tsv))

	for gt.Next() {
		fmt.Println(gt.Int())
		fmt.Println(gt.Float64())
		fmt.Println(gt.String())
	}

	fmt.Println(gt.Error())

	/*

		Output:

		1
		2.1
		a
		4
		5.2
		b
		7
		8.3
		c
		<nil>

	*/

	invalidTSV()
}

func invalidTSV() {
	tsvInvalid := "1\t2.1\ta\n" +
		"4\t5.2\tb\n" +
		"a\t8.3\tc\n" // first column is not int

	gt := gtsv.New(bytes.NewBufferString(tsvInvalid))

	for gt.Next() {
		fmt.Println(gt.Int())
		fmt.Println(gt.Float64())
		fmt.Println(gt.String())
	}

	if err := gt.Error(); err != nil {
		er := err.(gtsv.Error)
		fmt.Printf("error row: %d, col: %d", er.Row(), er.Col())

		// Output: error row: 3, col: 1
	}
}
```
