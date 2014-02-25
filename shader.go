package main

import (
	"bytes"
	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"io/ioutil"
	"log"
)

// Fill the shader with the source, compile and go!
func fillShader(program gl.Program, shader gl.Shader, filename string) {
	vertexData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	vertexSource := bytes.NewBuffer(vertexData).String()
	shader.Source(vertexSource)
	shader.Compile()
	program.AttachShader(shader)
}

// Create the vertex shader
func createVertex(program gl.Program) gl.Shader {
	shader := gl.CreateShader(gl.VERTEX_SHADER)
	fillShader(program, shader, "./vertex.glsl")
	return shader
}

// create the fragment shader
func createFragment(program gl.Program) gl.Shader {
	shader := gl.CreateShader(gl.FRAGMENT_SHADER)
	fillShader(program, shader, "./fragment.glsl")
	return shader
}

func shader(window *glfw.Window) (program gl.Program) {
	program = gl.CreateProgram()
	createVertex(program)
	createFragment(program)

	// Say to which buffer shall fragment data go
	program.BindFragDataLocation(0, "finalColor")

	program.Link()

	// Get `vert` input attribute of vertex shader
	attributeLocation := program.GetAttribLocation("vert")

	attributeLocation.AttribPointer(
		3,        // amount of values per vertex
		gl.FLOAT, // type
		false,    // if not float, normalize?
		0,        // stride (how much data lays in between the vertices in the array)
		nil,      // array buffer offset
	)

	// needs to be enabled
	attributeLocation.EnableArray()

	return program
}
