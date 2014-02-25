package main

import (
	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"log"
)

func CreateWindow() *glfw.Window {
	window, err := glfw.CreateWindow(1024, 768, "Testing", nil, nil)
	if err != nil {
		log.Fatal("Failed to open GLFW window: " + err.Error())
	}
	return window
}

func loop(window *glfw.Window, vertexBuffer gl.Buffer) {

	log.Print("bind Vertex buffer")
	vertexBuffer.Bind(gl.ARRAY_BUFFER)

	log.Print("really loop now")
	for !window.ShouldClose() {
		// clear screen
		gl.Clear(gl.COLOR_BUFFER_BIT)

		log.Print("Draw")
		// draw the triangle
		// FIXME here it crashes
		gl.DrawArrays(
			gl.TRIANGLES, // We want a triangle
			0,            // skip that many vertices from the beginning
			3,            // how many vertices to process
		)

		log.Print("Swap buffers")
		window.SwapBuffers()
	}
}

// Init OpenGL and a window
func initGl() *glfw.Window {
	glfw.Init()
	window := CreateWindow()
	window.MakeContextCurrent()

	if gl.Init() != 0 {
		log.Fatal("Failed to init GL")
	}
	return window
}

// Fill the vertex buffer wiuth data
func fillVBO(vertices []float32) (vertexBuffer gl.Buffer) {
	// Create VertexBuffer on graphics card
	log.Print("generating buffer")
	vertexBuffer = gl.GenBuffer()

	// make the buffer the active buffer
	log.Print("binding buffer")
	vertexBuffer.Bind(gl.ARRAY_BUFFER)

	// upload data to graphic memory
	log.Print("uploading data")
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, vertices, gl.STATIC_DRAW)

	log.Print("returning")
	return vertexBuffer
}

func main() {
	log.Print("initilaizing window")
	window := initGl()

	log.Print("Building shader")
	program := shader(window)

	program.Use()

	log.Print("Clear color")
	gl.ClearColor(0.0, 0.0, 0.3, 0.0)

	vao := gl.GenVertexArray()
	vao.Bind()

	vertices := []float32{
		-1.0, -1.0, 0.0, // Vertex 1
		1.0, -1.0, 0.0, // Vertex 2
		0.0, 1.0, 0.0, // Vertex 3
	}

	log.Print("fill vertex buffer")
	vertexBuffer := fillVBO(vertices)

	log.Print("loop")
	loop(window, vertexBuffer)
}
