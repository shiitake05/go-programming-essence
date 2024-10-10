// I/Oを担当するパッケージ
package main

import (
	"io"
	"os"
)

// io.Reader, io.Writer, io.ReadWriterの定義
// type Reader interface {
// 	Read(p []byte) (n int, err error)
// }

// type Writer interface {
// 	Write(p []byte) (n int, err error)
// }

// type ReadWriter interface {
// 	Reader
// 	Writer
// }

// 自分で作成したstructなどにWriteメソッドを実装と、fml.Fprintやio.Copyの引数として渡すことができる
type Foo struct {
}

func (f *Foo) Write(p []byte) (n int, err error) {
	return 0, nil
}

func main() {
	var foo Foo

	io.Copy(&foo, os.Stdin)
}
