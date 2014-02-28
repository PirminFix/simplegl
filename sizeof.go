// have something than can give me sizeof of basic types
package main

import "log"

// sizeof to be meant to be used only with primitive datatypes:
func sizeof(i interface{}) int {
	switch i.(type) {
	case []float32:
		return len(i.([]float32)) * 4
	}
	log.Fatal("sizeof: unknown type")
	return -1
}
