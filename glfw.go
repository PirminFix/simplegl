package main

import "github.com/go-gl/glfw3"
import "log"

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

func glfwStuff() *glfw3.Window {
	// First step: We need a window with opengl context
	if !glfw3.Init() {
		log.Fatal("glfw initialization failed")
	}
	glfw3.SetErrorCallback(glfwErrorCallback)

	window, err := glfw3.CreateWindow(1024, 768, NAME, nil, nil)
	if err != nil {
		log.Fatalf("CreateWindow failed: %v\n", err)
	}
	window.MakeContextCurrent()
	window.SetKeyCallback(glfwKeyCallback)
	return window
}
