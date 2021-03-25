package example1

import (
	"strconv"
	"strings"
)

const (
	SUM = "sum"
	SUB = "sub"
	MUL = "mul"
	DIV = "div"
)

type polishNotationStack []int

func (p *polishNotationStack) Push(s int) {
	*p = append(*p,s)
}

func (p *polishNotationStack) Pop() int {
	length := len(*p)
	if length > 0 {
		temp := (*p)[length-1]
		*p = (*p)[:length-1]
		return temp
	}
	return 0
}


func Calculate(o string) (int, error) {
	stack := polishNotationStack{}
	terms := strings.Split(o, " ")
	for _, term := range terms {
		if isOperator(term){
			right := stack.Pop()
			left := stack.Pop()
			mathFunc := getOperationFunc(term)
			res := mathFunc(left, right)
			stack.Push(res)


		} else {
			val, err := strconv.Atoi(term)
			if err != nil {
				return 0, err
			}
			stack.Push(val)
		}
	}
	return stack.Pop(), nil
}

func isOperator(o string) bool {
	return o == SUM || o == SUB || o == MUL || o == DIV
}

func getOperationFunc(o string ) func(a,b int) int {
	switch o {
	case SUM:
		return func(a, b int) int {
			return a + b
		}
	case SUB:
		return func(a, b int) int {
			return a - b
		}
	case MUL:
		return func(a, b int) int {
			return a * b
		}
	case DIV:
		return func(a, b int) int {
			return a / b
		}
	}
	return nil
}


