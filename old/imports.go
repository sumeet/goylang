package main

import (
	"go/types"
	"golang.org/x/tools/go/packages"
	"path/filepath"
)

func getTypeForPackage(pkgname string, name string) types.Type {
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName | packages.NeedImports | packages.NeedTypes}, pkgname)
	if err != nil {
		panic(err)
	}

	for _, pkg := range pkgs {
		found := pkg.Types.Scope().Lookup(name)
		if found != nil {
			return found.Type()
		}

	}
	return nil
}

func isVersionSuffix(s string) bool {
	if s[0] != 'v' {
		return false
	}

	for _, c := range s[1:] {
		if c < '0' || c > '9' {
			return false
		}
	}

	return true
}

func packageScopeNameFromPackagePath(pkgpath string) string {
	base := filepath.Base(pkgpath)
	if isVersionSuffix(base) {
		return filepath.Base(filepath.Dir(pkgpath))
	}
	return base
}
