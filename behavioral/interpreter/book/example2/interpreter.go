package example2

type Interpreter interface {
  Read() int
}

const (
	SUM = "sum"
	SUB = "sub"
	MUL = "mul"
	DIV = "div"
)


type value int

func (v *value) Read() int{
	return int(*v)
}

type operation struct {
	Left Interpreter
	Right Interpreter
}

type operationSum struct {
  operation
}

type operationSub struct {
	operation
}
type operationMul struct {
	operation
}
type operationDiv struct {
	operation
}

func (o *operationSum) Read() int {
	return o.Left.Read() + o.Right.Read()
}

func (o *operationSub) Read() int {
	return o.Left.Read() - o.Right.Read()
}

func (o *operationMul) Read() int {
	return o.Left.Read() * o.Right.Read()
}

func (o *operationDiv) Read() int {
	return o.Left.Read() / o.Right.Read()
}

func operatorFactory(o string, left, right Interpreter) Interpreter {
	switch o {
	case SUM:
		return &operationSum{
		operation{
			Left:  left,
			Right: right,
			},
		}
	case SUB:
		return  &operationSub{
			operation{
				Left:  left,
				Right: right,
			},
		}
	}
	return nil
}

type polishNotationStack []Interpreter

func (p *polishNotationStack) Push(s Interpreter) {
	*p = append(*p,s)
}

func (p *polishNotationStack) Pop() Interpreter {
	length := len(*p)
	if length > 0 {
		temp := (*p)[length-1]
		*p = (*p)[:length-1]
		return temp
	}
	return nil
}


