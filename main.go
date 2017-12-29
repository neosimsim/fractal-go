package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"math/rand"
	"os"
	"time"
)

const (
	width  = 256
	height = 256
)

func writeImage(img image.Image) {
	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

type cPoint struct {
	color color.Color
	x     int
	y     int
}

// Convert the coordination of a point in a Canvac
// c = [0,n]x[0,m] point within a pane p= [a,b]x[c,d]
//
// No sanity checks are done, i.e a < c etc.
func (point *cPoint) toComplex(a, b, c, d float64) complex128 {
	pWidth := b - a
	pHeigth := d - c
	re := float64(point.x)/float64(width)*pWidth + a
	im := float64(point.y)/float64(height)*pHeigth + c
	return complex(re, im)
}

func isInSet(c complex128) bool {
	return cmplx.Abs(c) < 0.8
}

func do(x, y int, ch chan *cPoint) {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	p := new(cPoint)
	p.x = x
	p.y = y
	c := p.toComplex(-1, 1, -1, 1)
	if isInSet(c) {
		p.color = color.Black
		ch <- p
	} else {
		ch <- nil
	}
}

func calcPoint(width, heigt int, out chan cPoint) {
	ch := make(chan *cPoint, width*height)
	for y := 0; y < width; y++ {
		for x := 0; x < width; x++ {
			go do(x, y, ch)
		}
	}
	for i := 0; i < width*height; i++ {
		p := <-ch
		if p != nil {
			out <- *p
		}
	}
	close(out)
}

func main() {
	const width, height = 256, 256
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	writeImage(img)

	c := make(chan cPoint)

	go calcPoint(width, height, c)

	for point := range c {
		img.Set(point.x, point.y, point.color)
		fmt.Println(point.x, point.y)
	}

	writeImage(img)
}
