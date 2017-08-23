package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type packageInfo struct {
	Dir        string
	ImportPath string
	Name       string
	Doc        string
	Root       string
	Imports    []string
	Deps       []string
}

func goList(pkgs []string) ([]packageInfo, error) {
	args := []string{"list", "-json"}
	args = append(args, pkgs...)
	out, err := exec.Command("go", args...).Output()
	if err != nil {
		return nil, err
	}

	var packageInfoList []packageInfo
	jsonDec := json.NewDecoder(bytes.NewReader(out))
	for {
		var pkg packageInfo
		err := jsonDec.Decode(&pkg)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		packageInfoList = append(packageInfoList, pkg)
	}
	fmt.Println(packageInfoList)
	return packageInfoList, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("need set package in arg")
	}

	inputPkgs, err := goList(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	// not work....
	allPkgs, err := goList([]string{"..."})
	fmt.Printf("%q", allPkgs)
	if err != nil {
		log.Fatal(err)
	}
	var result []packageInfo
	for _, pkg := range allPkgs {
		if isDep(pkg, inputPkgs) {
			result = append(result, pkg)
		}
	}

	for _, r := range result {
		fmt.Println(r.Dir)
	}
}

func isDep(pkg packageInfo, deps []packageInfo) bool {
	for _, dep := range deps {
		if contains(pkg.Deps, dep.ImportPath) {
			return true
		}
	}
	return false
}

func contains(list []string, target string) bool {
	for _, item := range list {
		if target == item {
			return true
		}
	}
	return false
}
