package extract

import (
	"os"
)

type tikaExtractor struct{}

func (t *tikaExtractor) ExtractFile(file *os.File) (string, error) {
	panic("TODO - implement me!")
}
