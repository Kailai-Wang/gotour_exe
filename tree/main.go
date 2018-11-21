package main

import "golang.org/x/tour/tree"
import "fmt"
import "strings"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
  if t == nil {
    return
  }
  
  Walk(t.Left, ch)
  ch <- t.Value
  Walk(t.Right, ch)
}

var trimfunc = func (r rune) rune {
  if r == '(' || r == ')' {
    return -1
  }
    return r
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
  s1 := strings.Map(trimfunc, t1.String())
  s2 := strings.Map(trimfunc, t2.String())
  return s1 == s2
}

func main() {
  t := tree.New(1)
  ch := make(chan int)
  defer close(ch)
  
  go Walk(t, ch)
  for i := 0; i < 10 ; i++ {
    fmt.Println(<-ch)
  }

  fmt.Println(Same(tree.New(1), tree.New(1)))
  fmt.Println(Same(tree.New(1), tree.New(2)))
}

