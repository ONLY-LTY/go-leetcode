package string

import (
	"fmt"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	res := findAnagrams("aa", "bb")
	fmt.Println(res)
}
