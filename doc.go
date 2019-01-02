/*
Package gtsv is the package to read TSV formatted text.

It is designed to type-specifically use each tsv columns,
and easily find row and column number which error happened.

This is the example to treat valid TSV.


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
    }

While gt.Next() returns true, you can load column value type-specific,
like `gt.Int()` , `gt.String()` .

Of course, you can use this with struct.

    type user struct {
      name string
      age  int
      male bool
    }

    func main() {
      tsv := "john\t18\ttrue\n" +
        "emily\t16\tfalse\n"

      gt := gtsv.New(bytes.NewBufferString(tsv))

     var users []*user
      for gt.Next() {
        users = append(users, &user{name: gt.String(), age: gt.Int(), male: gt.Bool()})
      }

      fmt.Println(gr.Error())
    }

If you change the order to call `gt.String()` and `gt.Int()` ,
`gt.Error()` will be non-nil because it's not the same as TSV column order.

Sometimes, we want to know the error position of these files.
If you need it, you can cast `gt.Error()` (this is just a `error`) into `gtsv.Error`
to know error position. like this:


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
    }

if cast `err.(gtsv.Error)` succeeded, `Row()` and `Col()` show you error position.

*/
package gtsv
