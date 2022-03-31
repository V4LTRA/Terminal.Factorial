package main

import (
  "fmt"
  "strconv"
)

const (
  End       = "Конец"
  Factorial = "Факториал"
)

func main(){workWithFactorial}
func One() {
  for {
    var msg string
    msg = ""
    fmt.Println("Введите команду:")
    fmt.Scanln(&msg)
    fmt.Println(msg)
    if msg == End {
      break
    }
    if msg == Factorial {
      if err := workWithFactorial(); err != nil {
        break
      }
    }
  }
}

func workWithFactorial() error {
  for {
    var msg string
    msg = ""
    fmt.Println("Введите команду окончания или число:")
    fmt.Scanln(&msg)
    if msg == End {
      break
    }
    num, err := strconv.Atoi(msg)
    if err != nil {
      fmt.Println("введена команда")
      continue
    }
    fmt.Println("Факториал вашего числа равен =", factorial(uint64(num)))
  }
  return nil
}

func factorial(n uint64) uint64 {
  if n == 0 {
    return 1
  }
  return n * factorial(n-1)
}