package wordy

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	// NUMBERS all numbers in the system
	NUMBERS = `-?\d+`

	//Operators Operators
	Operators = `[\+-/*]?`

	//Question Question
	Question = `\d+\?$`
)

// getNum Get the number
func getNum(n string) (int, error) {
	val, err := strconv.Atoi(n)
	return val, err
}

//isQuestion Check if it is a question
func isQuestion(q string) bool {
	question := regexp.MustCompile(Question)
	return question.Match([]byte(q))
}

// Answer Answers the statement.
func Answer(q string) (int, bool) {
	var stack []string = make([]string, 0)
	var sum int = 0
	numbers := regexp.MustCompile(NUMBERS)
	operators := regexp.MustCompile(Operators)

	// Matches the query
	q = strings.TrimSpace(q)

	if !isQuestion(q) {
		return 0, false
	}

	var elements []string = strings.Split(q, " ")
    log.Printf("split: %v",elements)

	for _, v := range elements {
        log.Printf("item: %s",v)
		switch v {
		case "plus":
			stack = append(stack, "+")
		case "minus":
			stack = append(stack, "-")
		case "multiplied":
			stack = append(stack, "*")
		case "divided":
			stack = append(stack, "/")
		default:
			n := numbers.FindAllString(v, -1)
			if len(n) == 1 {
				stack = append(stack, n[0])
			}
		}
	}

	// No valid data found
	if len(stack) == 0 {
		return 0, false
	}

	for i, el := range stack {
		// If it is the first element
		if i == 0 {
			val, err := getNum(el)
			if err != nil {
				return 0, false
			}
			sum = val
		} else {

			if operators.Match([]byte(el)) {
				switch el {
				case "+":
					val, err := getNum(stack[i+1])
					if err != nil {
						return 0, false
					}
					sum += val
				case "-":
					val, err := getNum(stack[i+1])
					if err != nil {
						return 0, false
					}
					sum -= val
				case "/":
					val, err := getNum(stack[i+1])
					if err != nil {
						return 0, false
					}
					sum /= val
				case "*":
					val, err := getNum(stack[i+1])
					if err != nil {
						return 0, false
					}
					sum *= val
				default:
					if i%2 != 0 {
						return 0, false
					}
				}
			}
		}
	}

	log.Println(" Read: ", stack)
	return sum, true
}




package wordy

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: d137db1 Format using prettier (#1917)

type wordyTest struct {
	description string
	question    string
	expectError bool
	expected    int
}

var tests = []wordyTest{
	{
		description: "just a number",
		question:    "What is 5?",
		expectError: false,
		expected:    5,
	},
	{
		description: "addition",
		question:    "What is 1 plus 1?",
		expectError: false,
		expected:    2,
	},
	{
		description: "more addition",
		question:    "What is 53 plus 2?",
		expectError: false,
		expected:    55,
	},
	{
		description: "addition with negative numbers",
		question:    "What is -1 plus -10?",
		expectError: false,
		expected:    -11,
	},
	{
		description: "large addition",
		question:    "What is 123 plus 45678?",
		expectError: false,
		expected:    45801,
	},
	{
		description: "subtraction",
		question:    "What is 4 minus -12?",
		expectError: false,
		expected:    16,
	},
	{
		description: "multiplication",
		question:    "What is -3 multiplied by 25?",
		expectError: false,
		expected:    -75,
	},
	{
		description: "division",
		question:    "What is 33 divided by -3?",
		expectError: false,
		expected:    -11,
	},
	{
		description: "multiple additions",
		question:    "What is 1 plus 1 plus 1?",
		expectError: false,
		expected:    3,
	},
	{
		description: "addition and subtraction",
		question:    "What is 1 plus 5 minus -2?",
		expectError: false,
		expected:    8,
	},
	{
		description: "multiple subtraction",
		question:    "What is 20 minus 4 minus 13?",
		expectError: false,
		expected:    3,
	},
	{
		description: "subtraction then addition",
		question:    "What is 17 minus 6 plus 3?",
		expectError: false,
		expected:    14,
	},
	{
		description: "multiple multiplication",
		question:    "What is 2 multiplied by -2 multiplied by 3?",
		expectError: false,
		expected:    -12,
	},
	{
		description: "addition and multiplication",
		question:    "What is -3 plus 7 multiplied by -2?",
		expectError: false,
		expected:    -8,
	},
	{
		description: "multiple division",
		question:    "What is -12 divided by 2 divided by -3?",
		expectError: false,
		expected:    2,
	},
	{
		description: "unknown operation",
		question:    "What is 52 cubed?",
		expectError: true,
		expected:    0,
	},
	{
		description: "Non math question",
		question:    "Who is the President of the United States?",
		expectError: true,
		expected:    0,
	},
	{
		description: "reject problem missing an operand",
		question:    "What is 1 plus?",
		expectError: true,
		expected:    0,
	},
	{
		description: "reject problem with no operands or operators",
		question:    "What is?",
		expectError: true,
		expected:    0,
	},
	{
		description: "reject two operations in a row",
		question:    "What is 1 plus plus 2?",
		expectError: true,
		expected:    0,
	},
	{
		description: "reject two numbers in a row",
		question:    "What is 1 plus 2 1?",
		expectError: true,
		expected:    0,
	},
	{
		description: "reject postfix notation",
		question:    "What is 1 2 plus?",
		expectError: true,
		expected:    0,
	},
	{
		description: "reject prefix notation",
		question:    "What is plus 1 2?",
		expectError: true,
		expected:    0,
	},
}
