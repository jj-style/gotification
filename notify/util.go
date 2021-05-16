package notify

import (
	"fmt"
)

func prepareCodeBlock(message string, language string) string {
	return fmt.Sprintf("```%s\n%s\n```", language, message)
}