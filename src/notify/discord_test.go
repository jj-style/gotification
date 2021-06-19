package notify

import (
	"testing"
)

const (
	goCodeBlockIn = `package main
import "fmt"
func main() {
	fmt.Println("Hello world!")
}`
	goCodeBlockOut    = "```go\n" + goCodeBlockIn + "\n```"
	pythonCodeBlockIn = `from random import randint
print(f"Your number is {randint(1,10)}")`
	pythonCodeBlockOut = "```python\n" + pythonCodeBlockIn + "\n```"
)

func TestPrepareCodeBlock(t *testing.T) {
	testCases := []struct {
		input    string
		language string
		output   string
	}{
		{goCodeBlockIn, "go", goCodeBlockOut},
		{pythonCodeBlockIn, "python", pythonCodeBlockOut},
	}

	for _, testCase := range testCases {
		output := prepareCodeBlock(testCase.input, testCase.language)
		if testCase.output != output {
			t.Errorf("expected %s, got %s", testCase.output, output)
		}
	}
}
