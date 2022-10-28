package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	go func() {
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
	}()
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	println(`Program Sonlandırılıyor`)
}
