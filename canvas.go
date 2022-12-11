package main

import "github.com/lxn/walk"

type (
	Canvas struct {
		*walk.Canvas
	}
)

func NewCanvas() *Canvas {
	return new(Canvas)
}
