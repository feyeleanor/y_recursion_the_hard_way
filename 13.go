package main
import "os"
import "strconv"

var limit int

func init() {
  if x, e := strconv.Atoi(os.Args[1]); e == nil {
    limit = x
  } else {
    os.Exit(1)
  }
}

# limit 67108726 on a 64GB Intel MacBook Pro with Go version 1.24.3 darwin/amd64

func main() {
  limit--
  if limit > 0 {
    main()
  }
}
