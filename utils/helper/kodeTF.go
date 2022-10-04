package helper

import (
	"strconv"
	"time"
)

func GenerateTF(userID int) string {
	unix := strconv.Itoa(int(time.Now().Unix()))
	uid := strconv.Itoa(userID)
	return "TF-" + unix + uid
}
