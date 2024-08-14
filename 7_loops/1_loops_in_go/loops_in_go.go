package main

func bulkSend(numMessages int) float64 {
	var total float64
	for i := range numMessages {
		total += 1.0 + float64(i)/100.0
	}
	return total
}
