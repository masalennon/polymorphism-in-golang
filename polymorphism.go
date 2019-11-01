package main

import (
	"fmt"
)

type income interface {
    calculate() int
}

type toyota struct {
    age, workingYear, baseSalary int
    performance                  int
}

type google struct {
    age, workingYear, baseSalary int
    performance                  int
}

type sony struct {
    age, workingYear, baseSalary int
    performance                  int
}

func (t toyota) calculate() int {
    return t.baseSalary + (1100 + t.performance) + (t.workingYear * 10)
}

func (g google) calculate() int {
    return g.baseSalary + (1000 * g.performance)
}

func (s sony) calculate() int {
    return s.baseSalary + (10 * s.performance) + (s.workingYear * 100)
}

func main() { //ポリモーフィズムを使った例
    taro := toyota{
        age:         33,
        workingYear: 10,
        baseSalary:  250000,
        performance: 80,
    }
    hanako := google{
        age:         28,
        workingYear: 5,
        baseSalary:  100000,
        performance: 190,
    }
    ichiro := sony{
        age:         40,
        workingYear: 15,
        baseSalary:  300000,
        performance: 130,
    }
    workers := []income{taro, hanako, ichiro}
    fmt.Printf("Total income: %d\n", calculateIncome(workers))
}

func calculateIncome(ic []income) int {
    sum := 0
    for _, worker := range ic {
        sum += worker.calculate()
    }
    return sum
}
