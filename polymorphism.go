package main

import (
	"fmt"
)

type worker struct {
    age, workingYear, baseSalary int
    performance                  int
    company                      string
}

func main() { //ポリモーフィズムアンチパターン例
    taro := worker{
        age:         33,
        workingYear: 10,
        baseSalary:  250000,
        performance: 80,
        company:     "toyota",
    }
    hanako := worker{
        age:         28,
        workingYear: 5,
        baseSalary:  100000,
        performance: 190,
        company:     "google",
    }
    ichiro := worker{
        age:         40,
        workingYear: 15,
        baseSalary:  300000,
        performance: 130,
        company:     "sony",
    }
    workers := []worker{taro, hanako, ichiro}
    fmt.Printf("Total income: %d\n", calculateIncome(workers))
}

func calculateIncome(workers []worker) int {
    sum := 0
    for _, worker := range workers {
        switch worker.company {
        case "toyota":
            sum += worker.baseSalary + (1100 + worker.performance) + (worker.workingYear * 10)
        case "google":
            sum += worker.baseSalary + (1000 * worker.performance)
        case "sony":
            sum += worker.baseSalary + (10 * worker.performance) + (worker.workingYear * 100)
        default:
            sum += 0
        }
    }
    return sum
}