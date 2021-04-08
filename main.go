package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

func Task1(text string) {
	fmt.Println("Task 1")
	reg := regexp.MustCompile(`\b[eyuioaEYUIOA]+[a-zA-Z]*\b`)

	matchesIndex := reg.FindAllStringIndex(text, -1)
	matchWords := reg.FindAllString(text, -1)

	fmt.Println("Positions: ")
	for i, match := range matchesIndex {
		fmt.Println(matchWords[i], ":", match[0])
	}
}

func Task2(text string) {
	fmt.Println("Task 2")
	reg := regexp.MustCompile(`\b[xyz]\w{5,17}[xyz]\b`)
	matchWords := reg.FindAllString(text, -1)

	fmt.Println("Words:", matchWords)
}

func Task3(text string) {
	var count int = 0
	fmt.Println("Task 3")
	reg := regexp.MustCompile(`\d+`)
	matchWords := reg.FindAllString(text, -1)
	for _, word := range matchWords {
		num, e := strconv.Atoi(word)
		if e != nil {
			fmt.Println(e)
			os.Exit(-1)
		}
		if num%10 == 0 {
			num /= 10
			strNum := strconv.Itoa(num)
			numReg := regexp.MustCompile(word)
			text = numReg.ReplaceAllString(text, strNum)
			count++
		}
	}
	fmt.Print("Replace count: ", count, "\n", text)
}

func main() {
	bufFileData, e := ioutil.ReadFile("f.txt")
	refString := string(bufFileData)
	if e != nil {
		fmt.Println("ouf, something happend: ", e)
	}
	Task1(refString)
	Task2(refString)
	Task3(refString)
	fmt.Println("\nPress Enter to exit")
	fmt.Scanf("a")
}
