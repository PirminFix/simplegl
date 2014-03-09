// Wrapper around matrix stuff
package main

import (
	"log"
	"math"

	"github.com/skelterjohn/go.matrix"
)

func rotate(deg float64) *[16]float32 {
	const DegToRad = math.Pi / 180.0

	θ := DegToRad * deg

	rotationMatrix := matrix.Eye(4)
	rotationMatrix.Set(0, 0, math.Cos(θ))
	rotationMatrix.Set(0, 1, -math.Sin(θ))
	rotationMatrix.Set(1, 0, math.Sin(θ))
	rotationMatrix.Set(1, 1, math.Cos(θ))

	var mat32 [16]float32
	for i, v := range rotationMatrix.Array() {
		mat32[i] = float32(v)
	}
	return &mat32
}

// assumes a vector
func normalizeDense(mat *matrix.DenseMatrix) {
	norm := mat.NumElements()
	var sum float64
	sum = 0
	for i := 0; i < norm; i++ {
		sum += mat.Get(i, 0) * mat.Get(i, 0)
	}
	mat.Scale(math.Sqrt(sum))
}

func lookAt(camPos, center, upVector *matrix.DenseMatrix) [16]float64 {
	forward, err := center.TimesDense(camPos)
	if err != nil {
		log.Panic(err)
	}
	normalizeDense(forward)

	side, err := forward.TimesDense(upVector)
	if err != nil {
		log.Panic(err)
	}
	normalizeDense(side)

	up, err := side.TimesDense(forward)
	if err != nil {
		log.Panic(err)
	}
	var rawMat [16]float64
	rawMat[0] = side.Get(0, 0)
	rawMat[4] = side.Get(1, 0)
	rawMat[8] = side.Get(2, 0)
	rawMat[12] = 0.0

	rawMat[1] = up.Get(0, 0)
	rawMat[5] = up.Get(1, 0)
	rawMat[9] = up.Get(2, 0)
	rawMat[13] = 0.0

	rawMat[2] = -forward.Get(0, 0)
	rawMat[6] = -forward.Get(1, 0)
	rawMat[10] = -forward.Get(2, 0)
	rawMat[14] = 0.0

	rawMat[3] = 0.0
	rawMat[15] = 1.0

	return rawMat
}

// Return world transormation matrix
// This sets the camera angle
// formula ripped from https://www.opengl.org/wiki/GluLookAt_code
func LookAt(camPos, center, upVector [3]float64) *[16]float32 {
	cam := matrix.MakeDenseMatrix(camPos[:], 3, 1)
	cen := matrix.MakeDenseMatrix(center[:], 3, 1)
	up := matrix.MakeDenseMatrix(upVector[:], 3, 1)
	result := lookAt(cam, cen, up)
	var result32 [16]float32
	for i, v := range result {
		result32[i] = float32(v)
	}
	return &result32
}
