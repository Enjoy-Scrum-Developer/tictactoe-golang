package helpers

func GenerateArray(size int, symbol string) []string {
	array := make([]string, size)
	for i := range array {
		array[i] = symbol
	}
	return array
}
