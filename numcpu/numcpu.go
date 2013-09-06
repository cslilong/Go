package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("numCPU=", runtime.NumCPU())
	
timeout := make(chan bool,1)
go func() {
	time.Sleep(1e9)
	timeout <- true
}

select {
	case <-ch:
	// 从ch中读数据
	case <-timeout:
	// 一直没有从ch读数据，但从timeout读到数据，说明超时了
}
}
