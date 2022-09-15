package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var s byte = 0

func delay() {
	strings := make([]string, 1024*128)
	for i := 0; i < len(strings); i++ {
		strings[i] = strconv.Itoa(int(s))
		s++
	}
	time.Sleep(30 * time.Second)
}

func main() {
	t := time.Now()
	for i := 0; i < 1000; i++ {
		go delay()
	}
	fmt.Println(time.Since(t).Milliseconds())
	time.Sleep(600 * time.Second)
}
func main2() {
	var t = time.Now()
	var wg sync.WaitGroup
	n := 1000000
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() { time.Sleep(1 * time.Second); wg.Done() }()
		if i%30000 == 0 {
			wg.Wait()
		}
	}
	wg.Wait()
	fmt.Println(time.Since(t))
}
