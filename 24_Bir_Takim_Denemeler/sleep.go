package main

import (
	"fmt"
	"time"
)

func Sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	fmt.Println("Program başladı.")
	Sleep(2 * time.Second)
	fmt.Println("2 saniye sonra bu yazı ekranda görünecek.")
}
