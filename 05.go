package main
import "os"
import "strconv"

var sum int

func main() {
  for _, v := range os.Args[1:] {
    x, _ := strconv.Atoi(v)
    accumulate(x)
  }
  os.Exit(accumulate(0))
}

func accumulate(x int) int {
  sum += x
  return sum
}
