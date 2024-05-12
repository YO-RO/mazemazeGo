package main

import (
	"github.com/fatih/color"
)

func NewTile(str string, attributes ...color.Attribute) Tile {
	if attributes == nil {
		attributes = make([]color.Attribute, 0)
	}
	return Tile{
		str:        str,
		attributes: attributes,
	}
}

type Tile struct {
	str string
	// フォントの色や背景の色、書体を指定する
	attributes []color.Attribute
}

func (base Tile) Overwrite(new Tile) Tile {
	str := base.str
	if new.str != "" {
		str = new.str
	}
	var attributes []color.Attribute
	attributes = append(attributes, base.attributes...)
	attributes = append(attributes, new.attributes...)
	return NewTile(str, attributes...)
}

func (t Tile) String() string {
	c := color.New(t.attributes...)
	return c.Sprint(t.str)
}
