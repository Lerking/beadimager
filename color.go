package main

import "github.com/lxn/walk"

type (
	BeadColor struct {
		Checkbox        *walk.CheckBox
		backgroundColor walk.Color
		Red             uint8
		Green           uint8
		Blue            uint8
	}
)
