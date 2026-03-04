package util

import (
	"strconv"
	"strings"
)

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
