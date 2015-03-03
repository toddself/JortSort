[![build status](https://secure.travis-ci.org/toddself/JortSort.png)](http://travis-ci.org/toddself/JortSort)

# JortSort

This is an implementation of Jenn Schiffer's famous [JortSort](http://jort.technology) in Go. This will tell you if you need to sort your array or not.

## Installing

`go get github.com/toddself/JortSort`

## Usage

```
import (
  "github.com/toddself/JortSort"
  "fmt"
)

func Main() {
  arr := []string{"a", "b", "c", "d"}
  jortArray = jort.Strings2Sortable(arr)
  sorted := jort.JortSort(jortArray)

  if sorted {
    fmt.Println("Your array is sorted")
  } else {
    fmt.Println("Your array is not sorted")
  }
}
```

Since Go is a pain in the ass and "strongly-typed" (whatever that means), I had to implement a generic interface and make that interface sortable. So you need to convert all your arrays into these special sortable types. There are three helper functions for this: `String2Sortable`, `Int2Sortable` and `Float642Sortable`.  If you need any other number type, you should re-think your project's data probably.

## Tests

Everything has tests rights?

```
↳ go test -v
=== RUN TestStringSort
--- PASS: TestStringSort (0.00s)
=== RUN TestIntSort
--- PASS: TestIntSort (0.00s)
=== RUN TestStringToSortable
--- PASS: TestStringToSortable (0.00s)
=== RUN TestIntToSortable
--- PASS: TestIntToSortable (0.00s)
=== RUN TestFloat64ToSortable
--- PASS: TestFloat64ToSortable (0.00s)
PASS
ok    github.com/toddself/JortSort  0.009s
```

## License
MIT