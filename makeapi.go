package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

func createGoFile(data TemplateData, tplfile, gofile string) error {
	targetFile := filepath.Join(data.PkgPath, gofile)
	fmt.Println("    create Go file:", targetFile)

	tpl, err := ioutil.ReadFile(filepath.Join(data.ViewPath, tplfile))
	if err != nil {
		return err
	}
	t1 := template.New("t1")
	t1, err = t1.Parse(string(tpl))
	if err != nil {
		return err
	}

	err = os.MkdirAll(data.PkgPath, 0600)
	if err != nil {
		return err
	}

	f, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer f.Close()
	t1.Execute(f, data)

	return nil
}

func MakePkg(tData TemplateData) (err error) {
	fmt.Println("    [api count] =", len(tData.Apis), "    [struct count] =", len(tData.Structs))

	var oneAPIData TemplateData
	oneAPIData.PkgDesc = tData.PkgDesc
	oneAPIData.PkgName = tData.PkgName
	oneAPIData.PkgPath = tData.PkgPath
	oneAPIData.ViewPath = tData.ViewPath
	oneAPIData.Apis = nil
	oneAPIData.PackageSite = confMain.PackageSite

	for _, api := range tData.Apis {
		oneAPIData.Apis = nil
		oneAPIData.Apis = append(oneAPIData.Apis, api)
		goFile := api.Name + ".go"
		err = createGoFile(oneAPIData, "api.tpl", goFile)
		if err != nil {
			return err
		}
	}

	err = createGoFile(tData, "struct.tpl", tData.PkgName+"_structs.go")
	if err != nil {
		return err
	}

	err = createGoFile(tData, "doc.tpl", tData.PkgName+"_doc.go")
	if err != nil {
		return err
	}

	fmt.Println("    format:", tData.PkgFullName)
	cmd := exec.Command("go", "fmt", tData.PkgFullName)
	err = cmd.Run()
	if err != nil {
		return err
	}

	fmt.Println("    install:", tData.PkgFullName)
	cmd = exec.Command("go", "install", tData.PkgFullName)
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil

}

type TemplateData struct {
	MetaVersionNo string
	PkgName       string
	PkgFullName   string
	PkgPath       string
	PkgDesc       string
	ViewPath      string
	Apis          []*ApiT
	Structs       []*StructT
	PackageSite   string
}

func MakeApis(confPackage *ConfPackageT, data *DataT) *[]string {

	fmt.Println("\n----------------------------------------------")
	fmt.Println("make apis")
	fmt.Println("----------------------------------------------")

	var GOPATH string
	GOPATH = os.Getenv("GOPATH")
	goPaths := strings.Split(GOPATH, ";")
	defaultPath := goPaths[0] + "/src/" + confMain.PackageSite

	os.MkdirAll(defaultPath, os.ModeDir)

	makeCore(defaultPath)

	errArray := make([]string, 0)

	var tdata TemplateData
	tdata.MetaVersionNo = data.MetaVersionNo

	for _, v := range confPackage.Mx {
		if !v.PkgChoose {
			continue
		}

		tdata.PkgFullName = confMain.PackageSite + "/api/" + v.Name
		fmt.Println("  make:", tdata.PkgFullName)

		tdata.PkgDesc = v.Desc
		tdata.PkgName = v.Name
		tdata.PkgPath = defaultPath + "/api/" + v.Name
		tdata.ViewPath = "./template/"
		tdata.Apis = data.MapPkgApi[v.Name]
		tdata.Structs = data.MapPkgStruct[v.Name]
		tdata.PackageSite = confMain.PackageSite
		err := MakePkg(tdata)
		if err != nil {
			fmt.Println("      [error]", err)
			errArray = append(errArray, "package ["+tdata.PkgFullName+"] make err!")
		}
	}

	return &errArray
}

func makeCore(defaultPath string) error {
	var tdata TemplateData
	tdata.PkgPath = defaultPath
	tdata.ViewPath = "./template/"
	tdata.PackageSite = confMain.PackageSite

	err := createGoFile(tdata, "Taobao.tpl", "Taobao.go")
	if err != nil {
		return err
	}

	err = createGoFile(tdata, "TaobaoAuth.tpl", "TaobaoAuth.go")
	if err != nil {
		return err
	}

	err = createGoFile(tdata, "TaobaoErrResponse.tpl", "TaobaoErrResponse.go")
	if err != nil {
		return err
	}

	err = createGoFile(tdata, "TaobaoRequest.tpl", "TaobaoRequest.go")
	if err != nil {
		return err
	}

	err = createGoFile(tdata, "TaobaoMethodRequest.tpl", "TaobaoMethodRequest.go")
	if err != nil {
		return err
	}

	return nil
}
