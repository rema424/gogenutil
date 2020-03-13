package ccc

import (
	"fmt"
	"log"
)

const (
	CCC_01_00 = 0
	CCC_01_01 = 1
	CCC_01_02 = 2
)

var (
	ccc_01_00 = "0"
	ccc_01_01 = "1"
	ccc_01_02 = "2"
)

type CCC01 struct {
	Field1 int
	Field2 string
}

func PrintCCC01() {
	fmt.Println(ccc_01_00, ccc_01_01, ccc_01_01)
}

func (c *CCC01) Print() {
	log.Println(c.Field1, c.Field2)
}
