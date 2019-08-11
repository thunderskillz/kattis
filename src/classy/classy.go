package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	numberOfCases, _ := strconv.Atoi(sc.Text())
	for i := 0; i < numberOfCases; i++ {
		persons := readInput(sc)
		sort.Strings(persons)
		printPersons(persons)
	}
}

func readInput(sc *bufio.Scanner) []string {
	sc.Scan()
	numberOfPersons, _ := strconv.Atoi(sc.Text())
	persons := make([]string, numberOfPersons)
	for j := 0; j < numberOfPersons; j++ {
		var name, class string
		sc.Scan()
		fmt.Sscanf(sc.Text(), "%s %s class\n", &name, &class)
		name = name[:len(name)-1]
		newPerson := createClassText(class) + name
		persons[j] = newPerson
	}
	return persons
}

func createClassText(class string) string {
	var classTextBuilder strings.Builder
	classTextBuilder.Grow(10)
	splitClass := strings.Split(class, "-")
	length := len(splitClass)
	for index := 0; index < 10; index++ {
		var value string
		if index < length {
			switch splitClass[length-index-1] {
			case "upper":
				value = "a"
			case "middle":
				value = "b"
			case "lower":
				value = "c"
			}
		} else {
			value = "b"
		}
		classTextBuilder.WriteString(value)
	}
	return classTextBuilder.String()
}

func printPersons(persons []string) {
	for _, person := range persons {
		name := person[10:]
		fmt.Println(name)
	}
	fmt.Println("==============================")
}
