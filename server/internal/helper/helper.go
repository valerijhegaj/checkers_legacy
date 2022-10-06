package helper

func StringToHttpStatus(status string) int {
	return int(status[0]-'0')*100 +
		int(status[1]-'0')*10 +
		int(status[2]-'0')
}

func HttpStatusToString(status int) string {
	return string(rune('0'+status/100)) +
		string(rune('0'+status%100/10)) +
		string(rune('0'+status%10))
}
