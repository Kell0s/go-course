package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		img = append(img, make([]uint8, 0, dx))
	}

	return img
}

func main() {
	pic.Show(Pic)
}
