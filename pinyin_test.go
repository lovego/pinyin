package pinyin_test

import (
	"fmt"

	"github.com/lovego/pinyin"
)

func ExampleInitials() {
	fmt.Println(pinyin.Initials("长大", false))
	fmt.Println(pinyin.Initials("长城abc", false))
	fmt.Println(pinyin.Initials(" A长b大°,c ", true))
	// Output:
	// ZD
	// CCABC
	//  AZBD°,C
}
