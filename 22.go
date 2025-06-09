package main
import "fmt"
import "os"
import "strconv"

type Integer interface {
  ~int | ~int8 | ~int16 | ~int32 | ~int64 |
  ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func main() {
	skipUndefined := Catch(NoDefinedValue("factorial"))
	for _, v := range os.Args[1:] {
		skipUndefined(PrintFactorial(v))
	}
}

func Factorial[T Integer](n T) T {
	if n == 0 {
		return 1
	}
	return n * Factorial(n - 1)
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

func PrintFactorial(v string) func() {
	return func() {
		if x, e := strconv.Atoi(v); e == nil && x > -1 {
			fmt.Printf("%v! = %v\n", x, Factorial(x))
		} else {
			panic(v)
		}
	}
}
