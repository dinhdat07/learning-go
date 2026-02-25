package processor

import (
	calError "calculator/error"
	"strconv"
	"strings"
)

var ErrInvalidFloatList = &calError.SyntaxError{
	Message: "invalid float list",
}

func parseNumber(input string, i int) (float64, int, error) {
	start := i
	dot := 0

	for i < len(input) {
		ch := input[i]

		if ch == '.' {
			dot++
			if dot > 1 {
				return 0, i, ErrInvalidExpression
			}
			i++
			continue
		}

		if ch < '0' || ch > '9' {
			break
		}

		i++
	}

	numStr := input[start:i]
	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return 0, i, ErrInvalidExpression
	}
	return num, i - 1, nil
}

func parseKeyword(calculator *Calculator, input string, i int) (float64, int, error) {
	start := i

	for i < len(input) {
		ch := input[i]
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
			i++
		} else {
			break
		}
	}

	keyword := input[start:i]

	switch keyword {
	case "ans":
		return calculator.ans, i - 1, nil
	case "preAns":
		return calculator.preAns, i - 1, nil
	default:
		if val, ok := calculator.variables[keyword]; ok {
			return val, i - 1, nil
		}
		return 0, i, ErrInvalidExpression
	}
}

func ParseFloatList(s string) ([]float64, error) {
	fields := strings.Fields(s)
	output := make([]float64, 0, len(fields))

	for _, v := range fields {
		num, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, ErrInvalidFloatList
		}
		output = append(output, num)
	}

	return output, nil
}
