package main
import "fmt"
import "os"
import "strconv"

type Integer interface {
  ~int | ~int8 | ~int16 | ~int32 | ~int64 |
  ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func main() {
	f := MakeFactorial[int]()
	skipUndefined := Catch(NoDefinedValue("factorial"))
	Each(os.Args[1:], func(v string) {
		skipUndefined(PrintResult(v, f))
	})
}

func MakeFactorial[T Integer]() (f func(T) T) {
	c := map[T] T { 0: 1 }
	return func(n T) (r T) {
		if n < 0 {
			panic(n)
		}
		if r = c[n]; r == 0 {
			r = n * f(n - 1)
		}
		c[n] = r
		return
	}
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

func PrintResult[T Integer](v string, f func(T) T) func() {
	return func() {
		if x, e := strconv.Atoi(v); e == nil {
			fmt.Printf("f(%v) = %v\n", x, f(T(x)))
		} else {
			panic(v)
		}
	}
}
