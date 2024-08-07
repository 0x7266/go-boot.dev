package main

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	return len(mToSend.sender.name) > 0 && mToSend.sender.number > 0 && len(mToSend.recipient.name) > 0 && mToSend.recipient.number > 0
}
