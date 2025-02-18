package main

import "fmt"

func reformat(message string, formatter func(string) string) string {
	return fmt.Sprintf("TEXTIO: %s", formatter(formatter(formatter(message))))
}
