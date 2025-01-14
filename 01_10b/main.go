package main

import (
	"flag"
	"log"
	"strconv"
	"strings"
)

// operators is the map of legal operators and their functions
var operators = map[string]func(x, y float64) float64{
	"+": func(x, y float64) float64 { return x + y },
	"-": func(x, y float64) float64 { return x - y },
	"*": func(x, y float64) float64 { return x * y },
	"/": func(x, y float64) float64 { return x / y },
}

// parseOperand parses a string to a float64
func parseOperand(op string) float64 {
	parsedOp, error := strconv.ParseFloat(op, 64)
	if error != nil {
		log.Fatal(error)
	}
	return parsedOp
}

// calculate returns the result of a 2 operand mathematical expression
func calculate(expr string) float64 {
	ops := strings.Fields(expr)
	if len(ops) != 3 {
		log.Fatal("Bad expression. It requieres three arguments")
	}
	left := parseOperand(ops[0])
	right := parseOperand(ops[2])
	f := operators[ops[1]]
	if (f == nil) {
		log.Fatal("Bad function operator. It only supports basic operators (+, -, *, /)")
	}
	result := f(left, right)
	return result
}

func main() {
	expr := flag.String("expr", "",
		"The expression to calculate on, separated by spaces.")
	flag.Parse()
	result := calculate(*expr)
	log.Printf("%s = %.2f\n", *expr, result)
}
