package extract

import (
	"os"
)

type Extractor interface {
	ExtractFile(reader *os.File) (string, error)
}
