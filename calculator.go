package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

// recivers permite agregar una función al struct calc
func (calc) operate(input string, operator string) int {
	cleanInput := strings.Split(input, operator)
	value1 := parseInt(cleanInput[0])
	value2 := parseInt(cleanInput[1])

	switch operator {
	case "+":
		return (value1 + value2)
	case "-":
		return (value1 - value2)
	case "*":
		return (value1 * value2)
	case "/":
		return (value1 / value2)
	default:
		fmt.Println("Operador no soportado")
	}
	return 0
}

func parseInt(input string) int {
	value, _ := strconv.Atoi(input)
	return value
}

func ReadInput(message string) string {
	fmt.Println(message)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
