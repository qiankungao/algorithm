package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {

	r := gin.Default()
	r.Use(Recovery())
	r.GET("/get", GetTest)

	r.Run(":8000")
}

func GetTest(c *gin.Context) {
	str := "nihao哥"
	//panic("程序抛异常")
	GoHelp(GoTest)
	//Go(GoTest)
	time.Sleep(10 * time.Second)
	c.String(200, str)
}

func GoTest() {
	fmt.Println("另外起一个goroutine")
	panic(" 子线程出错")
}

func Go(x func()) {
	if err := recover(); err != nil {
		fmt.Println("panic in Go: ", err)
	}
	go x()
}

func GoHelp(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("panic in GoHelp: ", err)
			}
		}()
		f()
	}()
}


func Recovery() gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("捕获异常:", err)
			}
		}()
		c.Next()
		fmt.Println("后面的要不还要执行")
	}
}
