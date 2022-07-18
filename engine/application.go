package engine

import "log"

type Application interface {
	Run()
}

type CassiniApp struct {
}

func (c CassiniApp) Run() {
	log.Println("Hello from my cassini app!")
}
