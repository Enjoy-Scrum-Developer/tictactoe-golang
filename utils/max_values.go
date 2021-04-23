package utils

func MaxValues(values []int) (int, IntegerLinkedList) {
	maxValue := 0
	list := NewIntegerLinkedList()

	for i := 0; i < len(values); i++ {
		value := values[i]
		if value > maxValue {
			maxValue = value
		}
	}

	for i := 0; i < len(values); i++ {
		value := values[i]
		if value == maxValue {
			list.Add(i)
		}
	}

	return maxValue, *list
}
