package main

import (
	"fmt"
	"math"
)

var paperLengths = make([]float64, 30)

//AOnePaper Calculates how much tape needs to be used when taping papers of size A2 and smaller into an A1 paper.
func AOnePaper() {
	smallestSize, papers := readInput()
	usedPapers, possible := calculateUsedPapers(smallestSize, papers)
	if possible {
		tapeLengthNeeded := calculateTapeNeeded(usedPapers)
		fmt.Print(tapeLengthNeeded)
	} else {
		fmt.Print("impossible")
	}
}

func readInput() (int, []int) {
	var smallestSize int
	fmt.Scan(&smallestSize)
	papers := make([]int, smallestSize-1)
	for index := 0; index < smallestSize-1; index++ {
		fmt.Scan(&papers[index])
	}
	return smallestSize, papers
}

func calculateUsedPapers(smallestSize int, papers []int) ([]int, bool) {
	usedPapers := make([]int, smallestSize-1)
	paperSizeNeeded := int(math.Pow(2, float64(smallestSize-1)))

	for index, numberOfPapers := range papers {
		conversionFactor := int(math.Pow(2, float64(smallestSize-index-2)))
		sizeOfPapers := numberOfPapers * conversionFactor
		if paperSizeNeeded > sizeOfPapers {
			paperSizeNeeded -= sizeOfPapers
			usedPapers[index] = numberOfPapers
		} else {
			numberOfPapersNeeded := int(paperSizeNeeded / conversionFactor)
			usedPapers[index] = numberOfPapersNeeded
			return usedPapers, true
		}
	}

	return usedPapers, false
}

func calculateTapeNeeded(papers []int) float64 {
	var totalTapeNeeded float64
	for index, amountOfPaper := range papers {
		if amountOfPaper > 0 {
			for i := index; i >= 0; i-- {
				totalTapeNeeded += float64(amountOfPaper) * calculateLengthOfPaper(i) / math.Pow(2, float64(index-i+1))
			}
		}
	}
	return totalTapeNeeded
}

func calculateLengthOfPaper(size int) float64 {
	paperLength := paperLengths[size]
	if paperLength != 0 {
		return paperLength
	}
	paperLength = math.Pow(2, (-3-2*float64(size))/4)
	paperLengths[size] = paperLength
	return paperLength
}
