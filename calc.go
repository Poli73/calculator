package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func calculator(num1 int, num2 int, operation string) (int, error) {
	switch operation {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		return num1 / num2, nil
	default:
		return 0, errors.New("unknown operation")
	}
}

func convert(any string) (int, string) {
	switch any {
	case "1":
		return 1, "arabic"
	case "2":
		return 2, "arabic"
	case "3":
		return 3, "arabic"
	case "4":
		return 4, "arabic"
	case "5":
		return 5, "arabic"
	case "6":
		return 6, "arabic"
	case "7":
		return 7, "arabic"
	case "8":
		return 8, "arabic"
	case "9":
		return 9, "arabic"
	case "10":
		return 10, "arabic"
	case "I":
		return 1, "roman"
	case "II":
		return 2, "roman"
	case "III":
		return 3, "roman"
	case "IV":
		return 4, "roman"
	case "V":
		return 5, "roman"
	case "VI":
		return 6, "roman"
	case "VII":
		return 7, "roman"
	case "VIII":
		return 8, "roman"
	case "IX":
		return 9, "roman"
	case "X":
		return 10, "roman"
	default:
		panic("Неподходящее число")
	}
}
func ArabicToRoman(number int) string {

	conversions := []struct {
		value int
		digit string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}
	return roman.String()
}

func main() {

	fmt.Println("Введите первое и второе число, введите операцию (+, -, *, /): ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	s := strings.Split(text, " ")

	if len(s) > 3 {
		panic("Много входящих параметров")
	} else if len(s) < 3 {
		panic("Недостаточно входящих параметров")
	}

	n1, t1 := convert(s[0])
	n2, t2 := convert(s[2])

	if t1 != t2 {
		fmt.Println("Пожалуйста, используйте только арабские или только римские цифры от 1 до 10")
		return
	}
	result, err := calculator(n1, n2, s[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	if t1 == "roman" && (result) <= 0 {
		panic("В римской системе нет отрицательных чисел")
	} else if t1 == "roman" && (result) > 0 {
		fmt.Println("Результат:", ArabicToRoman(result))
	} else {
		fmt.Println("Результат:", result)
	}

}
