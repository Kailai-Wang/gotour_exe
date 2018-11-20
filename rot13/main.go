package main

import (
  "io"
  "os"
  "strings"
)

type rot13Reader struct {
  r io.Reader
}

func Convert(in byte) byte {
  if (in >= 'A' && in <= 'M') || (in >= 'a' && in <= 'm') {
    return in + 13
  } else if (in >= 'N' && in <= 'Z') || (in >= 'n' && in <= 'z') {
    return in - 13
  } else {
    return in
  }
}

func (rr rot13Reader) Read(b []byte) (int, error) {
  len, err := rr.r.Read(b)
  if err != nil {
    return len, err
  }
  
  for i,v := range b {
    b[i] = Convert(v)
  }
  
  return len,err
}

func main() {
  s := strings.NewReader("Lbh penpxrq gur pbqr!")
  r := rot13Reader{s}
  io.Copy(os.Stdout, &r)
}

