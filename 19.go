package main
import "fmt"
import "os"
import "strconv"

type Integer interface {
  ~int | ~int8 | ~int16 | ~int32 | ~int64 |
  ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func main() {
	for _, v := range os.Args[1:] {
		func() {
			defer func() {
				if x := recover(); x != nil {
					fmt.Printf("no factorial defined for %v\n", x)
				}
			}()
			if x, e := strconv.Atoi(v); e == nil && x > -1 {
				fmt.Printf("%v! = %v\n", x, Factorial(x))
			} else {
				panic(v)
			}
		}()
	}
}

func Factorial[T Integer](n T) T {
	if n == 0 {
		return 1
	}
	return n * Factorial(n - 1)
}
