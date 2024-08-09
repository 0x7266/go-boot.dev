package main

const SYSTEM_ALERT_IMPORTANCE_SCORE int = 100

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

func (dm directMessage) importance() int {
	if dm.isUrgent {
		return 50
	}
	return dm.priorityLevel
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

func (gm groupMessage) importance() int {
	return gm.priorityLevel
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (sa systemAlert) importance() int {
	return SYSTEM_ALERT_IMPORTANCE_SCORE
}

func processNotification(n notification) (string, int) {
	switch n := n.(type) {
	case directMessage:
		return n.senderUsername, n.importance()
	case groupMessage:
		return n.groupName, n.importance()
	case systemAlert:
		return n.alertCode, n.importance()
	default:
		return "", 0
	}
}
