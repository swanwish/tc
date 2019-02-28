package tc

import (
	"fmt"
	"testing"
)

func TestText(t *testing.T) {
	fmt.Printf("%s\n", Text("Red", TextColorRed))
	fmt.Printf("%s\n", Text("Red", TextColorRed, TextStyleBold))
}

func TestNewColorTemplate(t *testing.T) {
	redBoldText := NewColorTemplate(TextColorRed, TextStyleBold)
	fmt.Printf("%s\n", redBoldText.Text("Hello world"))
	fmt.Printf("%s\n", redBoldText.Text(100))
}
