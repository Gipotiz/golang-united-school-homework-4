package string_sum

import (
	"fmt"
	"testing"
)

func TestStringSum2(t *testing.T) {
	v, err := StringSum(" -   21-2")
	//StringSum("35+5")
	fmt.Println(v, err)
}
