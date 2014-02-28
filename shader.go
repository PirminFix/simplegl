package main

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
)

// Fill the shader with the source, compile and go!
func fillShader(program gl.Program, shader gl.Shader, filename string) {
	shaderData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	shaderSource := bytes.NewBuffer(shaderData).String()
	shader.Source(shaderSource)
	shader.Compile()
	if shader.Get(gl.COMPILE_STATUS) != gl.TRUE {
		shaderLog := shader.GetInfoLog()
		log.Fatalf("Compiling shader %v failed!\n%v\n", filename, shaderLog)
	}
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

// Create and return shader program from hardcoded files.
// It is not yet linked.
func shaderProgram(window *glfw.Window) (program gl.Program) {
	program = gl.CreateProgram()
	createVertex(program)
	createFragment(program)
	return program
}
