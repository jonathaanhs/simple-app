package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func ReverseLowerAndUpperCaseChar(input string) {
	result := ""

	for _, v := range input {
		if unicode.IsLetter(v) {
			if v >= 'A' && v <= 'Z' {
				v = unicode.ToLower(v)
			} else {
				v = unicode.ToUpper(v)
			}
		}

		result = result + string(v)
	}

	fmt.Println("=== Reverse Lower and Upper case character ===")
	fmt.Println("Input : ", input)
	fmt.Println("Result : ", result)
	fmt.Println()

	return
}

func ConvertTotalCharInString(input string) {
	var result string
	tmp := map[string]bool{}

	for _, v := range input {
		if !tmp[string(unicode.ToLower(v))] && unicode.IsLetter(v) {
			total := strings.Count(strings.ToLower(input), string(unicode.ToLower(v)))
			tmp[string(unicode.ToLower(v))] = true

			if total > 1 {
				result = result + string(v) + strconv.Itoa(total)
			} else {
				result = result + string(v)
			}
		}
	}

	fmt.Println("=== Convert String to Total Character ===")
	fmt.Println("Input : ", input)
	fmt.Println("Result : ", result)
	fmt.Println()

	return
}

func GetWordStatistic(input string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}

	processedString := reg.ReplaceAllString(string(b), " ")

	countWord := map[string]int{}
	for _, v := range strings.Fields(processedString) {
		countWord[v] = countWord[v] + 1
	}

	countWordOnlyShowOnce := 0
	tmpMapWordCount := map[string]int{}
	highsetWordCount := ""
	lowestWordCount := ""
	for i, v := range countWord {
		if v == 1 {
			countWordOnlyShowOnce++
		}

		if tmpMapWordCount["highest"] < v {
			highsetWordCount = i
			tmpMapWordCount["highest"] = v
		}

		if tmpMapWordCount["lowest"] > v || tmpMapWordCount["lowest"] == 0 {
			lowestWordCount = i
			tmpMapWordCount["lowest"] = v
		}
	}

	fmt.Println("=== Get Word Statistics ===")
	fmt.Println("Total Word : ", len(strings.Fields(processedString)))
	fmt.Println("Word Count Every Word : ", countWord)
	fmt.Println("Number of words that only show up once : ", countWordOnlyShowOnce)
	fmt.Println("Word that has the highest count : ", highsetWordCount)
	fmt.Println("Word that has the smallest count : ", lowestWordCount)
}

func main() {
	ReverseLowerAndUpperCaseChar("Lorem Ipsum Dolor Sit Amet")
	ConvertTotalCharInString("Team Engineering PT. Raksasa Laju Lintang")
	GetWordStatistic("example.txt")
}
