package main
import "os"
import "strconv"

type Integer interface {
  ~int | ~int8 | ~int16 | ~int32 | ~int64 |
  ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Scalar interface {
  Integer | ~float32 | ~float64
}

func main() {
  var sum int
  for _, v := range os.Args[1:] {
    x, _ := strconv.Atoi(v)
    sum = add(sum, x)
  }
  os.Exit(sum)
}

func add[T Scalar](x, y T) T {
  return x + y
}
