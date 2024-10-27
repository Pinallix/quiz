package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Welcome to the Quiz game.")
	questions := getFileData()
	timer, challangeLevel := setQuizParamenters(len(questions))
	// go setTimmer(timer, c)

	results := startGame(questions, timer, challangeLevel)

	fmt.Println("you got ", results, "answers correct!")
}

// func setTimmer(sleepTime time.Duration, c chan bool) {
// 	time.Sleep(sleepTime)
// 	fmt.Println("Time's up!")
// 	c <- true
// }

func startGame(questions map[string]int, time time.Duration, challangeLevel int) int {
	score := 0
	counter := 1
	g := 0
	fmt.Println("challange set to: ", time, "s and ", challangeLevel, "questions.")
	reader := bufio.NewReader(os.Stdin)
	for question, answer := range questions {
		fmt.Println(question, ":")
		guess := parseInput(reader)
		g, _ = strconv.Atoi(guess)
		if g == answer {
			score++
		}
		if counter >= challangeLevel {
			break
		}
		counter++
	}
	return score
}

func parseInput(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("invalid command")
	}
	input = strings.TrimSpace(input)
	return input
}

func getFileData() map[string]int {
	reader := bufio.NewReader(os.Stdin)
	fileFound := false
	data := []byte{}
	var err error
	for !fileFound {
		fmt.Println("Please specify the file name you wish to load or press enter for default problems.csv:")

		fileName := parseInput(reader)
		if len(fileName) == 0 {
			fileName = "problems.csv"
		}
		data, err = openFile(fileName)
		if err != nil {
			fmt.Println("File not found")
		}
		if len(data) != 0 {
			fileFound = true
		}
	}
	return parseData(data)
}

func openFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)

}

func parseData(data []byte) map[string]int {
	questions := make(map[string]int)
	reader := csv.NewReader(strings.NewReader(string(data)))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("error parsing data on file")
			os.Exit(1)
		}
		result, err := strconv.Atoi(record[1])
		if err != nil {
			fmt.Println("error parsing answer value")
			os.Exit(1)
		}
		questions[record[0]] = result

	}
	fmt.Println("loaded ", len(questions), "problems")
	return questions
}

func setQuizParamenters(numberOfQuestions int) (time.Duration, int) {
	reader := bufio.NewReader(os.Stdin)
	parametersSet := false
	var timeInput int
	challange := 0
	var err error
	for !parametersSet {
		fmt.Println("Set Quiz challange time (in seconds) or press enter for default time(30s):")
		inputVal := parseInput(reader)
		if len(inputVal) == 0 {
			timeInput = 30
			fmt.Println("Default time set to 30s")
		} else {
			timeInput, err = strconv.Atoi(inputVal)
		}
		if err != nil {
			fmt.Println("invalid input of time")
		}
		if timeInput > 0 {
			parametersSet = true
		}
	}
	timer := time.Duration(timeInput) * time.Second
	parametersSet = false
	for !parametersSet {
		fmt.Println("Choose Quiz challange from 1 to ", numberOfQuestions, " or press enter to choose all:")
		inputVal := parseInput(reader)
		if len(inputVal) == 0 {
			challange = numberOfQuestions
			fmt.Println("Default challange set to ", numberOfQuestions)
		} else {
			challange, _ = strconv.Atoi(inputVal)
		}
		if challange > 0 && challange <= numberOfQuestions {
			parametersSet = true
		} else {
			fmt.Println("invalid input of questions")
		}
	}
	return timer, challange
}
