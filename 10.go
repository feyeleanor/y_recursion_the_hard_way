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

type Intish interface {
  Int() int
}

type Accumulator[T Scalar] func(T) T

func main() {
  var n []int
  for _, v := range os.Args[1:] {
    x, _ := strconv.Atoi(v)
    n = append(n, x)
  }
  os.Exit(MakeAccumulator(n...).Int())
}

func MakeAccumulator[T Scalar](s ...T) (a Accumulator[T]) {
  var sum T
  a = func(x T) T {
    sum += x
    return sum
  }
  for _, v := range s {
    a.Add(v)
  }
  return
}

func (a Accumulator[T]) Add(x any) Accumulator[T] {
  switch x := x.(type) {
  case T:
    a(x)
  case Intish:
    a(T(x.Int()))
  }
  return a
}

func (a Accumulator[T]) Int() int {
  return int(a(0))
}
