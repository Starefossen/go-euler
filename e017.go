package main

import (
  "fmt"
)

var base = []string{"zero", "one", "two", "three", "four", "five", "six",
"seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen",
"fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}

var base10 = []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy",
"eighty", "ninety"}

func main() {
  var a, b, c, d, sum int

  // a = 1 to 10
  for i := 1; i < 10; i++ {
    a += len(base[i])
  }

  // b = 10 to 19
  for i := 10; i < 20; i++ {
    b += len(base[i])
  }

  // c = 20 to 99
  for i := 2; i < 10; i++ {
    c += (len(base10[i]) * 10) + a + 9
  }

  // d = 1 to 99
  d = a + b + c

  sum = d

  // 1 to 999
  for i := 1; i < 10; i++ {
    sum += ((len(base[i]) + len("hundred")) * 100) + ((len("and") + d) * 99)
  }

  sum += len(base[1]) + len("thousand")

  fmt.Println(sum)
}
