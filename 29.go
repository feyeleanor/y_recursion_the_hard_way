package main
import "fmt"
import "os"
import "strconv"

type Integer interface {
  ~int | ~int8 | ~int16 | ~int32 | ~int64 |
  ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Function[T, R any] func(T) R
type Transformer[T, R any] func(Function[T, R]) Function[T, R]

type Recursor[T, R any] func(Recursor[T, R]) Function[T, R]

func (r Recursor[T, R]) Apply(t Transformer[T, R]) Function[T, R] {
	return t(r(r))
}

func main() {
	skipUndefined := Catch(NoDefinedValue("factorial"))
	factorial := Y(func(h Function[int, int]) Function[int, int] {
		return func(n int) (r int) {
			switch {
			case n < 0:
				panic(n)
			case n > 1:
				return n * h(n-1)
			}
			return 1
		}
	})

	Each(os.Args[1:], func(v string) {
		skipUndefined(PrintResult(v, factorial))
	})
}

// Y = ->f.(->x.f(x x))(->x.f(x x))

func Y[T, R any](t Transformer[T, R]) Function[T, R] {
	g := func(r Recursor[T, R]) Function[T, R] {
		return func(x T) R {
			return r.Apply(t)(x)
		}
	}
	return g(g)
}

func Each[T any](s []T, f func(T)) {
	if len(s) > 0 {
		f(s[0])
		Each(s[1:], f)
	}
}

func Catch(e func()) func(func()) {
	return func(f func()) {
		defer e()
		f()
	}
}

func NoDefinedValue(s string) func() {
	return func() {
		if x := recover(); x != nil {
			fmt.Printf("no %v defined for %v\n", s, x)
		}
	}
}

func PrintResult[T, R Integer](v string, f Function[T, R]) func() {
	return func() {
		if x, e := strconv.Atoi(v); e == nil {
			fmt.Printf("f(%v) = %v\n", x, f(T(x)))
		} else {
			panic(v)
		}
	}
}
