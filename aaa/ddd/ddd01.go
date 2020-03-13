package ddd

import (
	"fmt"
	"log"
)

const (
	DDD_01_00 = 0
	DDD_01_01 = 1
	DDD_01_02 = 2
)

var (
	ddd_01_00 = "0"
	ddd_01_01 = "1"
	ddd_01_02 = "2"
)

type DDD01 struct {
	Field1 int
	Field2 string
}

func PrintDDD01() {
	fmt.Println(ddd_01_00, ddd_01_01, ddd_01_01)
}

func (d *DDD01) Print() {
	log.Println(d.Field1, d.Field2)
}
