package day25

import (
	"aoc22/utils"
	"fmt"
)

type Result struct {
	part1 string
	part2 string
}

func Pow5(p int) int {
	res := 1
	for p > 0 {
		res *= 5
		p--
	}
	return res
}

func ParseSnafu(num string) int {
	result := 0
	for i := len(num) - 1; i >= 0; i-- {
		symbol := num[i]
		power := len(num) - 1 - i
		digit := 0
		switch symbol {
		case '2':
			digit = 2
		case '1':
			digit = 1
		case '0':
			digit = 0
		case '-':
			digit = -1
		case '=':
			digit = -2
		default:
			panic(fmt.Sprintf("Unexpected symbol: %s", string(symbol)))
		}

		result += digit * Pow5(power)
	}

	return result
}

func DivMod(a, b int) (int, int) {
	return int(a / b), a % b
}

func Base5(num int) []int {
	result := []int{}

	p := 0
	for num != 0 {
		q, r := DivMod(num, 5)
		result = append([]int{r}, result...)
		p++
		num = q
	}

	return result
}

func IsValidSnafu(num []int) bool {
	for _, n := range num {
		if n < -2 || n > 2 {
			return false
		}
	}

	return true
}

func MakeValidIterate(num []int) []int {
	for i := len(num) - 1; i > 0; i-- {
		d := num[i]
		if d > 2 {
			// D*5^P = 5^(P+1) + (D-5)*5^P
			num[i-1] += 1
			num[i] = d - 5
		} else if d < -2 {
			// -D*5^P = -5^(P+1) + (5-D)*5^P
			num[i-1] -= 1
			num[i] = 5 + d
		}
	}

	if num[0] > 2 {
		num = append([]int{1}, num...)
		num[1] -= 5
	} else if num[0] < -2 {
		num = append([]int{-1}, num...)
		num[1] += 5
	}

	return num
}

func ArrToSnafu(num []int) string {
	result := make([]byte, len(num))
	for i, d := range num {
		if d >= 0 {
			result[i] = byte(d) + '0'
		} else {
			switch d {
			case -1:
				result[i] = '-'
			case -2:
				result[i] = '='
			}
		}
	}
	return string(result)
}

func ParseDecimal(num int) string {
	num5 := Base5(num)

	for !IsValidSnafu(num5) {
		num5 = MakeValidIterate(num5)
	}

	return ArrToSnafu(num5)
}

func Run(fileName string) Result {
	lines := utils.ReadFileToLines(fileName)

	sum := 0
	for _, line := range lines {
		sum += ParseSnafu(line)
	}
	snafuSum := ParseDecimal(sum)

	return Result{snafuSum, ""}
}
