package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func countResults(inputName, outputName string) {
	input, err := os.Open(inputName)
	if err != nil {
		fmt.Println("Ошибка при открытии файла: ", err)
		return
	}
	defer input.Close()

	output, err := os.Create(outputName)
	if err != nil {
		fmt.Println("Ошибка при создании файла: ", err)
		return
	}
	defer output.Close()

	scanner := bufio.NewScanner(input)
	writer := bufio.NewWriter(output)

	re := regexp.MustCompile(`([0-9]+)([+-]+)([0-9]+)`)

	for scanner.Scan() {
		s := scanner.Text()
		matches := re.FindStringSubmatch(s)

		if matches != nil {
			num1, _ := strconv.Atoi(matches[1])
			operator := matches[2]
			num2, _ := strconv.Atoi(matches[3])

			var res int

			if operator == "+" {
				res = num1 + num2
			} else if operator == "-" {
				res = num1 - num2
			}

			resStr := matches[0] + "=" + strconv.Itoa(res)

			writer.Write([]byte(resStr + "\n"))
		}
	}
	writer.Flush()
}

func main() {
	countResults("./input.txt", "./output.txt")
}
