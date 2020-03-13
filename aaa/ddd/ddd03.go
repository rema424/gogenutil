package ddd

import (
	"fmt"
	"log"
)

const (
	DDD_03_00 = 0
	DDD_03_01 = 1
	DDD_03_02 = 2
)

var (
	ddd_03_00 = "0"
	ddd_03_01 = "1"
	ddd_03_02 = "2"
)

type DDD03 struct {
	Field1 int
	Field2 string
}

func PrintDDD03() {
	fmt.Println(ddd_03_00, ddd_03_01, ddd_03_01)
}

func (d *DDD03) Print() {
	log.Println(d.Field1, d.Field2)
}
