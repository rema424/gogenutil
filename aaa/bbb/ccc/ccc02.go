package ccc

import (
	"fmt"
	"log"
)

const (
	CCC_02_00 = 0
	CCC_02_01 = 1
	CCC_02_02 = 2
)

var (
	ccc_02_00 = "0"
	ccc_02_01 = "1"
	ccc_02_02 = "2"
)

type CCC02 struct {
	Field1 int
	Field2 string
}

func PrintCCC02() {
	fmt.Println(ccc_02_00, ccc_02_01, ccc_02_01)
}

func (c *CCC02) Print() {
	log.Println(c.Field1, c.Field2)
}
