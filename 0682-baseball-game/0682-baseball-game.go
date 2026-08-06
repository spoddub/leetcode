func calPoints(operations []string) int {
	record := make([]int, 0)

	for _, operation := range operations {
		switch operation {
		case "+":
			num := record[len(record)-1] + record[len(record)-2]
			record = append(record, num)
		case "D":
			num := record[len(record)-1] * 2
			record = append(record, num)
		case "C":
			record = record[:len(record)-1]
		default:
			num, _ := strconv.Atoi(operation)
			record = append(record, num)
		}
	}

	counter := 0

	for _, score := range record {
		counter += score
	}

	return counter
}