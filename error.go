package main

import "github.com/go-gl/gl"
import "log"

// glError: Handle gl errors
func glError(msg string) {
	var errMsg string
	glErr := gl.GetError()
	switch glErr {
	case gl.INVALID_ENUM:
		errMsg = "invalid enum"
	case gl.INVALID_VALUE:
		errMsg = "invalid value"
	case gl.INVALID_OPERATION:
		errMsg = "invalid Operation"
	case gl.INVALID_FRAMEBUFFER_OPERATION:
		errMsg = "invalid framebuffer operation"
	case gl.OUT_OF_MEMORY:
		errMsg = "out of memory"
	case gl.NO_ERROR:
		errMsg = ""
	default:
		errMsg = "unknown opengl error"
	}
	if errMsg != "" {
		log.Fatalf("%v: %v\n", msg, errMsg)
	}
}
