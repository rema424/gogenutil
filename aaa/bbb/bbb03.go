package bbb

import (
	"fmt"
	"log"
)

const (
	BBB_03_00 = 0
	BBB_03_01 = 1
	BBB_03_02 = 2
)

var (
	bbb_03_00 = "0"
	bbb_03_01 = "1"
	bbb_03_02 = "2"
)

type BBB03 struct {
	Field1 int
	Field2 string
}

func Printbbb03() {
	fmt.Println(bbb_03_00, bbb_03_01, bbb_03_01)
}

func (b *BBB03) Print() {
	log.Println(b.Field1, b.Field2)
}
