package main

import (
	"movingmedian/slidingwindow"
)

func main() {
	sw := slidingwindow.New(3)
	sw.Add(1)
	sw.Add(2)
	sw.Add(3)
	sw.Add(4)
	sw.Print()
}

