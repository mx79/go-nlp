package extractor

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	sent := "Une phrase de test"
	ext := NewEntityExtractor("test")
	fmt.Println(ext.GetEntity(sent))
	fmt.Println(time.Since(start))
}
