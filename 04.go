package main
import "os"
import "strconv"

var sum int

func main() {
  for i, v := range os.Args[1:] {
    x, _ := strconv.Atoi(v)
    if i / 2 == 0 {
      sum += x
    } else {
      accumulate(x)
    }
  }
  os.Exit(sum)
}

func accumulate(x int) {
  sum += x
}
