package array

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTowSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	assert.Equal(t, []int{0, 1}, towSum(nums, 9))

	nums = []int{3, 2, 4}
	assert.Equal(t, []int{1, 2}, towSum(nums, 6))

	nums = []int{3, 3}
	assert.Equal(t, []int{0, 1}, towSum(nums, 6))
}

func TestTreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum2(nums)
	fmt.Println(res)
	//[[-1,-1,2],[-1,0,1]]
}

type Student struct {
	name string
	age  int
}

func TestRangeError(t *testing.T) {
	m := make(map[string]*Student)
	students := []Student{
		{name: "小王子", age: 18},
		{name: "大王子", age: 20},
		{name: "大王八", age: 1800},
	}
	//遍历student地址复用
	for _, student := range students {
		m[student.name] = &student
	}
	for k, v := range m {
		fmt.Println(k, "=>", v)
	}
}

func TestRangeNormal(t *testing.T) {
	m := make(map[string]*Student)
	students := []Student{
		{name: "小王子", age: 18},
		{name: "大王子", age: 20},
		{name: "大王八", age: 1800},
	}
	for i := 0; i < len(students); i++ {
		m[students[i].name] = &students[i]
	}
	for k, v := range m {
		fmt.Println(k, "=>", v)
	}
}
