package cd

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var numberOfCDs1, numberOfCDs2 int
	var cds []int
	readNumberOfCDs(&numberOfCDs1, &numberOfCDs2, scanner)
	for numberOfCDs1 != 0 && numberOfCDs2 != 0 {
		readFirstCDs(numberOfCDs1, &cds, scanner)
		fmt.Println(findNumberOfMatches(numberOfCDs2, cds, scanner))
		readNumberOfCDs(&numberOfCDs1, &numberOfCDs2, scanner)
	}
}

func readNumberOfCDs(numberOfCDs1 *int, numberOfCDs2 *int, scanner *bufio.Scanner) {
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d %d", numberOfCDs1, numberOfCDs2)
}

func readFirstCDs(numberOfCDs int, cds *[]int, scanner *bufio.Scanner) {
	*cds = make([]int, numberOfCDs)
	for index := 0; index < numberOfCDs; index++ {
		scanner.Scan()
		(*cds)[index], _ = strconv.Atoi(scanner.Text())
	}
}

func findNumberOfMatches(numberOfCDs int, cds []int, scanner *bufio.Scanner) int {
	cdCounter := 0
	for index := 0; index < numberOfCDs; index++ {
		scanner.Scan()
		catalogNumber, _ := strconv.Atoi(scanner.Text())
		i := sort.Search(len(cds), func(j int) bool {
			return cds[j] >= catalogNumber
		})
		if i < len(cds) && cds[i] == catalogNumber {
			cdCounter++
		}
		cds = cds[i:]
	}
	return cdCounter
}
