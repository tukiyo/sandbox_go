package main

import (
  "fmt"
  "errors"
)

func main() {
  if err := genError(); err != nil {
    fmt.Println(err)
  }
}

func genError() error {
  return errors.New("ERROR")
}
