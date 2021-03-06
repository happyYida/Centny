package main

import (
	"fmt"
	"org.centny.test/tcode"
	"runtime"
	"time"
)

func rtimes(c chan int, times int) {
	fmt.Println("Started Testing: Should run", times, "times.")
	var count int
	for i := 0; i < times; i++ {
		count++
		time.Sleep(time.Second)
	}
	fmt.Println("Finished:", times, " times")
	c <- 1
}

func main() {
	for i := 0; i < 1000; i++ {
		time.Sleep(1000)
	}
	runtime.GOMAXPROCS(4)
	tcode.ShowChinese()
	const C_SIZE int = 3
	var chs [C_SIZE]chan int
	for i := 0; i < C_SIZE; i++ {
		chs[i] = make(chan int)
	}
	for i := C_SIZE; i > 0; i-- {
		go rtimes(chs[i-1], i*10000000)
	}
	for i := 0; i < C_SIZE; i++ {
		<-chs[i]
	}
}

type TestInc interface {
	showMsg()
}

func testList(inf TestInc) {

}
