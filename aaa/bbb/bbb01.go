package bbb

import (
	"fmt"
	"log"
)

const (
	BBB_01_00 = 0
	BBB_01_01 = 1
	BBB_01_02 = 2
)

var (
	bbb_01_00 = "0"
	bbb_01_01 = "1"
	bbb_01_02 = "2"
)

type BBB01 struct {
	Field1 int
	Field2 string
}

func Printbbb01() {
	fmt.Println(bbb_01_00, bbb_01_01, bbb_01_01)
}

func (b *BBB01) Print() {
	log.Println(b.Field1, b.Field2)
}
