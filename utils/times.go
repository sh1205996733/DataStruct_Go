package utils

import (
	"fmt"
	"time"
)

type TaskFunc func()

// Times 耗时
func Times(title string, task TaskFunc) {
	if task == nil {
		return
	}
	if title != "" {
		title = "【" + title + "】"
	}
	fmt.Println(title)
	fmt.Println("开始：" + time.Now().Format("2006-01-02 15:04:05"))
	begin := time.Now().UnixMilli()
	task()
	end := time.Now().UnixMilli()
	fmt.Println("结束：" + time.Now().Format("2006-01-02 15:04:05"))
	delta := (end - begin) / 1000
	fmt.Printf("耗时%d秒\n", delta)
	fmt.Println("-------------------------------------")
}
