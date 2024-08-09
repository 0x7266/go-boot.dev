package main

type UserInputError struct {
	message string
}

func (uie UserInputError) Error() string {
	return uie.message
}

func validateStatus(status string) error {
	if status == "" {
		return UserInputError{"status cannot be empty"}
	}
	if len(status) > 140 {
		return UserInputError{"status exceeds 140 characters"}
	}
	return nil
}
