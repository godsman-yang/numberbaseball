package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	// 결정사항 - 숫자 세자리? 3자리수 숫자, 숫자 3개, 문자열

	// 임의 숫자 세자리 만들기
	// var computerNumber int
	// computerNumber = 648

	computerNumber := MakeComputerNumber()

	count := 0
	for {
		count++
		// 입력 세자리
		fmt.Printf("Input your number %2d: ", count)
		userNumber := InputUserNumber()

		if !CheckNumber(userNumber) {
			count--
			continue
		}

		// 숫자 비교
		result := CompareNumber(computerNumber, userNumber)

		if result {
			fmt.Println("3 strike(s) - congratulation")
			break
		}
	}
}

func MakeComputerNumber() int {
	result := 648

	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

	if a[0] != 0 {
		result = a[0] * 100
	} else {
		result = a[3] * 100
	}
	result += (a[1] * 10)
	result += a[2]
	// result = 101 + r1.Intn(900)
	// fmt.Println("Random: ", result)
	return result
}

func InputUserNumber() int {
	var number int
	_, _ = fmt.Scanf("%d", &number)

	return number
}

func CheckNumber(number int) bool {

	// 3자리가 아니면
	if number <= 100 || number >= 1000 {
		return false
	}

	// 중복되는 숫자가 있으면
	nums := strconv.Itoa(number)
	first := nums[0:1]
	second := nums[1:2]
	third := nums[2:3]

	if strings.Count(nums, first) > 1 {
		return false
	} else if strings.Count(nums, second) > 1 {
		return false
	} else if strings.Count(nums, third) > 1 {
		return false
	}

	return true
}

func CompareNumber(computerNumber, userNumber int) bool {
	strikes := 0
	balls := 0

	if computerNumber == userNumber {
		strikes = 3
		balls = 0
		PrintBallCount(strikes, balls)
		return true
	}

	cNumString := strconv.Itoa(computerNumber)
	first := cNumString[0:1]
	second := cNumString[1:2]
	third := cNumString[2:3]
	u_nums := strconv.Itoa(userNumber)

	// fmt.Println("fist-second-third: ", first, second, third)
	// fmt.Printf("cNumString: %c, %c, %c\n", cNumString[0], cNumString[1], cNumString[2])

	result := strings.Index(u_nums, first)
	if result < 0 {
		// fmt.Println("first-none: ", u_nums, result)
	} else if result == 0 {
		strikes++
		// fmt.Println("first-strikes: ", u_nums, result)
	} else {
		balls++
		// fmt.Println("first-balls: ", u_nums, result)
	}
	result = strings.Index(u_nums, second)
	if result < 0 {
		// fmt.Println("second-none: ", u_nums, result)
	} else if result == 1 {
		strikes++
		// fmt.Println("second-strikes: ", u_nums, result)
	} else {
		balls++
		// fmt.Println("second-balls: ", u_nums, result)
	}
	result = strings.Index(u_nums, third)
	if result < 0 {
		// fmt.Println("third-none: ", u_nums, result)
	} else if result == 2 {
		strikes++
		// fmt.Println("third-strikes: ", u_nums, result)
	} else {
		balls++
		// fmt.Println("third-balls: ", u_nums, result)
	}

	PrintBallCount(strikes, balls)
	return false
}

func PrintBallCount(strikes, balls int) {
	msg := "continue"
	if strikes == 3 && balls == 0 {
		msg = "complete"
	}
	fmt.Printf("%s-> %d strikes, %d balls - %s\n", strings.Repeat(" ", 30), strikes, balls, msg)
}
