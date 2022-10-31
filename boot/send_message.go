package boot

import (
	"encoding/json"
	"net/http"
	"strings"
)

func SendGroupMsg(groupID string, message string) error {
	body, _ := json.Marshal(map[string]interface{}{
		"group_id":    groupID,
		"message":     message,
		"auto_escape": false,
	})
	_, err := http.DefaultClient.Post("http://127.0.0.1:5700/send_group_msg", "application/json", strings.NewReader(string(body)))
	return err
}
