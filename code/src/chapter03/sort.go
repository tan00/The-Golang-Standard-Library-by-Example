package main

import (
	"fmt"
	"sort"
	"strconv"
)

//内置类型排序
func TestSort() {
	numSlice1 := []int{3, 34, 6, 56, 7, 787, 68, 463}
	numSlice2 := []int{35, 4, 254, 21, 45, 46, 6, 34, 535, 352, 1, 214, 36}

	intSlice := sort.IntSlice(numSlice1)
	intSlice.Sort()
	fmt.Println(intSlice)

	//func Ints(a []int) { Sort(IntSlice(a)) }
	sort.Ints(numSlice2)
	fmt.Println(numSlice2)

	//func IntsAreSorted(a []int) bool { return IsSorted(IntSlice(a)) }
	fmt.Println(sort.IntsAreSorted(numSlice2))
}

//自定义类型排序 需要实现 sort.Interface 接口
// type Interface interface {
// 	// Len is the number of elements in the collection.
// 	Len() int
// 	// Less reports whether the element with
// 	// index i should sort before the element with index j.
// 	Less(i, j int) bool
// 	// Swap swaps the elements with indexes i and j.
// 	Swap(i, j int)
// }

type Person struct {
	age  int
	sex  int
	name string
}
type Persons []Person

func (persons Persons) Len() int {
	return len(persons)
}
func (persons Persons) Less(i, j int) bool {
	return persons[i].age < persons[j].age
}

func (persons Persons) Swap(i, j int) {
	persons[i], persons[j] = persons[j], persons[i]
}

func (persons Persons) String() string {
	var str string
	for _, v := range persons {
		//str += fmt.Sprintf("%s, age=%d , sex=%d \n", v.name, v.age, v.sex)
		str += "Person, name:" + v.name + "\tage:" + strconv.Itoa(v.age) + "\tsex:" + strconv.Itoa(v.sex) + "\n"
	}
	return str
}

func TestSort2() {
	var (
		persons = Persons{
			{18, 1, "James"},
			{5, 1, "Sfsaf"},
			{23, 0, "WEAFs"},
			{76, 1, "dvweg"},
			{2, 0, "Wefwrqw"},
			{36, 1, "Werfscd"},
		}
	)

	sort.Sort(persons)

	//下面三行都会调用 Persons.String()
	fmt.Println(persons)
	fmt.Printf("%v", persons)
	fmt.Printf("%s", persons)
}
