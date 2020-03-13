package main

import (
	"gogenutil/gogenutil"
	"log"

	"github.com/k0kubun/pp"
)

func main() {
	path := "./aaa/aaa01.go"
	if err := run(path); err != nil {
		log.Fatalln(err)
	}
}

func run(path string) error {
	pp.Println("path", path)

	dir, err := gogenutil.AbsDir(path)
	if err != nil {
		log.Println(err)
		return err
	}
	pp.Println("dir", dir)

	pkgs, err := gogenutil.LoadPkgs(dir, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	pp.Println("len(pkgs)", len(pkgs))
	// pp.Println("pkgs", pkgs[0])

	for _, pkg := range pkgs {
		gogenutil.PrintPackage(pkg)
		// for _, astFile := range pkg.Syntax {
		// 	// pp.Println("astFile", astFile)
		// 	// fmt.Println("astFile", astFile)
		// }
	}

	return nil
}
