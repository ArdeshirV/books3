package main

import (
	"fmt"
	"sync"
	"time"
)

func mainX() {
	stepSix()
}

func stepSix() {
	time.Sleep(time.Millisecond * 100)
	const numWorkers = 4

	for range 1024 {
		dataChan := make(chan rune, 4)
		var wg sync.WaitGroup
		colors := [4]string{"\033[1;34m", "\033[1;32m", "\033[1;33m", "\033[1;35m"}
		wg.Add(numWorkers)
		for i := range numWorkers {
			go func(index int) {
				defer wg.Done()
				for data := range dataChan {
					time.Sleep(time.Millisecond * 100)
					fmt.Print(colors[index], string(data), "\033[0m ")
				}
			}(i)
		}

		go func() {
			for c := '6'; c <= 'Z'; c++ {
				dataChan <- c
			}
			close(dataChan)
		}()

		wg.Wait()
		fmt.Println()
	}
}


