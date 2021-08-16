package gutil

import "strconv"

func ConvStr2Prom(str string) float64 {
	if str == "" {
		return 0
	}
	result, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return result
}
