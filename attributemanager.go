package main

import (
	"fmt"
	"unsafe"

	"github.com/go-gl/gl"
)
import "log"

type AttributeManager struct {
	stride  int
	offset  uintptr
	program gl.Program
	names   []string
	values  []uint
}

func NewAttributeManager(program gl.Program) *AttributeManager {
	am := AttributeManager{
		program: program,
		names:   nil,
		values:  nil,
	}
	return &am
}

func (am *AttributeManager) Add(name string, values uint) {
	log.Printf("am.Set(%s, ...)", name)
	am.names = append(am.names, name)
	am.values = append(am.values, values)
	am.stride += int(values)
}

func (am *AttributeManager) Set() {
	var offset uintptr
	offset = 0
	log.Printf("am.names == %v", am.names)
	for i := 0; i < len(am.names); i++ {
		log.Printf("am.Exec(%s, ...)", am.names[i])
		am.setAttrib(
			am.names[i],
			am.values[i],
			offset,
		)
		offset += uintptr(am.values[i])
	}

}

func (am *AttributeManager) setAttrib(
	name string,
	values uint,
	offset uintptr,
) {
	log.Printf("setAttrib(%s, ...)", name)
	attrib := am.program.GetAttribLocation(name)
	if attrib < 0 {
		log.Fatalf("%v < 0", name)
	}
	attrib.EnableArray()
	glError(fmt.Sprintf("%v enable attrib array", name))
	attrib.AttribPointer(
		values,   // Amount of values for a vertex (X, Y)
		gl.FLOAT, // Type of the values
		false,    // normalize? (only if not floats)
		am.stride*int(unsafe.Sizeof(float32(0))), // bytes between values (stride)
		offset*unsafe.Sizeof(float32(0)),         // offset in the array (whyever this needs to be a pointer)
	)
	glError(fmt.Sprintf("%v attrib error", name))
}
