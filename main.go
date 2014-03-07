package main

import (
	"log"
	"unsafe"

	"github.com/go-gl/gl"
	"github.com/go-gl/glfw3"
)

const NAME = "simplegl"
const VERSION = "0.0.1"

const TEXTUREFILE = "./texture.jpg"

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
		-0.5, 0.5, 0.0, 0.0, // Top-left
		0.5, 0.5, 1.0, 0.0, // Top-right
		0.5, -0.5, 1.0, 1.0, // Bottom-right
		-0.5, -0.5, 0.0, 1.0, // Bottom-left
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
	pixels, imgWidth, imgHeight := png2array(TEXTUREFILE)
	gl.TexImage2D(
		gl.TEXTURE_2D, // work on 2d texture
		0,             // Level of detail
		gl.RGBA,       // Format for the gpu
		imgWidth,      // width
		imgHeight,     // height
		0,             // border, always 0
		gl.RGBA,       // format of the image
		//gl.UNSIGNED_INT, // datatype of the image
		gl.FLOAT,
		pixels, // image array
	)
	glError("tex")
	// make this mirrored wrap!
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.MIRRORED_REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.MIRRORED_REPEAT)
	glError("wrap")
	//gl.GenerateMipmap(gl.TEXTURE_2D)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	glError("mip")
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

	log.Print("Set Attributes")
	am := NewAttributeManager(program)
	am.Add("position", 2)
	//am.Add("color", 3)
	am.Add("texcoord", 2)
	am.Set()

	// Give us some dog and cat textures
	catPix, catWidth, catHeight := png2array("./cat.jpg")
	dogPix, dogWidth, dogHeight := png2array("./puppy.png")

	catTex := gl.GenTexture()
	dogTex := gl.GenTexture()

	gl.ActiveTexture(gl.TEXTURE0)
	glError("active texture")
	catTex.Bind(gl.TEXTURE_2D)
	glError("bind texture")
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		catWidth,
		catHeight,
		0,
		gl.RGBA,
		gl.FLOAT,
		catPix,
	)
	glError("upload cat")
	uniLocTexCat := program.GetUniformLocation("texCat")
	glError("cat uniform location")
	uniLocTexCat.Uniform1i(0)

	glError("tex")
	// make this mirrored wrap!
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.MIRRORED_REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.MIRRORED_REPEAT)
	glError("wrap")
	gl.GenerateMipmap(gl.TEXTURE_2D)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	glError("mip")

	gl.ActiveTexture(gl.TEXTURE1)
	dogTex.Bind(gl.TEXTURE_2D)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		dogWidth,
		dogHeight,
		0,
		gl.RGBA,
		gl.FLOAT,
		dogPix,
	)
	uniLocTexDog := program.GetUniformLocation("texDog")
	uniLocTexDog.Uniform1i(1)

	glError("tex")
	// make this mirrored wrap!
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.MIRRORED_REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.MIRRORED_REPEAT)
	glError("wrap")
	gl.GenerateMipmap(gl.TEXTURE_2D)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	glError("mip")

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
