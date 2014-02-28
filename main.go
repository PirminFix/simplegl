package main

import (
	"log"

	"github.com/go-gl/gl"
	"github.com/go-gl/glfw3"
)

const NAME = "simplegl"
const VERSION = "0.0.1"

// callback handler for glfw errors
func glfwErrorCallback(error glfw3.ErrorCode, description string) {
	log.Fatalf("glfw error: %v: %v\n", error, description)
}

// callback handler for glfw key events
func glfwKeyCallback(window *glfw3.Window, key glfw3.Key, scancode int, action glfw3.Action, mods glfw3.ModifierKey) {
	if key == glfw3.KeyEscape && action == glfw3.Press {
		window.SetShouldClose(true)
	}
}

func main() {
	// First step: We need a window with opengl context
	if !glfw3.Init() {
		log.Fatal("glfw initialization failed")
	}
	defer glfw3.Terminate()
	glfw3.SetErrorCallback(glfwErrorCallback)

	window, err := glfw3.CreateWindow(1024, 768, NAME, nil, nil)
	if err != nil {
		log.Fatalf("CreateWindow failed: %v\n", err)
	}
	defer window.Destroy()
	window.MakeContextCurrent()
	window.SetKeyCallback(glfwKeyCallback)

	gl.Init()

	// Triangle
	vertices := []float32{
		0.0, 0.5, // Vertex 1 (X, Y)
		0.5, -0.5, // Vertex 2 (X, Y)
		-0.5, -0.5, // Vertex 3 (X, Y)
	}

	// Create Vertex Buffer Object to have some space in video ram for our vertices
	// Then upload the vertices to that buffer
	vbo := gl.GenBuffer()
	vbo.Bind(gl.ARRAY_BUFFER)
	// each float32 consumes 4 bytes
	gl.BufferData(gl.ARRAY_BUFFER, sizeof(vertices), vertices, gl.STATIC_DRAW)

	for !window.ShouldClose() {
		// Might be used as a timer or something
		// leaving this here as a reminder of its existence
		//time := glfw3.GetTime()
		glfw3.PollEvents()
		window.SwapBuffers()
	}
}
