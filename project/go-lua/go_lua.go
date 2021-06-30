package main

import (
	"fmt"
	"github.com/aarzilli/golua/lua"
)

func main() {
	l := lua.NewState()
	//l.OpenLibs()
	defer l.Close()

	//家在lua文件
	if err := l.DoString(""); err != nil {
		fmt.Println("")
	}

}
