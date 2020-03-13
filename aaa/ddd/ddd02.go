package ddd

import (
	"fmt"
	"log"
)

const (
	DDD_02_00 = 0
	DDD_02_01 = 1
	DDD_02_02 = 2
)

var (
	ddd_02_00 = "0"
	ddd_02_01 = "1"
	ddd_02_02 = "2"
)

type DDD02 struct {
	Field1 int
	Field2 string
}

func PrintDDD02() {
	fmt.Println(ddd_02_00, ddd_02_01, ddd_02_01)
}

func (d *DDD02) Print() {
	log.Println(d.Field1, d.Field2)
}
