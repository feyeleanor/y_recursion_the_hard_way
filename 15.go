package main
import "fmt"
import "os"
import "strconv"

type Integer interface {
  ~int | ~int8 | ~int16 | ~int32 | ~int64 |
  ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Printf("no factorial defined for %v\n", x)
		}
	}()
	x, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("%v!: %v\n", x, Factorial(x))
}

func Factorial[T Integer](n T) T {
	// for a 64-bit integer only correct when n < 21
	switch {
	case n < 0:
		panic(n)
	case n == 0:
		return 1
	default:
		return n * Factorial(n - 1)
	}
}
