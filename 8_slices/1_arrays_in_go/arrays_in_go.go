package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	primary_len := len(primary)
	secondary_len := len(secondary) + primary_len
	tertiary_len := len(tertiary) + secondary_len
	return [3]string{primary, secondary, tertiary}, [3]int{primary_len, secondary_len, tertiary_len}
}
