package logs

import "fmt"

// TODO checkout logsrus
func DebugMessage(status int, msg string) string {
	debug := fmt.Sprintf("********** DEBUG: (%d) %v", status, msg)
	fmt.Println(debug)
	return debug
}
