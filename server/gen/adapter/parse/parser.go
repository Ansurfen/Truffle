package parse

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	ARGC    = 3
	newLine = "\n\n"
)

var (
	argsReg = regexp.MustCompile(`@\w+`)
	tagReg  = regexp.MustCompile(`ddd:"(.*?)"`)
)

type (
	SrcOpt struct {
		Cap    bool
		OutPkg string
		InDir  string
		OutDir string
		Module string
		Filter []string
	}

	SrcPkg struct {
		opt   *SrcOpt
		files map[string]*SrcFile // .go -> file
		dirs  []string
	}

	SrcFile struct {
		funcs map[string]*Function
	}

	Function struct {
		req *SrcStc
		ret *SrcStc
	}

	SrcStc struct {
		tag         string
		packageName string
		strcutName  string
		fields      map[string]struct {
			_type string
			tags  map[string][]string
		}
	}
)

func NewSrcPkg(opt *SrcOpt) *SrcPkg {
	return &SrcPkg{
		opt:   opt,
		files: make(map[string]*SrcFile),
		dirs:  make([]string, 0),
	}
}

func NewSrcFile() *SrcFile {
	return &SrcFile{
		funcs: make(map[string]*Function),
	}
}

func (srcPkg *SrcPkg) ReadDirs() {
	dirs, err := ioutil.ReadDir(srcPkg.opt.InDir)
	if err != nil {
		log.Fatal(err)
	}
	filterDic := make(map[string]bool)
	for _, rule := range srcPkg.opt.Filter {
		filterDic[rule] = true
	}
	filterDic[srcPkg.opt.OutPkg] = true
	for _, dir := range dirs {
		dirName := dir.Name()
		if _, ok := filterDic[dirName]; !ok {
			srcPkg.dirs = append(srcPkg.dirs, dirName)
		}
	}
}

func (srcPkg *SrcPkg) Parse() {
	for _, dir := range srcPkg.dirs {
		pkg, err := parser.ParseDir(token.NewFileSet(), srcPkg.opt.InDir+dir, func(fi fs.FileInfo) bool { return true }, parser.ParseComments)
		if err != nil {
			log.Fatal(err)
		}
		if pkg[dir] == nil {
			continue
		}
		var (
			structName string
			sf         *SrcFile
		)
		for _, file := range pkg[dir].Files {
			pn := pkg[dir].Name
			for _, decl := range file.Decls {
				if stc, ok := decl.(*ast.GenDecl); ok && stc.Tok == token.TYPE && stc.Doc != nil {
					fields := make(map[string]struct {
						_type string
						tags  map[string][]string
					})
					for _, spec := range stc.Specs {
						if tp, ok := spec.(*ast.TypeSpec); ok {
							structName = tp.Name.Name
							if stp, ok := tp.Type.(*ast.StructType); ok {
								if !stp.Struct.IsValid() {
									continue
								}
								for _, field := range stp.Fields.List {
									if len(field.Names) == 1 {
										fieldName := field.Names[0].Name
										fieldType := ""
										if ft, ok := field.Type.(*ast.Ident); ok {
											fieldType = ft.Name
										}
										fields[fieldName] = struct {
											_type string
											tags  map[string][]string
										}{
											_type: fieldType,
											tags:  make(map[string][]string),
										}
										if field.Tag != nil {
											tag := tagReg.FindStringSubmatch(field.Tag.Value)
											if len(tag) > 0 {
												funcOrder := strings.Split(tag[len(tag)-1], ";")
												for _, order := range funcOrder {
													temp := strings.Split(order, ":")
													if len(temp) != 2 {
														continue
													}
													fields[fieldName].tags[temp[0]] = strings.Split(temp[1], ",")
												}
											}
										}
									}
								}
							}
						}
					}
					docs := strings.Split(stc.Doc.Text(), "\n")
					for _, doc := range docs {
						if argsReg.MatchString(doc) {
							args := strings.Split(doc, " ")
							if len(args) != ARGC {
								log.Printf("argc must is %d\n at %s", ARGC, doc)
								continue
							}
							file := args[2]
							if srcPkg.files[file] == nil {
								srcPkg.files[file] = &SrcFile{
									funcs: make(map[string]*Function),
								}
							}
							sf = srcPkg.files[file]
							funcName := args[0]
							funcName = funcName[1:]
							tag := strings.Split(args[1], ":")[1]
							if strings.HasPrefix(args[1], "req:") {
								if sf.funcs[funcName] == nil {
									sf.funcs[funcName] = new(Function)
								}
								sf.funcs[funcName].req = &SrcStc{
									tag:         tag,
									packageName: pn,
									strcutName:  structName,
									fields: make(map[string]struct {
										_type string
										tags  map[string][]string
									}),
								}
								sf.funcs[funcName].req.fields = fields
							} else if strings.HasPrefix(args[1], "ret:") {
								if sf.funcs[funcName] == nil {
									sf.funcs[funcName] = new(Function)
								}
								sf.funcs[funcName].ret = &SrcStc{
									tag:         tag,
									packageName: pn,
									strcutName:  structName,
									fields: make(map[string]struct {
										_type string
										tags  map[string][]string
									}),
								}
								sf.funcs[funcName].ret.fields = fields
							} else {
							}
						}
					}
				}
			}
		}
	}
}

func (srcPkg *SrcPkg) Print() {
	for filename, file := range srcPkg.files {
		fmt.Println(filename)
		for fn, f := range file.funcs {
			fmt.Printf("\t%v: \n", fn)
			if f.req != nil {
				fmt.Printf("\t\treq: %s %s\t%v\n", f.req.tag, f.req.packageName, f.req.fields)
			}
			if f.ret != nil {
				fmt.Printf("\t\tret: %s %s\t%v\n\n", f.ret.tag, f.ret.packageName, f.ret.fields)
			}
		}
	}
}

func (srcPkg *SrcPkg) Write() {
	for filename, file := range srcPkg.files {
		targetCode := "package " + srcPkg.opt.OutPkg + newLine
		packageNames := make(map[string]string)
		funcs := []string{}
		for funcName, f := range file.funcs {
			if f.req == nil || f.ret == nil {
				continue
			}
			packageNames[f.req.packageName] = "\"" + srcPkg.opt.Module + "/" + f.req.packageName + "\""
			packageNames[f.ret.packageName] = "\"" + srcPkg.opt.Module + "/" + f.ret.packageName + "\""
			funcCode := "func " + funcName + "(req *" + f.req.packageName + "." + f.req.strcutName + ") *" + f.ret.packageName + "." + f.ret.strcutName +
				" {\n\t return &" + f.ret.packageName + "." + f.ret.strcutName + "{"
			if funcCode[len(funcCode)-1] == ',' {
				funcCode = funcCode + "\n\t"
			}
			tag := f.req.tag // req:tag
			if len(tag) > 0 && f.req.fields != nil {
				for fieldName, field := range f.req.fields {
					for _, v := range field.tags[tag] {
						if srcPkg.opt.Cap {
							v = Capitalize(v)
						}
						if _, ok := f.ret.fields[v]; !ok {
							continue
						}
						funcCode += "\n\t\t" + v + ":\t" + "req." + fieldName + ","
					}
					for _, v := range field.tags["*"] {
						if srcPkg.opt.Cap {
							v = Capitalize(v)
						}
						if _, ok := f.ret.fields[v]; !ok {
							continue
						}
						funcCode += "\n\t\t" + v + ":\t" + "req." + fieldName + ","
					}
				}
			}
			if funcCode[len(funcCode)-1] == ',' {
				funcCode = funcCode + "\n\t"
			}
			funcCode += "}" + "\n}" + newLine
			funcs = append(funcs, funcCode)
		}
		if len(packageNames) > 0 {
			targetCode += "import (\n"
			for _, pn := range packageNames {
				targetCode += "\t" + pn + "\n"
			}
			targetCode += ")" + newLine
		}
		for _, fn := range funcs {
			targetCode += fn
		}
		fp, _ := os.Create(srcPkg.opt.OutDir + filename)
		defer fp.Close()
		fp.Write([]byte(targetCode))
	}
}

func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {
				vv[i] -= 32
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}
