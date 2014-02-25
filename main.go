package main

import (
	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"log"
)

func main() {
	glfw.Init()
	window, err := glfw.CreateWindow(
		1024,       // width
		768,        // height
		"simplegl", // title
		nil,        // monitor
		nil,        // window to share context with
	)
	if err != nil {
		log.Fatalf("Failed to open GLFW window: %v\n", err.Error())
	}
	defer window.Destroy()

	glInitErr := gl.Init()
	if glInitErr != 0 {
		log.Fatalf("Failed to init GL: %v\n", glInitErr)
	}
	window.MakeContextCurrent()
	defer glfw.Terminate()

	// TODO
	//shader := shaderProgram(window)
	//shader.Link()
	//shader.Use()

	for !window.ShouldClose() {
		// TODO
	}
}
