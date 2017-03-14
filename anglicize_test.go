package anglicize

import (
	"fmt"
)

func ExampleKeywords() {
	result := Anglicize("foo ÂÇİĞÖŞÜÑ bar âçığöşüñ qux")

	fmt.Println(result)
	// Output: foo ACIGOSUN bar acigosun qux
}
