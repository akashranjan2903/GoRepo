package exercise

import (
	"math"
	"sort"
	"strings"
)

// unction that takes a slice of integers and returns a new slice with only the even numbers.
func Slice(slice []int) []int {
	length := len(slice)
	for i := 0; i < length; {
		if slice[i]%2 != 0 {
			slice = append(slice[:i], slice[i+1:]...)
			length--

		} else {
			i++
		}
	}
	return slice
}

// function that takes a slice of strings and returns a new slice with all strings sorted in alphabetical order
func Sortedslice(slice1 []string) []string {

	sort.Strings(slice1)
	return slice1
}

// function that takes a string and returns a map with the frequency count of each character in the string
func Alphabeticalfrequency(name string) map[string]int {

	newmap := map[string]int{}
	for _, value := range name {

		val, ok := newmap[string(value)]
		if ok {
			newmap[string(value)] = val + 1
		} else {
			newmap[string(value)] = 1

		}

	}
	return newmap
}

// function for sum and largest number and avg
func SumAvgLargest(slice []int) (int, int, int, int) {

	largest := slice[0]
	sum := slice[0]
	var avg int
	for i := 1; i < len(slice); i++ {
		if largest < slice[i] {
			largest = slice[i]

		}
		sum = sum + slice[i]
		avg = sum / len(slice)
	}
	// we can also sort the array first and find the largest no directly.
	sort.Ints(slice)
	secondLargestNo := slice[len(slice)-2]

	return largest, secondLargestNo, sum, avg
}

// function that takes a slice of strings and returns a new slice with all strings that have more than 5 characters
func Slicewithmaxletter(slice []string) []string {
	length := len(slice)
	for i := 0; i < length; {

		if len(slice[i]) <= 5 {
			slice = append(slice[:i], slice[i+1:]...)
			length--
		} else {
			i++
		}
	}
	return slice
}

func Constantstr(name string) string {

	arr := []string{}
	for _, value := range name {

		if string(value) != "A" && string(value) != "E" && string(value) != "I" && string(value) != "O" && string(value) != "U" && string(value) != "a" && string(value) != "e" && string(value) != "i" && string(value) != "o" && string(value) != "u" {
			arr = append(arr, string(value))

		}
	}
	name = strings.Join(arr, "")
	return name
}

func Removeduplicate(slice []int) []int {

	newMap := map[int]bool{}
	newslice := []int{}
	for _, value := range slice {

		if _, ok := newMap[value]; !ok {

			newMap[value] = true
			newslice = append(newslice, value)
		}

	}
	return newslice
}

func MatchString(slice []string, str string) bool {

	for _, value := range slice {

		if value == str {
			return true
		}
	}
	return false
}
func FirstOccurenceint(slice []int, num int) int {

	for i, value := range slice {

		if value == num {
			return i
		}
	}
	return -1
}

func TotalOccurence(slice []int, num int) int {

	counter := 0
	for _, value := range slice {

		if value == num {
			counter++
		}
	}
	return counter
}

// reversing a slice
// Another way is to do by sort function in slice
func ReverseSlice(slice []int) []int {

	start := 0
	end := len(slice) - 1
	for start <= end {

		temp := slice[start]
		slice[start] = slice[end]
		slice[end] = temp

		start++
		end--
	}
	return slice
}

func CheckSorted(slice []int) bool {

	for i := 0; i < len(slice)-1; i++ {

		if slice[i+1] < slice[i] {
			return false
		}
	}
	return true
}

func FirstOccurencestr(slice []string, name string) int {

	for i, value := range slice {

		if value == name {
			return i
		}
	}
	return -1
}

func TakestrFromletter(slice []string, str string) []string {

	newSlice := []string{}
	for _, value := range slice {

		if strings.HasPrefix(value, str) {
			newSlice = append(newSlice, value)
		}
	}
	return newSlice
}

func Checkpalindromestr(slice []string) []string {

	newslice := []string{}
	for _, value := range slice {

		chars := []rune(value)
		length := len(chars)

		newStr := []string{}
		for i := length - 1; i >= 0; i-- {
			newStr = append(newStr, string(chars[i]))
		}
		str := strings.Join(newStr, "")
		if str == value {
			newslice = append(newslice, value)
		}

	}
	return newslice
}

// reverse a string
func Reverserunestr(str []rune) []rune {

	start := 0
	end := len(str) - 1
	for start <= end {

		temp := str[start]
		str[start] = str[end]
		str[end] = temp

		start++
		end--
	}
	return str
}

// reverse a string with start and end points
func Reversestr(str []rune, start, end int) []rune {

	for start < end {

		temp := str[start]
		str[start] = str[end]
		str[end] = temp

		start++
		end--
	}
	return str
}

// reverse words of the string
func ReversestrArray(str string) string {

	runearray := []rune(str)
	newarray := Reverserunestr(runearray)
	length := len(runearray)
	for i := 0; i < length; {

		for i < length && string(newarray[i]) == " " {
			i++
		}
		if i == length {
			break
		}

		start := i
		for i < length && string(newarray[i]) != " " {
			i++
		}
		end := i - 1
		newarray = Reversestr(newarray, start, end)

	}
	return string(newarray)

}

func Prime(slice []int) []int {

	length := len(slice)
	for index := 0; index < length; {

		var i int
		if slice[index] == 1 {

			slice = append(slice[:index], slice[index+1:]...)
			length--
			index++

		}
		for i = 2; i <= int(math.Sqrt(float64(slice[index]))); i++ {
			if slice[index]%i == 0 {
				length--
				slice = append(slice[:index], slice[index+1:]...)
				break
			}
		}
		if index < length && i > int(math.Sqrt(float64(slice[index]))) {
			index++
		}
	}
	return slice
}

// func alternate

func Checkdivisibility(slice []int, num int) []int {

	length := len(slice)
	for index := 0; index < length; {

		if slice[index]%num != 0 {
			slice = append(slice[:index], slice[index+1:]...)
			length--
		} else {
			index++
		}

	}
	return slice
}

func Checknumispreset(slice []int, num int) bool {

	for _, value := range slice {

		if value == num {
			return true
		}

	}
	return false
}

func Slicewithminletter(slice []string, num int) []string {
	length := len(slice)
	for i := 0; i < length; {

		if len(slice[i]) < num {
			slice = append(slice[:i], slice[i+1:]...)
			length--
		} else {
			i++
		}
	}
	return slice
}

func Checkletterpresent(slice []string, letter string) []string {

	length := len(slice)
	for i := 0; i < length; {

		chars := []rune(slice[i])
		flag := false
		for _, v := range chars {

			if string(v) == letter {
				flag = true
			}
		}
		if !flag {
			slice = append(slice[:i], slice[i+1:]...)
			length--
		} else {
			i++
		}

	}
	return slice
}

func Numberpalindrome(slice []int) []int {

	length := len(slice)
	for index := 0; index < length; {

		reversenum := 0
		a := slice[index]
		if a < 10 {
			slice = append(slice[:index], slice[index+1:]...)
			length--
			continue
		}
		for a != 0 {
			remainder := a % 10
			reversenum = reversenum*10 + remainder
			a = a / 10
		}

		if reversenum != slice[index] {

			slice = append(slice[:index], slice[index+1:]...)
		} else {
			index++
		}

	}
	return slice
}

func CheckdescendingSorted(slice []int) bool {

	for i := 0; i < len(slice)-1; i++ {

		if slice[i+1] > slice[i] {
			return false
		}
	}
	return true
}
func Lcm(slice []int) int {

	smallest := slice[0]

	for i := 0; i < len(slice)-1; i++ {

		if smallest > slice[i] {
			smallest = slice[i]

		}

	}
	/* Use of For Loop as a While Loop*/
	for {
		flag := false
		for i, value := range slice {
			if smallest%value != 0 {
				break
			}
			if i == len(slice)-1 {
				flag = true
			}
		}
		if !flag {
			smallest++
		} else {
			break
		}

	}
	return smallest // Return without any value
}
