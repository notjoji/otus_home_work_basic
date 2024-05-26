package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

func calculateArea(s any) (any, error) {
	switch obj := s.(type) {
	case Shape:
		return obj.Area(), nil
	default:
		return nil, fmt.Errorf("переданный объект не является фигурой")
	}
}

func main() {
	var object Shape
	object = Circle{5}
	circleArea, err := calculateArea(object)
	if err == nil {
		fmt.Println("Круг: радиус", object.(Circle).Radius, "Площадь", circleArea)
	}
	object = Rectangle{10, 5}
	rectangleArea, err := calculateArea(object)
	if err == nil {
		fmt.Println("Прямоугольник: ширина", object.(Rectangle).Width, "высота", object.(Rectangle).Height,
			"Площадь", rectangleArea)
	}
	object = Triangle{8, 6}
	triangleArea, err := calculateArea(object)
	if err == nil {
		fmt.Println("Треугольник: ширина", object.(Triangle).Base, "высота", object.(Triangle).Height,
			"Площадь", triangleArea)
	}
	_, err = calculateArea("not object")
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
}
