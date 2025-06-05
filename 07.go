package main
import "os"
import "strconv"

type Accumulator func(int) int

func main() {
  a := MakeAccumulator()
  for _, v := range os.Args[1:] {
    x, _ := strconv.Atoi(v)
    a(x)
  }
  os.Exit(a(0))
}

func MakeAccumulator() Accumulator {
  var sum int
  return func(x int) int {
    sum += x
    return sum
  }
}
