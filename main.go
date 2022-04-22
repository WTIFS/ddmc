package main

import (
	"github.com/wtifs/ddmc/service/cart"
	"time"
)

func main() {
	runSingle()
	//runCron()
}

// 单次运行
func runSingle() {
	if cart.CheckCart() {
		cart.CheckDeliverTime()
	}
}

// 定时任务
func runCron() {
	ticker := time.NewTicker(time.Second)
	for t := range ticker.C {
		if t.Minute()%10 == 0 && t.Second() == 0 {
			runSingle()
		}
	}
}
