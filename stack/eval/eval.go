package eval

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// ValueStack stack of values extracted from expression and calculated in runtime
type ValueStack struct {
	Values [10]int
	Idx    int
}

// Push pushes new value onto the top of stack
func (v *ValueStack) Push(val int) {
	v.Values[v.Idx] = val
	v.Idx++
}

// Pop returns value from the top of stack
func (v *ValueStack) Pop() int {
	v.Idx--
	return v.Values[v.Idx]
}

// IsEmpty return boolean identifying is stack empty or not
func (v *ValueStack) IsEmpty() bool {
	return v.Idx == 0
}

// OperationStack is a stack of mathematical operators extracted from expression
type OperationStack struct {
	Values [10]string
	Idx    int
}

// Push pushes new value onto the top of stack
func (o *OperationStack) Push(val string) {
	o.Values[o.Idx] = val
	o.Idx++
}

// Pop returns value from the top of stack
func (o *OperationStack) Pop() string {
	o.Idx--
	return o.Values[o.Idx]
}

// IsEmpty return boolean identifying is stack empty or not
func (o *OperationStack) IsEmpty() bool {
	return o.Idx == 0
}

// Eval parses, evaluates and returns result of math expression
func Eval(expr string) int {
	vStack := &ValueStack{}
	opStack := &OperationStack{}

	for _, token := range expr {
		str := strings.Trim(fmt.Sprintf("%c", token), " ")

		if str == "(" || len(str) == 0 {
			continue
		}

		if str == ")" {
			op := opStack.Pop()
			val := vStack.Pop()
			res := 0
			switch op {
			case "+":
				res = vStack.Pop() + val
			case "-":
				res = vStack.Pop() - val
			case "*":
				res = vStack.Pop() * val
			case "s":
				res = int(math.Sqrt(float64(val)))
			case "/":
				res = vStack.Pop() / val
			}
			vStack.Push(res)
			continue
		}

		intVal, err := strconv.Atoi(str)

		if err != nil {
			opStack.Push(str)
			continue
		}

		vStack.Push(intVal)
	}

	return vStack.Pop()
}
