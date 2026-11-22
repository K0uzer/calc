package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func add(a, b float64) (float64, error) {
	return a + b, nil
}

func subtract(a, b float64) (float64, error) {
	return a - b, nil
}

func multiply(a, b float64) (float64, error) {

	if a == 0 || b == 0 {
		return 0, errors.New("Умножение на 0")
	}

	return a * b, nil
}

func divide(a, b float64) (float64, error) {

	if a == 0 || b == 0 {
		return 0, errors.New("Деление на ноль")
	}

	return a / b, nil
}

func greet(name string) (string, error) {
    if len(name) == 0 { return "", errors.New("Строка пуста") } 
    return fmt.Sprintf("Привет, %s!", name), nil
}

func validateAge(age int) error {
    if age > 150 || age < 150 { return errors.New("Возраст не прошел валидацию") }
	return nil
}

func safeSqrt(x float64) (float64, error) {
    if x < 0 { return 0, errors.New("Число меньше нуля") }
	return math.Sqrt(x), nil
}

func readNumber(prompt string) (float64, error) {
var input string
    fmt.Print(prompt)
    if _, err := fmt.Scan(&input); err != nil {
        return 0, fmt.Errorf("ошибка ввода: %w", err)
    }
    n, err := strconv.ParseFloat(input, 64)
    if err != nil {
        return 0, fmt.Errorf("%q — не число", input)
    }
    return n, nil
}

func main() {
	firstNumber, err := readNumber("12")
	if err != nil {
        fmt.Println("Ошибка:", err)
        return
    }
	secondNumber, err := readNumber("15")
	if err != nil {
        fmt.Println("Ошибка:", err)
        return
    }
	var operation string
	var result float64
	var state error
	fmt.Println("Введите операцию:")
	fmt.Scan(&operation)

	switch {
		case operation == "деление":
			result, state = divide(firstNumber, secondNumber)
		case operation == "сложение": 
			result, _ = add(firstNumber, secondNumber)
		case operation == "умножение": 
			result, state = multiply(firstNumber, secondNumber)
		case operation == "вычетание": 
			result, _ = subtract(firstNumber, secondNumber)
		default: 
			fmt.Println("Такой операции нет")
			return
	}

	if state != nil {
		fmt.Println(state)
		return
	} else {
		fmt.Printf("Ваш результат: %.1f", result)
	}

}