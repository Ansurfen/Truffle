package main

import (
	"flag"
	. "truffle/gen/adapter/flag"
	. "truffle/gen/adapter/parse"
)

var (
	autoCapitalize = flag.Bool("cap", true, "AutoCapitalize tag")
	outPackageName = flag.String("outpkg", "adapter", "output package name")
	inDir          = flag.String("indir", "./ws/ddd/", "input dir")
	outDir         = flag.String("outdir", "./ws/ddd/adapter/", "output dir")
	moduleDir      = flag.String("module", "truffle/ws/ddd", "module dir")
	filter         []string
)

func main() {
	flag.Var(NewSliceValue([]string{}, &filter), "filter", "filter dir")
	pkg := NewSrcPkg(&SrcOpt{
		Cap:    *autoCapitalize,
		OutPkg: *outPackageName,
		InDir:  *inDir,
		OutDir: *outDir,
		Module: *moduleDir,
		Filter: filter,
	})
	pkg.ReadDirs()
	pkg.Parse()
	// pkg.Print()
	pkg.Write()
}
