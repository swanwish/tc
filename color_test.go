package terminalcolor

import (
	"fmt"
	"testing"
)

func TestColorText(t *testing.T) {
	fmt.Printf("%s\n", ColorText("Red", TextColorRed))
}
