package main

import (
  "fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
  // note the float64() conversion here: otherwise it's inf loop
  return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
  if x >= 0 {
    z := 1.0
    for i := 0; i < 10; i++ {
      z -= (z*z - x) / (2*z)
    }
    return z, nil
  } else {
    return 0, ErrNegativeSqrt(x)
  }
}

func main() {
  fmt.Println(Sqrt(2))
  fmt.Println(Sqrt(-2))
}

