package main

import (
	"fmt"
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
		files := gogenutil.Files(pkg)
		for _, f := range files {
			// f.Print(os.Stdout)
			// fmt.Println(f.NodeByOffset(38))
			fmt.Println(f.NodeByOffset(225))
		}
		// for _, astFile := range pkg.Syntax {
		// 	// pp.Println("astFile", astFile)
		// 	// fmt.Println("astFile", astFile)
		// }
	}

	return nil
}
