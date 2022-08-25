package classifier

import (
	"fmt"
	"github.com/mx79/go-nlp/extractor"
	"time"
)

func main() {
	start := time.Now()
	sent := "je veux un véhicule de remplacement Suzuki ouais des tests le matin vers 9 h"
	// cls := classifier.NewIntentClassifier("SOC")
	// fmt.Println(cls.Predict(sent))
	ext := extractor.NewEntityExtractor("test")
	fmt.Println(ext.GetEntity(sent))
	fmt.Println(time.Since(start))
}
