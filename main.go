package main

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/go-gl/gl"
	"github.com/go-gl/glfw3"
)

const NAME = "simplegl"
const VERSION = "0.0.1"

const TEXTUREFILE = "./texture.png"

func genVao() gl.VertexArray {
	// Let's create a Vertex Array Object to save the relation of attributes and buffer object
	vao := gl.GenVertexArray()
	if vao < 0 {
		log.Fatal("vao < 0")
	}
	vao.Bind()
	return vao
}

// Create Vertex Buffer Object to have some space in video ram for our vertices
func genVbo() gl.Buffer {
	vbo := gl.GenBuffer()
	if vbo < 0 {
		log.Fatal("vbo < 0")
	}

	vbo.Bind(gl.ARRAY_BUFFER)
	// Triangle
	vertices := []float32{
		// Position, Color, Texcoords
		-0.5, 0.5, 1.0, 0.0, 0.0, 0.0, 0.0, // Top-left
		0.5, 0.5, 0.0, 1.0, 0.0, 1.0, 0.0, // Top-right
		0.5, -0.5, 0.0, 0.0, 1.0, 1.0, 1.0, // Bottom-right
		-0.5, -0.5, 1.0, 1.0, 1.0, 0.0, 1.0, // Bottom-left
	}
	gl.BufferData(
		gl.ARRAY_BUFFER,
		int(unsafe.Sizeof(vertices))*len(vertices),
		vertices,
		gl.STATIC_DRAW,
	)
	glError("fill vbo with data failed")
	return vbo
}

func genEbo() gl.Buffer {
	ebo := gl.GenBuffer()
	if ebo < 0 {
		log.Fatal("ebo < 0")
	}
	ebo.Bind(gl.ELEMENT_ARRAY_BUFFER)
	elements := []gl.GLuint{
		0, 1, 2,
		2, 3, 0,
	}
	gl.BufferData(
		gl.ELEMENT_ARRAY_BUFFER,
		int(unsafe.Sizeof(elements))*len(elements),
		elements,
		gl.STATIC_DRAW,
	)
	glError("fill ebo with data failed")
	return ebo
}

func genTex() {
	tex := gl.GenTexture()
	if tex < 0 {
		log.Fatal("tex < 0")
	}
	tex.Bind(gl.TEXTURE_2D)
	glError("texgen")
	//pixels, imgWidth, imgHeight := png2array(TEXTUREFILE)
	pixels := []uint32{
		1.0, 1.0, 1.0, 0.0, 0.0, 0.0,
		0.0, 0.0, 0.0, 1.0, 1.0, 1.0,
	}
	imgWidth := 2
	imgHeight := 2
	gl.TexImage2D(
		gl.TEXTURE_2D, // work on 2d texture
		0,             // Level of detail
		gl.RGB,        // Format for the gpu
		imgWidth,      // width
		imgHeight,     // height
		0,             // border, always 0
		gl.RGB,        // format of the image
		//gl.UNSIGNED_INT, // datatype of the image
		gl.FLOAT,
		pixels, // image array
	)
	defer tex.Delete()
	glError("tex")
	// make this mirrored wrap!
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.MIRRORED_REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.MIRRORED_REPEAT)
	glError("wrap")
	//gl.GenerateMipmap(gl.TEXTURE_2D)
	//gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST_MIPMAP_NEAREST)
	//gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST_MIPMAP_NEAREST)
	//glError("mip")
}

func setAttrib(name string, program gl.Program, values uint, stride int, offset uintptr) {
	attrib := program.GetAttribLocation(name)
	if attrib < 0 {
		log.Fatalf(fmt.Sprint("%v < 0", name))
	}
	attrib.EnableArray()
	attrib.AttribPointer(
		values,   // Amount of values for a vertex (X, Y)
		gl.FLOAT, // Type of the values
		false,    // normalize? (only if not floats)
		stride*int(unsafe.Sizeof(float32(0))), // bytes between values (stride)
		offset*unsafe.Sizeof(float32(0)),      // offset in the array (whyever this needs to be a pointer)
	)
}

func setPosAttrib(program gl.Program) {
	// Telling opengl how our attributes are connected:
	posAttrib := program.GetAttribLocation("position")
	posAttrib.EnableArray()
	// describes current VBO
	posAttrib.AttribPointer(
		2,        // Amount of values for a vertex (X, Y)
		gl.FLOAT, // Type of the values
		false,    // normalize? (only if not floats)
		7*int(unsafe.Sizeof(float32(0))), // bytes between values (stride)
		nil, // offset in the array (whyever this needs to be a pointer)
	)
	glError("posAttrib")
}
func setColAttrib(program gl.Program) {
	colAttrib := program.GetAttribLocation("color")
	colAttrib.EnableArray()
	colAttrib.AttribPointer(
		3,
		gl.FLOAT,
		false,
		7*int(unsafe.Sizeof(float32(0))),
		2*unsafe.Sizeof(float32(0)),
	)
	glError("colAttrib")
}

func setTexAttrib(program gl.Program) {
	texAttrib := program.GetAttribLocation("texcoord")
	if texAttrib == -1 {
		log.Fatal(`GetAttribLocation("texcoord") returned -1`)
	}
	glError(fmt.Sprintf("texAttribA %v", texAttrib))
	texAttrib.EnableArray()
	glError(fmt.Sprintf("texAttribB %v", texAttrib))
	texAttrib.AttribPointer(
		2,
		gl.FLOAT,
		false,
		7*int(unsafe.Sizeof(float32(0))),
		5*unsafe.Sizeof(float32(0)),
	)
	glError("texAttribC")
}

func main() {

	window := glfwStuff()
	defer glfw3.Terminate()
	defer window.Destroy()
	gl.Init()

	genVao()
	genVbo()
	genEbo()

	program := shaderProgram(window)
	program.Link()
	program.Use()
	glError("program")

	setPosAttrib(program)
	//bla := program.GetAttribLocation("bla")
	//if bla == -1 {
	//	log.Fatal("no bla :(")
	//}
	setColAttrib(program)
	//setTexAttrib(program)

	genTex()

	for !window.ShouldClose() {
		// Might be used as a timer or something
		// leaving this here as a reminder of its existence
		glfw3.PollEvents()
		gl.ClearColor(0.0, 0.0, 0.99999, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)
		//time := glfw3.GetTime()
		//gl.DrawArrays(gl.TRIANGLES, 0, 6)
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)
		window.SwapBuffers()
	}
}
