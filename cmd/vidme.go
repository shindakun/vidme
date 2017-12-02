package main

import (
	"time"

	"vidme"
)

func main() {
	for c := 0; c < 10000; c++ {
		vidme.Get(c)
		time.Sleep(time.Second * 3)
	}
}
