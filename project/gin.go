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


