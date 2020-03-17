package gogenutil

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
	"os"
	"path/filepath"

	"github.com/k0kubun/pp"
	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/packages"
	"golang.org/x/xerrors"
)

var (
	ErrNotDir = xerrors.New("gogenutil: not dir")
)

// func IsFileExists(path string) error {
// 	_, err := os.Stat(path)
// 	return err
// }

func IsDir(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		return false
	}
	return f.IsDir()
}

func AbsPath(relPath string) (string, error) {
	eval, err := filepath.EvalSymlinks(relPath)
	if err != nil {
		return "", err
	}
	return filepath.Abs(eval)
}

// filepath or dirpath is ok
func AbsDir(relPath string) (string, error) {
	absPath, err := AbsPath(relPath)
	if err != nil {
		return "", err
	}
	if IsDir(absPath) {
		return absPath, nil
	}
	return filepath.Dir(absPath), nil
}

func LoadPkgs(dir string, overlay map[string][]byte) ([]*packages.Package, error) {
	cfg := &packages.Config{
		Overlay: overlay,
		Mode:    packages.LoadAllSyntax,
		Tests:   true,
		Dir:     dir,
		Fset:    token.NewFileSet(),
		Env:     os.Environ(),
	}

	return packages.Load(cfg)
}

func PrintPackage(pkg *packages.Package) {
	fmt.Println("pkg.Name\t", pkg.Name)
	fmt.Println("pkg.PkgPath\t", pkg.PkgPath)
	fmt.Println("pkg.String()\t", pkg.String())
	fmt.Println("pkg.GoFiles\t", pkg.GoFiles)
	fmt.Printf("pkg.Imports\t %#v\n", pkg.Imports)
	fmt.Printf("pkg.Syntax\t %#v\n", pkg.Syntax)
	for i := 0; i < len(pkg.Syntax); i++ {
		fmt.Println("===================================================================")
		astFile := pkg.Syntax[i]
		fmt.Printf("astFile := pkg.Syntax[%d]\n", i)
		fmt.Printf("astFile.Package\t\t %#v\n", astFile.Package)
		fmt.Printf("astFile.Name\t\t %#v\n", astFile.Name)
		fmt.Printf("astFile.Pos()\t\t %#v\n", astFile.Pos())
		fmt.Printf("astFile.End()\t\t %#v\n", astFile.End())
		fmt.Println()
		pos := astFile.Pos()
		tokenFile := pkg.Fset.File(pos)
		fmt.Printf("tokenFile := pkg.Fset.File(%d)\n", pos)
		fmt.Printf("tokenFile.Name()\t %#v\n", tokenFile.Name())
		fmt.Printf("tokenFile.Base()\t %#v\n", tokenFile.Base())
		fmt.Printf("tokenFile.LineCount()\t %#v\n", tokenFile.LineCount())
		fmt.Printf("tokenFile.Size()\t %#v\n", tokenFile.Size())
		fmt.Println("===================================================================")
	}
	// fmt.Printf("pkg.Fset\t %#v\n", pkg.Fset.)
}

type File struct {
	Name      string
	Pkg       string
	PkgPath   string
	AstFile   *ast.File
	TokenFile *token.File
}

func Files(pkg *packages.Package) map[string]*File {
	tokenFileSet := pkg.Fset

	files := make(map[string]*File, len(pkg.GoFiles))
	for _, astFile := range pkg.Syntax {
		tokenFile := tokenFileSet.File(astFile.Pos())
		files[tokenFile.Name()] = &File{
			Name:      tokenFile.Name(),
			Pkg:       pkg.Name,
			PkgPath:   pkg.PkgPath,
			AstFile:   astFile,
			TokenFile: tokenFile,
		}
	}
	return files
}

func (f *File) Print(w io.Writer) {
	fmt.Println("===================================================================")
	fmt.Fprintf(w, "*File.Name\t%s\n", f.Name)
	fmt.Fprintf(w, "*File.Pkg\t%s\n", f.Pkg)
	fmt.Fprintf(w, "*File.PkgPath\t%s\n", f.PkgPath)
	fmt.Fprintf(w, "*File.AstFile\t%+v\n", f.AstFile)
	fmt.Fprintf(w, "*File.TokenFile\t%+v\n", f.TokenFile)
	fmt.Println("===================================================================")
}

func (f *File) NodeByOffset(offset int) (*ast.Node, token.Pos, error) {
	if offset > f.TokenFile.Size() {
		return nil, 0, fmt.Errorf("file size (%d) is smaller than given offset (%d)", f.TokenFile.Size(), offset)
	}

	tokenPos := f.TokenFile.Pos(offset)
	path, e := astutil.PathEnclosingInterval(f.AstFile, tokenPos, tokenPos)
	pp.Println(len(path), path, e)
	return nil, f.TokenFile.Pos(offset), nil
}

// func PrintPackages(pkgs []*packages.Package) {
//   for _, pkg :=
// }
