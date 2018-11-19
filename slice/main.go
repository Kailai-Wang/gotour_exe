package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
  r := make([][]uint8, dy)
  for i,_ := range r {
    r[i] = make([]uint8, dx)
  }

  for x,_ := range r {
    for y,_ := range r[x] {
      r[y][x] = uint8((x+y)/2)
    }
  }

  return r
}

func main() {
  pic.Show(Pic)
}

