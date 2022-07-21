package main

import (
	"fmt"
	"os"
	"strconv"
)

type Arguments struct {
	Command string
	Domain  int
	Pair    int
}

type Calculate interface {
	Calc()
	Add() int
	Substract() int
	Multiply() int
	Divide() float32
}

func (input *Arguments) Add() int {
	return input.Domain + input.Pair
}

func (input *Arguments) Substract() int {
	return input.Domain - input.Pair
}

func (input *Arguments) Multiply() int {
	return input.Domain * input.Pair
}

func (input *Arguments) Divide() float32 {
	return float32(input.Domain) / float32(input.Pair)
}

func (input *Arguments) Calc() interface{} {
	switch input.Command {
	case "add":
		return input.Add()
	case "substract":
		return input.Substract()
	case "multiply":
		return input.Multiply()
	case "divide":
		return input.Divide()
	default:
		return "unimplemented"
	}
}

func HandleArgsInput() Arguments {
	args := os.Args
	domain, _ := strconv.Atoi(args[2])
	pair, _ := strconv.Atoi(args[3])
	problem := Arguments{
		Command: args[1],
		Domain:  domain,
		Pair:    pair,
	}

	return problem
}

func Count() {
	problem := HandleArgsInput()
	fmt.Println(problem.Calc())
}

func main() {
	Count()
}
