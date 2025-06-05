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

type Accumulator[T Scalar] func(T) T

func main() {
  a := MakeAccumulator[int]()
  for _, v := range os.Args[1:] {
    x, _ := strconv.Atoi(v)
    a(x)
  }
  os.Exit(a.Int())
}

func MakeAccumulator[T Scalar]() Accumulator[T] {
  var sum T
  return func(x T) T {
    sum += x
    return sum
  }
}

func (a Accumulator[T]) Int() int {
  return int(a(0))
}
