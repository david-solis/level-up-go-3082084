package main

import (
	"flag"
	"log"
)

// isBalanced returns whether the given expression
// has balanced brackets.
func isBalanced(expr string) bool {
    var stack []rune
    for _, bracket := range expr {
        switch bracket {
        case '(', '{', '[':
            stack = append(stack, bracket)
        case ')', '}', ']':
            if len(stack) == 0 {
                return false
            }
            if bracket == ')' && stack[len(stack)-1] != '(' {
                return false
            }
            if bracket == '}' && stack[len(stack)-1] != '{' {
                return false
            }
            if bracket == ']' && stack[len(stack)-1] != '[' {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }
    return len(stack) == 0
}



// printResult prints whether the expression is balanced.
func printResult(expr string, balanced bool){ 
	if balanced {
		log.Printf("%s is balanced.\n", expr)
		return
	}
	log.Printf("%s is not balanced.\n", expr)
}

func main() {
	expr := flag.String("expr", "", "The expression to validate brackets on.")
	flag.Parse()
	printResult(*expr, isBalanced(*expr))
}
