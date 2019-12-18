package main

import (
	"fmt"
	"github.com/k0kubun/pp"
	"reflect"
)

type Hoge struct {
	Id   string `label:"fuga"`
	Name string `label:"1"`
}

func main() {
	hoge := Hoge{
		Id:   "hello",
		Name: "world",
	}
	t := reflect.TypeOf(hoge)
	fmt.Printf("Struct.Name => %s\n", t.Name())
	fmt.Printf("Struct.PkgPath => %s\n", t.PkgPath())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		//fmt.Println(f.Tag)
		fmt.Printf("Name => %s\n", f.Name)
		fmt.Printf("Type.Name => %s\n", f.Type.Name())
		fmt.Printf("Type.PkgPath => %s\n", f.Type.PkgPath())
		fmt.Printf("Anonymous => %v\n", f.Anonymous)
		fmt.Printf("PkgPath => %s\n", f.PkgPath)
		fmt.Printf("Index => %v\n", f.Index)
		fmt.Printf("Offset => %v\n", f.Offset)
		fmt.Println("Tag.Label => " + f.Tag.Get("label"))
	}
}

func debug(args ...interface{}) {
	pp.Println(args...)
}
