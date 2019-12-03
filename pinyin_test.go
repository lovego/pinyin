package pinyin_test

import (
	"fmt"

	"github.com/lovego/pinyin"
)

func ExampleInitials() {
	fmt.Println(pinyin.Initials("长大"))
	fmt.Println(pinyin.Initials("长城abc"))
	fmt.Println(pinyin.Initials(" A长b大,c "))
	// Output:
	// ZD
	// ZCABC
	// AZBD,C
}
