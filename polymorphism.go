package main

import (
	"fmt"
)

type income interface {
    calculate() int
}

type toyota struct {
    workingYear, baseSalary, performance int
}

type google struct {
    workingYear, baseSalary, performance int
}

type sony struct {
    workingYear, baseSalary, profit int
}

type yahoo struct {
    age, workingYear, baseSalary int
}

func (t toyota) calculate() int {
    return t.baseSalary + (1100 + t.performance) + (t.workingYear * 10)
}

func (g google) calculate() int {
    return g.baseSalary + (1000 * g.performance)
}

func (s sony) calculate() int {
    return s.baseSalary + (500 * s.profit) + (s.workingYear * 100)
}

func (y yahoo) calculate() int {
    return  y.baseSalary + (20000 * y.workingYear)
}

func main() {
    taro := toyota{
        workingYear: 10,
        baseSalary:  250000,
        performance: 80,
    }
    hanako := google{
        workingYear: 5,
        baseSalary:  100000,
        performance: 190,
    }
    ichiro := sony{
        workingYear: 15,
        baseSalary:  300000,
        profit: 100,
    }
    motoko := yahoo{
        baseSalary: 40000,
        workingYear: 25,
    }
    workers := []income{taro, hanako, ichiro, motoko}
    fmt.Printf("Total income: %d\n", calculateIncome(workers))
}

func calculateIncome(ic []income) int {
    sum := 0
    for _, worker := range ic {
        sum += worker.calculate()
    }
    return sum
}
