package main

func countConnections(groupSize int) int {
	var total int
	for i := range groupSize {
		if i == 0 {
			continue
		}
		total += groupSize - i
	}
	return total
}
