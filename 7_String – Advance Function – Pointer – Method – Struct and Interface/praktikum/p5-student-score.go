package main

import "fmt"

type Student struct {
	name  []string
	score []int
}

func (s Student) average() float64 {
	var result int
	for _, score := range s.score {
		result += score
	}
	return float64(result) / float64(len(s.score))
}

func (s Student) min() (score int, name string) {
	var index int = 0
	var min int = s.score[0]

	for i, _score := range s.score {
		if min > _score {
			min = _score
			index = i
		}
	}

	return s.score[index], s.name[index]
}

func (s Student) max() (score int, name string) {
	var index int = 0
	var max int = s.score[0]

	for i, _score := range s.score {
		if max < _score {
			max = _score
			index = i
		}
	}

	return s.score[index], s.name[index]
}

func main() {
	var inputs = Student{}

	for i := 0; i < 3; i++ {
		var name string
		fmt.Print("Input ", i+1, " student's name: ")
		fmt.Scan(&name)
		inputs.name = append(inputs.name, name)

		var score int
		fmt.Print("Input ", i+1, " student's score: ")
		fmt.Scan(&score)
		inputs.score = append(inputs.score, score)
	}

	fmt.Println("\n\nAverage of student's score:", inputs.average())
	maxScore, maxName := inputs.max()
	fmt.Println("Highest score of the student is", maxName, "with", maxScore, "score.")
	minScore, minName := inputs.min()
	fmt.Println("Lowest score of the student is", minName, "with", minScore, "score.")
}
