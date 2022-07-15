package golang_united_school_homework

import (
	"errors"
)

const (
	notInBound = "shape by index doesn't exist or index went out of the range"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if b.isBoxFull() {
		return errors.New("reach limit of box capacity")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if b.isIndexExist(i) {
		return b.shapes[i], nil
	}
	return nil, errors.New(notInBound)

}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if b.isIndexExist(i) {
		extractedShape := b.shapes[i]
		b.removeElement(i)
		return extractedShape, nil
	}
	return nil, errors.New(notInBound)
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if b.isIndexExist(i) {
		extractedShape := b.shapes[i]
		b.shapes[i] = shape
		return extractedShape, nil
	}
	return nil, errors.New(notInBound)
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var summ float64
	for _, v := range b.shapes {
		summ += v.CalcPerimeter()
	}
	return summ
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var summ float64
	for _, v := range b.shapes {
		summ += v.CalcArea()
	}
	return summ
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var deletedCount int = 0
	shapesCopy := make([]Shape, len(b.shapes))
	copy(shapesCopy, b.shapes)
	for k, v := range shapesCopy {
		if v.ShapeType() == "Circle" {
			b.removeElement(k - deletedCount)
			deletedCount += 1
		}
	}
	if deletedCount == 0 {
		return errors.New("no circles to delete")
	}
	return nil
}

func (b *box) removeElement(i int) error {
	if b.isIndexExist(i) {
		b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
		return nil
	}
	return errors.New("shape by index doesn't exist or index went out of the range")
}

func (b *box) isBoxFull() bool {
	return len(b.shapes) >= b.shapesCapacity
}

func (b *box) isIndexExist(i int) bool {
	return len(b.shapes) > i
}
