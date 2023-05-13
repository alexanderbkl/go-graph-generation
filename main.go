package main

import (
	"bytes"
	//"math"

	//"encoding/base64"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func Pic(dx, dy int) *bytes.Buffer {
	// Generate the image
	m := image.NewGray(image.Rect(-dx, -dy, dx, dy))
	for i := -dx; i < dx; i++ {
		for j := -dy; j < dy; j++ {
			if j == 0 {
				m.Set(i, j, color.Gray{Y: uint8(255)})
			}
			if i == 0 {
				m.Set(i, j, color.Gray{Y: uint8(255)})
			}
			if i == j {
				m.Set(i, j, color.Gray{Y: uint8(255)})
			}
			if i == -j {
				m.Set(i, j, color.Gray{Y: uint8(255)})
			}

			_, err := generateGraph(i, j)

			if !err {
					m.Set(i, j, color.Gray{Y: uint8(255)})
				
			}

						

		}
	}

	// Encode the image to PNG
	buf := new(bytes.Buffer)
	err := png.Encode(buf, m)
	if err != nil {
		log.Fatal(err)
	}

	return buf
}

func generateGraph(x int, y int) (float64, bool) {
	const (
		threshold = 0.1
		delta     = 0.2
	)
	var (
		x1 = float64(x) * delta
		y1 = float64(y) * delta
	)

	equation := x1*x1
	d := equation + y1

	if d < 0 {
		d = -d
	}

	if d < threshold {
		return d, false
	} else {
		return d, true
	}
}

func main() {

	// Encode the image to PNG
	buf := Pic(2000, 2000)

	//str := base64.StdEncoding.EncodeToString(buf.Bytes())

	// Write the binary data to a file with a ".png" extension
	err := os.WriteFile("image.png", buf.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Print the Base64-encoded string to the console
	//log.Print(str)
}
