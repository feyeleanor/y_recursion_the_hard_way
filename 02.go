package main
import "os"

type Integer interface {
  ~int | ~int8 | ~int16 | ~int32 | ~int64 |
  ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Scalar interface {
  Integer | ~float32 | ~float64
}

func main() {
  os.Exit(add(3, 4))
}

func add[T Scalar](x, y T) T {
  return x + y
}
