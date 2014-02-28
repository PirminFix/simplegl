package main

import (
	"log"
	"math"
	"unsafe"

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
		0.0, 0.5, 1.0, 0.0, 0.0, // Vertex 1 (X, Y) Red
		0.5, -0.5, 0.0, 1.0, 0.0, // Vertex 2 (X, Y) Green
		-0.5, -0.5, 0.0, 0.0, 1.0, // Vertex 3 (X, Y) Blue
	}

	// Create Vertex Buffer Object to have some space in video ram for our vertices
	// Then upload the vertices to that buffer
	vbo := gl.GenBuffer()
	vbo.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, int(unsafe.Sizeof(vertices))*len(vertices), vertices, gl.STATIC_DRAW)

	program := shaderProgram(window)
	program.Link()
	program.Use()

	// Let's create a Vertex Array Object to save the relation of attributes and buffer object
	vao := gl.GenVertexArray()
	vao.Bind()

	// Telling opengl how our attributes are connected:
	posAttrib := program.GetAttribLocation("position")
	// describes current VBO
	posAttrib.AttribPointer(
		2,        // Amount of values for a vertex (X, Y)
		gl.FLOAT, // Type of the values
		false,    // normalize? (only if not floats)
		5*int(unsafe.Sizeof(float32(0))), // bytes between values (stride)
		nil, // offset in the array (whyever this needs to be a pointer)
	)
	posAttrib.EnableArray()

	colAttrib := program.GetAttribLocation("color")
	colAttrib.AttribPointer(
		3,
		gl.FLOAT,
		false,
		5*int(unsafe.Sizeof(float32(0))),
		2*unsafe.Sizeof(float32(0)),
	)
	colAttrib.EnableArray()

	uniColor := program.GetUniformLocation("triangleColor")

	for !window.ShouldClose() {
		// Might be used as a timer or something
		// leaving this here as a reminder of its existence
		glfw3.PollEvents()
		time := glfw3.GetTime()
		uniColor.Uniform3f(float32((math.Sin(time*4.0)+1.0)/2.0), 0.0, 0.0)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)
		window.SwapBuffers()
	}
}
