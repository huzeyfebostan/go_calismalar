package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		println(`Program Sonlandırılıyor`)
		os.Exit(1)
	}()
	space := 0
	space_flag := false
	for {
		println("**********")
		time.Sleep(100 * time.Millisecond)
		if space_flag == false {
			space = space + 1
			for i := space; i >= 0; i-- {
				print(" ")
			}
		} else {
			space = space - 1
			for z := 0; z <= space; z++ {
				print(" ")
			}
			if space == 0 {
				space_flag = false
			}
		}
		if space == 10 {
			space_flag = true
		}
	}
}
