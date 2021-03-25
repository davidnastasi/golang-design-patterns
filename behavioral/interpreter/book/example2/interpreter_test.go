package example2

import (
	"strconv"
	"strings"
	"testing"
)

func TestInterpreter(t *testing.T)  {

	stack := polishNotationStack{}
	terms := strings.Split("3 4 sum 2 sub", " ")
	for _, term := range terms {
		if term == SUM || term == SUB {
			right := stack.Pop()
			left := stack.Pop()
			mathFunc := operatorFactory(term, left, right)
			res := value(mathFunc.Read())
			stack.Push(&res)
		} else {
			val, err := strconv.Atoi(term)
			if err != nil {
				t.Error(err)
			}
			temp := value(val)
			stack.Push(&temp)
		}
	}

	if res := stack.Pop().Read(); res != 5 {
		t.Errorf("exepected result not found: %d != %d\n", 5, res)
	}



}
