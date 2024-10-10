// structの合成
// 継承されているわけではないため、埋め込んだStructから埋め込まれたstructを参照することができない(委譲(コンポジション)としての継承は「composition」で実現している)
package main

import "fmt"

type Attr struct {
	Name string
	Age  int
}

func (a Attr) String() string {
	return fmt.Sprintf("%s(%d)", a.Name, a.Age)
}

type AttrEx struct {
	Name string
}

func (a AttrEx) String() string {
	return fmt.Sprintf("a.k.a. %s", a.Name)
}

type Teacher struct {
	Attr
	AttrEx
	Subject string
}

func main() {
	teacher := Teacher{
		Attr: Attr{
			Name: "John Schwartz",
			Age:  43,
		},
		AttrEx: AttrEx{
			Name: "JS",
		},
		Subject: "Math",
	}
	fmt.Println(teacher.Attr.String(), teacher.AttrEx.String(), teacher.Subject) // John Schwartz(43) a.k.a. JS Math
}
