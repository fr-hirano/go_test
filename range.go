package main

import (
    "fmt"
)

func main() {
    arr := [4]string{"a","b","c","d"}

    for i, s :=range arr {
      fmt.Println(i, s)
    }

}
