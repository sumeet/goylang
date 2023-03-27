package main

import (
	"golang.org/x/tools/go/packages"
	"path/filepath"
)

func typesForPackage(pkgname string) (map[string]string, error) {
	pkgs, err := packages.Load(&packages.Config{Mode: packages.NeedName | packages.NeedImports | packages.NeedTypes}, pkgname)
	if err != nil {
		return nil, err
	}

	types := make(map[string]string, 0)
	for _, pkg := range pkgs {
		scope := pkg.Types.Scope()
		for _, name := range scope.Names() {
			obj := scope.Lookup(name)
			if obj != nil {
				types[name] = scope.Lookup(name).Type().String()
			}
		}
	}

	return types, nil
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
