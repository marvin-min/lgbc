package main

import "fmt"

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box

func (b Box) Volume() float64 {
	return b.width * b.depth * b.height
}

func (b *Box) SetColor(c Color) {
	(*b).color = c
}

func (bl BoxList) BiggestColor() Color {
	v := 0.0
	k := Color(WHITE)
	for _, b := range bl {
		if b.Volume() > v {
			v = b.Volume()
			k = b.color
		}
	}
	return k
}

func (bl BoxList) PaintItBlack() {
	for i, _ := range bl {
		(&bl[i]).SetColor(BLACK)
	}
}

func (c Color) String() string {
	string := []string{"RED","BLACK", "WHITE", "BLUE",  "YELLOW"}
	return string[c]
}

func main() {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 3, 1, YELLOW},
		Box{9, 6, 5, WHITE},
		Box{10, 10, 10, BLUE},
	}

	fmt.Println(len(boxes))
	fmt.Println(boxes[0].Volume())
	fmt.Println(boxes.BiggestColor())
	//boxes.PaintItBlack()
	fmt.Println(boxes[0].color.String())
}
