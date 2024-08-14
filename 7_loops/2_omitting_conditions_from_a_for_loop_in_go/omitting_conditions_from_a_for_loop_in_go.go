package main

func maxMessages(thresh int) int {
	var total int
	for i := 0; ; i++ {
		total += 100 + i
		if total > thresh {
			return i
		}
	}
}
