package main

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.
//
// Copyrigth © Alexander Ben Nasrallah 2017 <abn@posteo.de>

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"math/cmplx"
	"math/rand"
	"os"
	"time"
)

const (
	width  = 1155
	height = 1050
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
	iterations int
	x          int
	y          int
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

func isInSet(c complex128) int {
	var z complex128 = 0

	iter := 0
	maxIter := 100
	for cmplx.Abs(z) < 2 && iter < maxIter {
		z = cmplx.Pow(z, 2) + c
		iter++
	}
	return iter
}

func do(x, y int, ch chan *cPoint) {
	p := new(cPoint)
	p.x = x
	p.y = y
	c := p.toComplex(-2.04, 0.6, -1.2, 1.2)
	p.iterations = isInSet(c)
	time.Sleep(time.Duration(rand.Intn(p.iterations)) * time.Nanosecond)
	ch <- p
}

func calcPoint(width, heigt int, out chan cPoint) {
	ch := make(chan *cPoint, width*height)
	go func() {
		for y := 0; y < height/2; y++ {
			for x := 0; x < width; x++ {
				go do(x, y, ch)
			}
		}
	}()
	go func() {
		for y := height - 1; height/2 <= y; y-- {
			for x := 0; x < width; x++ {
				go do(x, y, ch)
			}
		}
	}()
	for i := 0; i < width*height; i++ {
		p := <-ch
		if p != nil {
			out <- *p
		}
	}
	close(out)
}

func main() {
	c := make(chan cPoint)

	go calcPoint(width, height, c)

	for point := range c {
		fmt.Println(point.x, point.y, point.iterations)
	}
}
