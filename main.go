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

func loop(window *glfw.Window) {
	vertices := []float32{
		-1.0, -1.0, 0.0, // Vertex 1
		1.0, -1.0, 0.0, // Vertex 2
		0.0, 1.0, 0.0, // Vertex 3
	}
	// Create VertexBuffer on graphics card
	vertexBuffer := gl.GenBuffer()

	// make the buffer the active buffer
	vertexBuffer.Bind(gl.ARRAY_BUFFER)

	// upload data to graphic memory
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, vertices, gl.STATIC_DRAW)
	for {
		// clear screen
		gl.Clear(gl.COLOR_BUFFER_BIT)

		// first attribute buffer: vertices
		var vertexAttrib gl.AttribLocation = 0
		vertexAttrib.EnableArray()
		vertexBuffer.Bind(gl.ARRAY_BUFFER)
		vertexAttrib.AttribPointer(
			3, // size
			gl.FLOAT,
			false, // normalized?
			0,     // stride
			nil)   // array buffer offset

		// draw the triangle
		gl.DrawArrays(gl.TRIANGLES, 0, 3)

		vertexAttrib.DisableArray()

		window.SwapBuffers()
	}
}

func main() {
	glfw.Init()
	window := CreateWindow()
	window.MakeContextCurrent()

	if gl.Init() != 0 {
		log.Fatal("Failed to init GL")
	}

	gl.ClearColor(0.0, 0.0, 0.3, 0.0)

	program := shader(window)
	program.Use()
	loop(window)
}
