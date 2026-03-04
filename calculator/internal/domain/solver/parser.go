package solver

import (
	"strconv"
)

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

func parseKeyword(input string, i int, ans, preAns float64, variables map[string]float64) (float64, int, error) {
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
		return ans, i - 1, nil
	case "preAns":
		return preAns, i - 1, nil
	default:
		if val, ok := variables[keyword]; ok {
			return val, i - 1, nil
		}
		return 0, i, ErrInvalidExpression
	}
}
