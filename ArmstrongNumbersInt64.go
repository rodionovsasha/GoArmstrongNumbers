package main

import (
    "math"
    "fmt"
    "time"
    "sort"
)

const NUMBER_OF_SIMPLE_DIGITS = 10
const MAX_NUMBER = math.MaxInt64

var arrayOfPowers [NUMBER_OF_SIMPLE_DIGITS][20]int64

type uint64arr []int64

func (a uint64arr) Len() int {
    return len(a)
}
func (a uint64arr) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}
func (a uint64arr) Less(i, j int) bool {
    return a[i] < a[j]
}

func main() {
    startTime := time.Now()

    for i := 1; i < NUMBER_OF_SIMPLE_DIGITS; i++ {
        for j := 1; j <= getNumberOfDigits(MAX_NUMBER); j++ {
            arrayOfPowers[i][j] = pow(i, j)
        }
    }

    mapOfNumbers := make(map[int64]bool)
    var armstrongNumbers []int64

    //Main loop
    for i := int64(1); i < MAX_NUMBER; i = getNextNumber(i) {
        if i < 0 {
            break // the maximum value is reached
        }

        var sumOfPowers = getSumOfPowers(int64(i))
        if sumOfPowers <= MAX_NUMBER && isArmstrongNumber(sumOfPowers) {

            if _, ok := mapOfNumbers[int64(sumOfPowers)]; !ok {
                mapOfNumbers[int64(sumOfPowers)] = true
                armstrongNumbers = append(armstrongNumbers, sumOfPowers)
            }
        }
    }

    arrayOfNumbers := uint64arr(armstrongNumbers)
    sort.Sort(arrayOfNumbers)
    for k := range arrayOfNumbers {
        fmt.Printf("%d. %d\n", k+1, arrayOfNumbers[k])
    }

    elapsedTime := time.Since(startTime)
    fmt.Printf("\nExecution time: %s\n", elapsedTime)
}

func getNextNumber(number int64) int64 {
    var copyOfNumber = number
    if isGrowingNumber(copyOfNumber) { // here we have numbers where each digit not less than previous one and not more than next one: 12, 1557, 333 and so on.
        return copyOfNumber + 1
    }

    // here we have numbers which end in zero: 10, 20, ..., 100, 110, 5000, 1000000 and so on.
    var lastNumber = 1 //can be: 1,2,3..., 10,20,30,...,100,200,300,...

    for copyOfNumber%10 == 0 { // 5000 -> 500 -> 50: try to get the last non-zero digit
        copyOfNumber = copyOfNumber / 10
        lastNumber = lastNumber * 10
    }
    var lastNonZeroDigit = copyOfNumber % 10

    return number + (lastNonZeroDigit * int64(lastNumber) / 10) //e.g. number=100, lastNumber=10, lastNonZeroDigit=1
}

func pow(base int, exponent int) int64 {
    var pow = 1
    for i := 1; i <= exponent; i++ {
        pow *= base
    }
    return int64(pow)
}

/*
* 135 returns true:  1 < 3 < 5
* 153 returns false: 1 < 5 > 3
* */
func isGrowingNumber(number int64) bool {
    return (number+1)%10 != 1
}

func isArmstrongNumber(number int64) bool {
    return number == getSumOfPowers(number)
}

func getSumOfPowers(number int64) int64 {
    var currentNumber = number
    var power = getNumberOfDigits(currentNumber)
    var currentSum int64 = 0

    for currentNumber > 0 {
        currentSum = currentSum + arrayOfPowers[(int64)(currentNumber%10)][power] // get powers from array by indexes and then the sum.
        currentNumber /= 10
    }
    return int64(currentSum)
}

func getNumberOfDigits(number int64) int {
    return int(math.Log10(float64(number)) + 1)
}
