package main

import (
	"github.com/mttchpmn07/cassini/engine"
)

func main() {
	app := &engine.CassiniApp{}
	engine.Run(app)
}
