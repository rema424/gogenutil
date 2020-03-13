package ccc

import (
	"fmt"
	"log"
)

const (
	CCC_03_00 = 0
	CCC_03_01 = 1
	CCC_03_02 = 2
)

var (
	ccc_03_00 = "0"
	ccc_03_01 = "1"
	ccc_03_02 = "2"
)

type CCC03 struct {
	Field1 int
	Field2 string
}

func PrintCCC03() {
	fmt.Println(ccc_03_00, ccc_03_01, ccc_03_01)
}

func (c *CCC03) Print() {
	log.Println(c.Field1, c.Field2)
}
