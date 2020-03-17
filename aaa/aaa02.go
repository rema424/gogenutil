package aaa

import "fmt"

const (
	AAA_02_00 = 0
	AAA_02_01 = 1
	AAA_02_02 = 2
)

var (
	aaa_02_00 = "0"
	aaa_02_01 = "1"
	aaa_02_02 = "2"
)

// type AAA02 struct {
// 	Field1 int
// 	Field2 string
// }

func PrintAAA02() {
	fmt.Println(aaa_02_00, aaa_02_01, aaa_02_02)
}

// func (a *AAA02) Print() {
// 	log.Println(a.Field1, a.Field2)
// }
