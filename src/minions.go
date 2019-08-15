package minions

import "fmt"

type minion struct {
	min, max       int
	assignedToRoom bool
}

func main() {
	minions := readInput()
	numberOfRooms := assignToRooms(minions)
	fmt.Print(numberOfRooms)
}

func readInput() []minion {
	var numberOfMinions int
	fmt.Scan(&numberOfMinions)
	minions := make([]minion, numberOfMinions)
	for index := 0; index < numberOfMinions; index++ {
		var nextMinion minion
		fmt.Scan(&nextMinion.min)
		fmt.Scan(&nextMinion.max)
		minions[index] = nextMinion
	}
	return minions
}

func assignToRooms(minions []minion) int {
	numberOfRooms := 0
	for temperatureOfNextRoom := findSmallestUnassignedMax(minions); temperatureOfNextRoom > 0; temperatureOfNextRoom = findSmallestUnassignedMax(minions) {
		numberOfRooms++
		assignToRoom(minions, temperatureOfNextRoom)
		temperatureOfNextRoom = findSmallestUnassignedMax(minions)
	}
	return numberOfRooms
}

func assignToRoom(minions []minion, temperature int) {
	for index := 0; index < len(minions); index++ {
		_minion := &minions[index]
		if _minion.min <= temperature && temperature <= _minion.max {
			_minion.assignedToRoom = true
		}
	}
}

func findSmallestUnassignedMax(minions []minion) int {
	currentSmallestMax := 0
	for _, _minion := range minions {
		if !_minion.assignedToRoom {
			max := _minion.max
			if currentSmallestMax > max || currentSmallestMax == 0 {
				currentSmallestMax = max
			}
		}
	}
	return currentSmallestMax
}
