package Exercise

import (
	"sort"
)

// function that takes a slice of strings and returns a map with the frequency of each string in the slice
func Frequency(slice []string) map[string]int {

	var temp = map[string]int{}
	for _, v := range slice {

		value, ok := temp[v]
		if ok {
			temp[v] = value + 1
		} else {
			temp[v] = 1
		}
	}
	return temp
}

// function that takes a map and returns a slice of the keys in the map, sorted in alphabetical order
func Sortedslicemap(temp map[string]int) []string {

	var slice = []string{}
	for key := range temp {
		slice = append(slice, key)
	}
	sort.Strings(slice)
	return slice
}

// function that takes a map and returns a new map with all the values multiplied by 2.
func Newmap(temp map[string]int) map[string]int {
	for key, value := range temp {

		temp[key] = value * 2
	}
	return temp
}

// function that takes a map and a slice of strings, and returns a new slice with all the strings from the input slice that are keys in the map
func Mapslice(newMap map[string]int, slice []string) []string {

	var newslice = []string{}
	for _, value := range slice {

		for key := range newMap {

			if key == value {
				newslice = append(newslice, key)
				continue
			}
		}
	}
	return newslice
}

// function that takes a map and returns a new map with all the keys and values reversed
func Reversedmap(temp map[string]int) map[int]string {

	var newMap = map[int]string{}
	for key, value := range temp {

		newMap[value] = key
	}
	return newMap
}

// func ToCamelCase(s string) string {
// 	var slice = []string{}
// 	if strings.Contains(s, "_") {
// 		slice = strings.Split(s, "_")
// 	}
// 	if strings.Contains(s, "-") {
// 		slice = strings.Split(s, "-")
// 	}

// 	chars := []rune(slice[0])
// 	var str string
// 	println(slice[0])

// 	println(unicode.IsUpper(chars[0]))
// 	if unicode.IsUpper(chars[0]) {
// 		for _, value := range slice {
// 			strings.Title(value)
// 		}
// 		str = strings.Join(slice, "")
// 		return str
// 	} else {
// 		for i := 1; i < len(slice); i++ {
// 			value := strings.Title(slice[i])
// 			slice[i] = value

// 		}
// 		str = strings.Join(slice, "")
// 		return str

// 	}
// }

// func SpinWords(str string) string {

// 	slice := strings.Split(str, " ")
// 	for i, value := range slice {
// 		if len(slice[i]) >= 5 {

// 			chars := []rune(value)
// 			start := 0
// 			end := len(chars) - 1
// 			for start < end {

// 				chars[start], chars[end] = chars[end], chars[start]

// 				start++
// 				end--
// 			}
// 			slice[i] = string(chars)
// 		}
// 		str = strings.Join(slice, " ")
// 	}
// 	return str

// }

// func Accum(s string) string {

// 	chars := []rune(s)
// 	var slice = []string{}
// 	for key, value := range chars {
// 		temp := string(value)
// 		word := strings.Repeat(temp, key)
// 		upperCaseletter := strings.ToUpper(string(value))
// 		word = upperCaseletter + word
// 		slice = append(slice, word)
// 	}
// 	s = strings.Join(slice, "-")
// 	return s

// }
