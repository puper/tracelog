package main

import (
	"fmt"
	"time"

	"github.com/puper/tracelog"
)

func main() {
	tl := tracelog.New("main")
	tl.Caller().Env(1111, 2222)
	reply, err := T1(1)
	tl.Arg(1).Reply(reply).Error(err).Stop(time.Millisecond)
	t2 := tl.Module("tttttt")
	t2.Stop().Error("ffff")
	fmt.Println(string(tl.Json()))
}

func T1(i int) (int, error) {
	time.Sleep(time.Second)
	return i, nil
}

/**
日志过滤
	- 生成日志的策略
		- 根据策略筛选出需要的日志
*/
