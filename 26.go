package main
import "fmt"
import "os"
import "strconv"

type Integer interface {
  ~int | ~int8 | ~int16 | ~int32 | ~int64 |
  ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Function func(any) any

func main() {
	skipUndefined := Catch(NoDefinedValue("factorial"))
	factorial := Y(func(h any) any {
		return func(n any) (r any) {
			if n, ok := n.(int); ok {
				switch {
				case n == 0, n == 1:
					return 1
				case n > 1:
					return n * h.(func(any) any)(n-1).(int)
				}
			}
			panic(n)
		}
	})

	Each(os.Args[1:], func(v string) {
		skipUndefined(PrintResult(v, factorial))
	})
}

// Y = ->f.(->x.f(x x))(->x.f(x x))

func Recursor(f any) Function {
	return f.(func(any) any)(f).(func(any) any)
}

func Y(g Function) Function {
	return Recursor(func(f any) any {
		return g(func(x any) any {
			return Recursor(f)(x)
		})
	})
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

func PrintResult(v string, f Function) func() {
	return func() {
		if x, e := strconv.Atoi(v); e == nil {
			fmt.Printf("f(%v) = %v\n", x, f(x))
		} else {
			panic(v)
		}
	}
}
