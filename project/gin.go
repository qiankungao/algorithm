package main

import (
	"fmt"
	"github.com/robfig/cron"
	"sync"
	"time"
)

func main1() {

	sy := sync.WaitGroup{}
	sy.Add(1)
	go func() {
		defer sy.Done()
		Timer()
	}()
	sy.Wait()
	select {}
}

func Timer() {
	fmt.Println("定时任务")
	c := cron.New()
	c.AddFunc("*/3 * * * * *", func() {
		fmt.Println("美妙执行",time.Now().Unix())
	})
	c.Start()
}

//package main
//
//import (
//	"fmt"
//	_ "github.com/1975210542/BigDataAlgorithm/project/docs"
//	"github.com/gin-gonic/gin"
//	"github.com/swaggo/gin-swagger"
//	"github.com/swaggo/gin-swagger/swaggerFiles"
//)
//
//func main1() {
//	r := gin.Default()
//	//url := ginSwagger.URL("http://127.0.0.1:8000/swagger/doc.json")
//	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
//
//	r.GET("/ping", ping)
//	r.Run("0.0.0.0:8089") // listen and serve on 0.0.0.0:8080
//}
//
//// @Summary 新增文章标签
//// @Tags ping
//// @Produce  json
//// @Param name query string false "Name"
//// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
//// @Router /ping [POST]
//func ping(c *gin.Context) {
//	//输出json结果给调用方
//
//	name, _ := c.GetQuery("name")
//	fmt.Println("name:", name)
//	c.JSON(200, gin.H{
//		"message": "pong" + name,
//	})
//}
