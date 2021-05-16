package notify

import (
	"fmt"
	"io/ioutil"
	"os"
)

func prepareCodeBlock(message string, language string) string {
	return fmt.Sprintf("```%s\n%s\n```", language, message)
}

func extractFileContents(file *os.File) (string, error) {
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("ERROR|notify/util.extractFileContents()|failed to read file|%s", err.Error())
	}
	return string(fileBytes), nil
}