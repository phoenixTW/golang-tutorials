package main
import "fmt"

func main()  {
  action := func () (int) {
    return 0
  }
  fmt.Printf("%d\n", action())
}
