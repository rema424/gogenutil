package bbb

import (
	"fmt"
	"log"
)

const (
	BBB_02_00 = 0
	BBB_02_01 = 1
	BBB_02_02 = 2
)

var (
	bbb_02_00 = "0"
	bbb_02_01 = "1"
	bbb_02_02 = "2"
)

type BBB02 struct {
	Field1 int
	Field2 string
}

func Printbbb02() {
	fmt.Println(bbb_02_00, bbb_02_01, bbb_02_02)
}

func (b *BBB02) Print() {
	log.Println(b.Field1, b.Field2)
}
