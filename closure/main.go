package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
  arr := make([]int, 0)
  var i int

	return func() int {
    
    switch i {
      case 0:
        arr = append(arr, 1)
      case 1:
        arr = append(arr, 1)
      default:
        arr = append(arr, arr[i-1] + arr[i-2])
    }
    i++
    return arr[len(arr) - 1]
  }
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
