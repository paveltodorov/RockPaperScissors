package ai

import "fmt"

type AIError struct {
	Message string
}

func (e AIError) Error() string {
	return fmt.Sprintf("AI error: %s", e.Message)
}
