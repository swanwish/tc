package tc

import (
	"fmt"
	"testing"
)

func TestText(t *testing.T) {
	fmt.Printf("%s\n", Text("Red", TextColorRed))
	fmt.Printf("%s\n", Text("Red", TextColorRed, TextStyleBold))
}
