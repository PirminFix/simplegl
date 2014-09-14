// Wrapper around matrix stuff
package main

import (
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

func NormalizeVector(v *[3]float64) {
	sum := math.Pow(v[0], 2) +
		math.Pow(v[1], 2) +
		math.Pow(v[2], 2)
	f := math.Sqrt(sum)
	v[0] *= f
	v[1] *= f
	v[2] *= f
}

// ComputeNormalOfPlnaeFLOAT_2 from glhlib
// 2 pvectors are passed and normal is computed
func ComputeNormalOfPlane(normal, pvector1, pvector2 *[3]float64) {
	normal[0] = (pvector1[1] * pvector2[2]) - (pvector1[2] * pvector2[1])
	normal[1] = (pvector1[2] * pvector2[0]) - (pvector1[0] * pvector2[2])
	normal[2] = (pvector1[0] * pvector2[1]) - (pvector1[1] * pvector2[0])
}

// Return world transormation matrix
// This sets the camera angle
// formula ripped from https://www.opengl.org/wiki/GluLookAt_code
func LookAtf2(
	matrix *[16]float64,
	eyePosition3D, center3D, upVector3D *[3]float64,
) {
	var forward, side, up [3]float64

	forward[0] = center3D[0] - eyePosition3D[0]
	forward[1] = center3D[1] - eyePosition3D[1]
	forward[2] = center3D[2] - eyePosition3D[2]
	NormalizeVector(&forward)
	//------------------
	//Side = forward x up
	ComputeNormalOfPlane(&side, &forward, upVector3D)
	NormalizeVector(&side)
	//------------------
	//Recompute up as: up = side x forward
	ComputeNormalOfPlane(&up, &side, &forward)
	//------------------
	matrix[0] = side[0]
	matrix[4] = side[1]
	matrix[8] = side[2]
	matrix[12] = 0.0
	//------------------
	matrix[1] = up[0]
	matrix[5] = up[1]
	matrix[9] = up[2]
	matrix[13] = 0.0
	//------------------
	matrix[2] = -forward[0]
	matrix[6] = -forward[1]
	matrix[10] = -forward[2]
}

// ripped von glhlib/MathLibrary.h
// PURPOSE:	For square matrices. This is column major for OpenGL
func MultiplyMatrices4by4OpenGL_FLOAT(
	result, matrix1, matrix2 *[16]float64,
) {
	result[0] = matrix1[0]*matrix2[0] +
		matrix1[4]*matrix2[1] +
		matrix1[8]*matrix2[2] +
		matrix1[12]*matrix2[3]
	result[4] = matrix1[0]*matrix2[4] +
		matrix1[4]*matrix2[5] +
		matrix1[8]*matrix2[6] +
		matrix1[12]*matrix2[7]
	result[8] = matrix1[0]*matrix2[8] +
		matrix1[4]*matrix2[9] +
		matrix1[8]*matrix2[10] +
		matrix1[12]*matrix2[11]
	result[12] = matrix1[0]*matrix2[12] +
		matrix1[4]*matrix2[13] +
		matrix1[8]*matrix2[14] +
		matrix1[12]*matrix2[15]

	result[1] = matrix1[1]*matrix2[0] +
		matrix1[5]*matrix2[1] +
		matrix1[9]*matrix2[2] +
		matrix1[13]*matrix2[3]
	result[5] = matrix1[1]*matrix2[4] +
		matrix1[5]*matrix2[5] +
		matrix1[9]*matrix2[6] +
		matrix1[13]*matrix2[7]
	result[9] = matrix1[1]*matrix2[8] +
		matrix1[5]*matrix2[9] +
		matrix1[9]*matrix2[10] +
		matrix1[13]*matrix2[11]
	result[13] = matrix1[1]*matrix2[12] +
		matrix1[5]*matrix2[13] +
		matrix1[9]*matrix2[14] +
		matrix1[13]*matrix2[15]

	result[2] = matrix1[2]*matrix2[0] +
		matrix1[6]*matrix2[1] +
		matrix1[10]*matrix2[2] +
		matrix1[14]*matrix2[3]
	result[6] = matrix1[2]*matrix2[4] +
		matrix1[6]*matrix2[5] +
		matrix1[10]*matrix2[6] +
		matrix1[14]*matrix2[7]
	result[10] = matrix1[2]*matrix2[8] +
		matrix1[6]*matrix2[9] +
		matrix1[10]*matrix2[10] +
		matrix1[14]*matrix2[11]
	result[14] = matrix1[2]*matrix2[12] +
		matrix1[6]*matrix2[13] +
		matrix1[10]*matrix2[14] +
		matrix1[14]*matrix2[15]

	result[3] = matrix1[3]*matrix2[0] +
		matrix1[7]*matrix2[1] +
		matrix1[11]*matrix2[2] +
		matrix1[15]*matrix2[3]
	result[7] = matrix1[3]*matrix2[4] +
		matrix1[7]*matrix2[5] +
		matrix1[11]*matrix2[6] +
		matrix1[15]*matrix2[7]
	result[11] = matrix1[3]*matrix2[8] +
		matrix1[7]*matrix2[9] +
		matrix1[11]*matrix2[10] +
		matrix1[15]*matrix2[11]
	result[15] = matrix1[3]*matrix2[12] +
		matrix1[7]*matrix2[13] +
		matrix1[11]*matrix2[14] +
		matrix1[15]*matrix2[15]
}

func glhLookAtf2(
	matrix *[16]float64,
	eyePosition3D, center3D, upVector3D *[3]float64,
) {
	var (
		forward, side, up     *[3]float64
		matrix2, resultMatrix *[16]float64
	)

	forward[0] = center3D[0] - eyePosition3D[0]
	forward[1] = center3D[1] - eyePosition3D[1]
	forward[2] = center3D[2] - eyePosition3D[2]
	NormalizeVector(forward)

	//Side = forward x up
	ComputeNormalOfPlane(side, forward, upVector3D)
	NormalizeVector(side)

	//Recompute up as: up = side x forward
	ComputeNormalOfPlane(up, side, forward)

	matrix2[0] = side[0]
	matrix2[4] = side[1]
	matrix2[8] = side[2]
	matrix2[12] = 0.0

	matrix2[1] = up[0]
	matrix2[5] = up[1]
	matrix2[9] = up[2]
	matrix2[13] = 0.0

	matrix2[2] = -forward[0]
	matrix2[6] = -forward[1]
	matrix2[10] = -forward[2]
	matrix2[14] = 0.0

	matrix2[3] = 0.0
	matrix2[7] = 0.0
	matrix2[11] = 0.0
	matrix2[15] = 1.0

	MultiplyMatrices4by4OpenGL_FLOAT(resultMatrix, matrix, matrix2)
	glhTranslatef2(resultMatrix, -eyePosition3D[0], -eyePosition3D[1], -eyePosition3D[2])

	memcpy(matrix, resultMatrix, 16*sizeof(sreal))
}
