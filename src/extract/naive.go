package extract

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type naiveExtractor struct {
	allowedExtensions map[string]bool
}

var naiveAllowedExtensions = []string{"txt", "md", "csv", "c", "cpp", "go", "py", "rs", "conf", "ini", "toml", "yaml", "yml"}

func (n *naiveExtractor) ExtractFile(file *os.File) (string, error) {
	extension := filepath.Ext(file.Name())[1:]
	if _, allowed := n.allowedExtensions[extension]; allowed {
		content, err := ioutil.ReadAll(file)
		if err != nil {
			return "", fmt.Errorf("error reading file %s: %s", file.Name(), err.Error())
		}
		return string(content), nil
	}
	return "", fmt.Errorf("unsupported file extension: %s", extension)
}
