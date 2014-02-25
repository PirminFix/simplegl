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
	log.Printf("%v\n", err)
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
	program.Link()

	return program
}
