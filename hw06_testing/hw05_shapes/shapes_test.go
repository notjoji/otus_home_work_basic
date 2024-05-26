package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ShapeType int8

const (
	CIRCLE ShapeType = iota
	RECTANGLE
	TRIANGLE
	UNKNOWN
)

func TestCalculateArea(t *testing.T) {
	testCases := []struct {
		name      string
		shapeType ShapeType
		radius    float64
		width     float64
		height    float64
		base      float64
		area      float64
	}{
		{
			name:      "circle_area",
			shapeType: CIRCLE,
			radius:    5,
			width:     0,
			height:    0,
			base:      0,
			area:      78.53981633974483,
		},
		{
			name:      "rectangle_area",
			shapeType: RECTANGLE,
			radius:    0,
			width:     10,
			height:    5,
			base:      0,
			area:      50,
		},
		{
			name:      "triangle_area",
			shapeType: TRIANGLE,
			radius:    0,
			width:     0,
			height:    5,
			base:      10,
			area:      25,
		},
		{
			name:      "not_a_shape",
			shapeType: UNKNOWN,
			radius:    0,
			width:     0,
			height:    0,
			base:      0,
			area:      0,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var object any
			switch testCase.shapeType {
			case CIRCLE:
				object = Circle{testCase.radius}
			case RECTANGLE:
				object = Rectangle{testCase.width, testCase.height}
			case TRIANGLE:
				object = Triangle{testCase.base, testCase.height}
			case UNKNOWN:
				object = "not object"
			}
			expected, err := calculateArea(object)
			if err != nil {
				assert.Equal(t, "переданный объект не является фигурой", err.Error())
			} else {
				assert.Equal(t, testCase.area, expected)
			}
		})
	}
}
