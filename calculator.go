package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

var romanToArabic = map[string]int{
    "I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
    "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var arabicToRoman = []struct {
    Value  int
    Symbol string
}{
    {1000, "M"},
    {900, "CM"},
    {500, "D"},
    {400, "CD"},
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

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Введите выражение:")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)

    result, err := calculate(input)
    if err != nil {
        panic(err)
    }

    fmt.Println("Результат:", result)
}

func calculate(input string) (string, error) {
    var a, b int
    var op string
    var isRoman bool

    // Проверка, что введенное выражение содержит только один оператор
    if strings.Count(input, "+") + strings.Count(input, "-") + strings.Count(input, "*") + strings.Count(input, "/") != 1 {
        return "", fmt.Errorf("недопустимый формат ввода")
    }

    // Проверка, используются ли римские числа
    for roman := range romanToArabic {
        if strings.Contains(input, roman) {
            isRoman = true
            break
        }
    }

    if isRoman {
        var aRoman, bRoman string
        aRoman, op, bRoman = parseInput(input)
        var ok bool
        if a, ok = romanToArabic[aRoman]; !ok {
            return "", fmt.Errorf("недопустимое римское число: %s", aRoman)
        }
        if b, ok = romanToArabic[bRoman]; !ok {
            return "", fmt.Errorf("недопустимое римское число: %s", bRoman)
        }
    } else {
        _, err := fmt.Sscanf(input, "%d %s %d", &a, &op, &b)
        if err != nil {
            return "", fmt.Errorf("недопустимый формат ввода")
        }
    }

    if a < 1 || a > 10 || b < 1 || b > 10 {
        return "", fmt.Errorf("числа должны быть от 1 до 10 включительно")
    }

    var result int
    switch op {
    case "+":
        result = a + b
    case "-":
        result = a - b
    case "*":
        result = a * b
    case "/":
        if b == 0 {
            return "", fmt.Errorf("деление на ноль")
        }
        result = a / b
    default:
        return "", fmt.Errorf("недопустимая операция")
    }

    if isRoman {
        if result < 1 {
            return "", fmt.Errorf("римское число не может быть меньше I")
        }
        return toRoman(result), nil
    }
    return strconv.Itoa(result), nil
}

func parseInput(input string) (string, string, string) {
    input = strings.ReplaceAll(input, " ", "")
    for _, op := range []string{"+", "-", "*", "/"} {
        if strings.Contains(input, op) {
            parts := strings.Split(input, op)
            return parts[0], op, parts[1]
        }
    }
    panic("недопустимый формат ввода")
}

func toRoman(num int) string {
    var result strings.Builder
    for _, entry := range arabicToRoman {
        for num >= entry.Value {
            result.WriteString(entry.Symbol)
            num -= entry.Value
        }
    }
    return result.String()
}
