package classifier

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	sent := "je veux un véhicule de remplacement Suzuki ouais des tests le matin vers 9 h"
	cls := NewIntentClassifier("SOC")
	fmt.Println(cls.GetIntent(sent, false))
	fmt.Println(time.Since(start))
}
