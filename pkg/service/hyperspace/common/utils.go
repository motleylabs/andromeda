package common

import "fmt"

func ConvertLamports(sol float64) int64 {
	return int64(sol * LAMPORTS_PER_SOL)
}

func GetLamports(sol float64) string {
	return fmt.Sprintf("%d", ConvertLamports(sol))
}

func GetLamportsFromPointer(sol *float64) string {
	output := "0"
	if sol != nil {
		output = fmt.Sprintf("%d", ConvertLamports(*sol))
	}
	return output
}

func GetFromIntPointer(val *int) string {
	output := "0"
	if val != nil {
		output = fmt.Sprintf("%d", *val)
	}
	return output
}

func GetFromFloatPointer(val *float64) string {
	output := "0"
	if val != nil {
		output = fmt.Sprintf("%f", *val)
	}
	return output
}
