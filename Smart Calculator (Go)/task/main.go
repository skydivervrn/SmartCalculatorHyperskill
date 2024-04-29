package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var (
	result          Stack
	resultInt       StackInt
	operators       Stack
	variableStorage map[string]int
	priorityScore   = map[string]int{
		"+": 10,
		"-": 10,
		"*": 20,
		"/": 20,
		"^": 30,
	}
)

type Stack struct {
	storage []string
}

type StackInt struct {
	storage []int
}

func main() {
	variableStorage = make(map[string]int)
	for {
		result = Stack{}
		operators = Stack{}
		resultInt = StackInt{}
		userInput := askUserInput("")
		if userInput == "" {
			continue
		}
		if strings.Contains(userInput, "=") {
			assignment(userInput)
			continue
		}
		inputItems := parseInput(userInput)
		if !validateInput(inputItems) {
			continue
		}
		if len(inputItems) == 1 {
			handleSingeItemInput(inputItems[0])
			continue
		}
		err := buildPostfix(inputItems)
		if err != nil {
			fmt.Println(err)
			continue
		}
		//for _, j := range result.storage {
		//	fmt.Printf("'%s'\n", j)
		//}
		//fmt.Println(result.storage)
		fmt.Println(calculatePostfix())
	}
}

func buildPostfix(str []string) error {
	for _, n := range str {
		if isNumber(n) {
			result.push(n)
			continue
		}
		if isLetter(n) {
			result.push(n)
			continue
		}
		operator, err := defineOperator(n)
		if err != nil {
			return err
		}
		checkTopOperator(operator)
	}
	if len(operators.storage) > 0 {
		for range operators.storage {
			value, _ := operators.pop()
			result.storage = append(result.storage, value)
		}
	}
	return nil
}

func validateInput(str []string) bool {
	leftParenthesis := 0
	rightParenthesis := 0
	for _, n := range str {
		if n == "(" {
			leftParenthesis++
		}
		if n == ")" {
			rightParenthesis++
		}
	}
	if leftParenthesis != rightParenthesis {
		fmt.Println("Invalid expression")
		return false
	}
	return true
}

//>  a = 7
//>  b = 2
//> a * 4 / b - (3 - 1)
//> 7 * 4 / 2 - (3 - 1)

func parseInput(str string) []string {
	var res []string
	for _, v := range strings.Split(strings.TrimSpace(str), " ") {
		if strings.HasPrefix(v, "(") {
			for x, y := range v {
				if y == '(' {
					res = append(res, string(y))
					continue
				} else {
					res = append(res, v[x:])
					break
				}
			}
			continue
		}
		if strings.HasSuffix(v, ")") {
			first := true
			for x, y := range v {
				if y == ')' && first {
					res = append(res, v[:x])
					res = append(res, string(y))
					first = false
					continue
				}
			}
			continue
		}
		res = append(res, v)
	}
	//fmt.Println(res)
	return res
}

func calculatePostfix() int {
	for _, n := range result.storage {
		if isNumber(n) {
			number, _ := strconv.Atoi(n)
			resultInt.push(number)
			continue
		}
		if isLetter(n) {
			if varHasValue(n) {
				number, _ := variableStorage[n]
				resultInt.push(number)
				continue
			}
		}
		if n == "+" {
			num1, _ := resultInt.pop()
			num2, _ := resultInt.pop()
			resultInt.push(num1 + num2)
		}
		if n == "-" {
			num1, _ := resultInt.pop()
			num2, _ := resultInt.pop()
			resultInt.push(num2 - num1)
		}
		if n == "*" {
			num1, _ := resultInt.pop()
			num2, _ := resultInt.pop()
			resultInt.push(num1 * num2)
		}
		if n == "/" {
			num1, _ := resultInt.pop()
			num2, _ := resultInt.pop()
			resultInt.push(num2 / num1)
		}
		//fmt.Println(resultInt)
	}
	res, _ := resultInt.pop()
	return res
}

func checkTopOperator(operator string) {
	if operator == "(" {
		operators.push(operator)
		return
	}
	for {
		topOperator, err := operators.showTop()
		if err != nil {
			if err.Error() == "stack is empty" {
				operators.push(operator)
				return
			}
		}
		if topOperator == "(" {
			if operator == ")" {
				operators.pop()
				return
			}
			operators.push(operator)
			return
		}
		if operator != ")" && priorityScore[operator] > priorityScore[topOperator] {
			operators.push(operator)
			return
		} else {
			if topOperator == "(" {
				operators.pop()
			}
			val, _ := operators.pop()
			result.push(val)
		}
	}
}

func (s *StackInt) push(value int) {
	s.storage = append(s.storage, value)
}

func (s *StackInt) pop() (int, error) {
	last := len(s.storage) - 1
	if last < 0 { // check the size
		return 0, errors.New("stack is empty") // and return error
	}

	value := s.storage[last]     // save the value
	s.storage = s.storage[:last] // remove the last element

	return value, nil // return saved value and nil error
}

func (s *StackInt) showTop() (int, error) {
	last := len(s.storage) - 1
	if last <= -1 { // check the size
		return 0, errors.New("stack is empty") // and return error
	}
	value := s.storage[last] // save the value

	return value, nil // return saved value and nil error
}

func (s *Stack) push(value string) {
	s.storage = append(s.storage, value)
}

func (s *Stack) pop() (string, error) {
	last := len(s.storage) - 1
	if last < 0 { // check the size
		return "", errors.New("stack is empty") // and return error
	}

	value := s.storage[last]     // save the value
	s.storage = s.storage[:last] // remove the last element

	return value, nil // return saved value and nil error
}

func (s *Stack) showTop() (string, error) {
	last := len(s.storage) - 1
	if last <= -1 { // check the size
		return "", errors.New("stack is empty") // and return error
	}
	value := s.storage[last] // save the value

	return value, nil // return saved value and nil error
}

func varHasValue(s string) bool {
	_, ok := variableStorage[s]
	if ok {
		return true
	}
	return false
}

func isLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func isNumber(s string) bool {
	for _, r := range s {
		if !unicode.IsNumber(r) {
			return false
		}
	}
	return true
}

func defineOperator(str string) (string, error) {
	if len(str) == 1 {
		switch str {
		case "+":
			return "+", nil
		case "-":
			return "-", nil
		case "*":
			return "*", nil
		case "/":
			return "/", nil
		case "(":
			return "(", nil
		case ")":
			return ")", nil
		default:
			return "", errors.New("Invalid expression")
		}
	}
	minusCounter := 0
	for _, char := range str {
		if char != '+' && char != '-' {
			return "", errors.New("Invalid expression")
		}
		if char == '-' {
			minusCounter++
			continue
		}
		continue
	}
	if minusCounter%2 != 0 {
		return "-", nil
	} else {
		return "+", nil
	}
}

func subtraction(num1, num2 int) int {
	return num1 - num2
}

func addition(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func askUserInput(str string) string {
	reader := bufio.NewReader(os.Stdin)
	//fmt.Println(str)
	buf, _, err := reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	return string(buf)
}

func assignment(userInput string) {
	variablesFromInput := strings.Split(strings.TrimSpace(userInput), "=")
	for s, _ := range variablesFromInput {
		variablesFromInput[s] = strings.Replace(variablesFromInput[s], " ", "", -1)
	}
	if isLetter(variablesFromInput[0]) {
		if isNumber(variablesFromInput[1]) {
			intValue, _ := strconv.Atoi(variablesFromInput[1])
			variableStorage[variablesFromInput[0]] = intValue
			return
		}
		if isLetter(variablesFromInput[1]) {
			if varHasValue(variablesFromInput[1]) {
				variableStorage[variablesFromInput[0]] = variableStorage[variablesFromInput[1]]
				return
			}
		}
		fmt.Println("Invalid assignment")
		return
	}
	fmt.Println("Invalid identifier")
	return
}

func handleSingeItemInput(item string) {
	if item == "/exit" {
		fmt.Println("Bye!")
		os.Exit(0)
	}
	if item == "/help" {
		fmt.Println("The program calculates the sum of numbers")
		return
	}
	if isLetter(item) {
		if varHasValue(item) {
			fmt.Println(variableStorage[item])
			return
		}
	}
	if isNumber(item) {
		fmt.Println(item)
		return
	}
	if varHasValue(item) {
		fmt.Println(variableStorage[item])
		return
	}
	fmt.Println("Unknown command")
	return
}
