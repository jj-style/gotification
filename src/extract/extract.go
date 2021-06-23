package extract

import (
	"gotification/src/util"
	"log"
	"sync"
)

var (
	extractor     Extractor
	extractorOnce sync.Once
)

type ExtractorType string

const (
	EXTRACT_NAIVE ExtractorType = "naive"
	EXTRACT_TIKA  ExtractorType = "tika"
)

func Extract() Extractor {
	extractorOnce.Do(func() {
		switch ExtractorType(util.Config.Extract.Type) {
		case EXTRACT_NAIVE:
			exts := make(map[string]bool, len(naiveAllowedExtensions))
			for _, ext := range naiveAllowedExtensions {
				exts[ext] = true
			}
			extractor = &naiveExtractor{
				allowedExtensions: exts,
			}
		case EXTRACT_TIKA:
			extractor = &tikaExtractor{}
		default:
			log.Fatalf("unknown extract type: %s", util.Config.Extract.Type)
		}
	})
	return extractor
}
