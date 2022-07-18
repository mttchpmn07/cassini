package engine

import "log"

func init() {
	log.Println("Cassini Game Engine version 0")
}

func Run(app Application) {
	app.Run()
}
