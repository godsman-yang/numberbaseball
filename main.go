package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type scoreT struct {
	index          int
	name           string
	startTime      time.Time
	endTime        time.Time
	computerNumber int
	tryNumber      int
}

func main() {
	// 결정사항 - 숫자 세자리? 3자리수 숫자, 숫자 3개, 문자열

	// 임의 숫자 세자리 만들기
	// var computerNumber int
	// computerNumber = 648

	userName := InputUserName()
	fmt.Println("User Name: ", userName)

	computerNumber := MakeComputerNumber()
	startTime := time.Now()

	count := 0
	for {
		count++
		// 입력 세자리
		// fmt.Printf("Input your number [%s] %2d: ", userName, count)
		userNumber := InputUserNumber(userName, count)

		// 숫자 비교
		result := ComputeBallCount(computerNumber, userNumber)

		if result {
			fmt.Println("3 strike(s) - congratulation")
			break
		}
	}
	endTime := time.Now()
	// fmt.Printf("1, shyang, 2020-09-06 11:08:01, 10.1, 158, 8\n")

	loc, _ := time.LoadLocation("Asia/Seoul")
	startTimeLocal := startTime.In(loc)
	endTimeLocal := endTime.In(loc)

	score := scoreT{0, userName, startTimeLocal, endTimeLocal, computerNumber, count}
	WriteScore(score)

	fmt.Printf("1, %s, %s, %s, %d, %d\n", userName, startTimeLocal, endTimeLocal, computerNumber, count)
}

func InputUserName() string {
	var name string

	fmt.Printf("Input your name: ")
	_, _ = fmt.Scanf("%s", &name)

	return name

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
	// result = 123
	return result
}

func InputUserNumber(userName string, count int) int {
	var number int

	for {
		fmt.Printf("Input your number [%s] %2d: ", userName, count)
		_, _ = fmt.Scanf("%d", &number)

		if IsValidNumber(number) {
			break
		}
	}
	return number
}

func IsValidNumber(number int) bool {

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

func ComputeBallCount(computerNumber, userNumber int) bool {
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
	uNumString := strconv.Itoa(userNumber)

	// fmt.Println("fist-second-third: ", first, second, third)
	// fmt.Printf("cNumString: %c, %c, %c\n", cNumString[0], cNumString[1], cNumString[2])

	result := strings.Index(uNumString, first)
	if result < 0 {
		// fmt.Println("first-none: ", uNumString, result)
	} else if result == 0 {
		strikes++
		// fmt.Println("first-strikes: ", uNumString, result)
	} else {
		balls++
		// fmt.Println("first-balls: ", uNumString, result)
	}
	result = strings.Index(uNumString, second)
	if result < 0 {
		// fmt.Println("second-none: ", uNumString, result)
	} else if result == 1 {
		strikes++
		// fmt.Println("second-strikes: ", uNumString, result)
	} else {
		balls++
		// fmt.Println("second-balls: ", uNumString, result)
	}
	result = strings.Index(uNumString, third)
	if result < 0 {
		// fmt.Println("third-none: ", uNumString, result)
	} else if result == 2 {
		strikes++
		// fmt.Println("third-strikes: ", uNumString, result)
	} else {
		balls++
		// fmt.Println("third-balls: ", uNumString, result)
	}

	PrintBallCount(strikes, balls)
	return false
}

func PrintBallCount(strikes, balls int) {
	msg := "continue"
	if strikes == 3 && balls == 0 {
		msg = "complete"
	}
	fmt.Printf("%s-> %d strike(s), %d ball(s) - %s\n", strings.Repeat(" ", 34), strikes, balls, msg)
}

func ReadIndex() int {
	index := 0

	file, err := os.Open("score.txt")
	if err != nil {
		return 0
	}
	defer file.Close()

	// csv reader 생성
	rdr := csv.NewReader(bufio.NewReader(file))

	// csv 내용 모두 읽기
	rows, _ := rdr.ReadAll()

	if len(rows) <= 0 {
		return 0
	}

	index, error := strconv.Atoi(rows[len(rows)-1][0])
	if error != nil {
		return 0
	}
	// fmt.Println("rows: ", len(rows), index)

	// 행,열 읽기
	for i, row := range rows {
		for j := range row {
			fmt.Printf("%s ", rows[i][j])
		}
		fmt.Println()
	}
	return index
}

/*
func ReadIndex() int {
	index := 0

	file, err := os.Open("score.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
		return index
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	for _, eachline := range txtlines {
		fmt.Println(eachline)
	}
	return index
}
*/

func WriteScore(score scoreT) {
	// 파일 생성
	file, err := os.OpenFile("score.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
	}

	defer file.Close()
	score.index = ReadIndex() + 1
	diff := score.endTime.Sub(score.startTime).Seconds()
	// fmt.Fprintf(file, "%d, %s, %s, %s, %d, %d\n", score.index, score.name, score.startTime, score.endTime, diff, score.computerNumber, score.tryNumber)
	fmt.Fprintf(file, "%d, %s, %s, %.1f, %d, %d\n", score.index, score.name, score.startTime.Format("2006-01-02 15:04:05"), diff, score.computerNumber, score.tryNumber)
}
