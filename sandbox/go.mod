module github.com/mttchpmn07/cassini/sandbox

go 1.18

replace github.com/mttchpmn07/cassini/engine => ../engine

replace github.com/mttchpmn07/cassini/engine/events => ../engine/events

replace github.com/mttchpmn07/cassini/engine/primatives => ../engine/primatives

replace github.com/mttchpmn07/cassini/engine/graphics => ../engine/graphics

require (
	github.com/faiface/pixel v0.10.0
	github.com/mttchpmn07/cassini/engine v0.0.0-00010101000000-000000000000
	golang.org/x/image v0.0.0-20190523035834-f03afa92d3ff
)

require (
	github.com/Tarliton/collision2d v0.1.0 // indirect
	github.com/TrueFurby/goexplorer v0.0.1 // indirect
	github.com/faiface/glhf v0.0.0-20181018222622-82a6317ac380 // indirect
	github.com/faiface/mainthread v0.0.0-20171120011319-8b78f0a41ae3 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/go-gl/gl v0.0.0-20190320180904-bf2b1f2f34d7 // indirect
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20220712193148-63cf1f4ef61f // indirect
	github.com/go-gl/mathgl v0.0.0-20190416160123-c4601bc793c7 // indirect
	github.com/mattn/go-colorable v0.1.9 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
)
