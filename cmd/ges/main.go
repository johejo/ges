package main

import (
	"github.com/johejo/gohejo/logutils"
)

var logger = logutils.New()

func main() {
	logger.Print("wakeup")
	s := NewServer()
	logger.Fatalln(s.Run())
}
