package main

import (
	"fmt"

	"github.com/GoPractice/Exercise"
)

// package intializers

// import (
// 	"log"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func Databasecon() {

// 	var err error
// 	dsn := "host=localhost user=postgres password=@kash123 dbname=DemoDatabase port=8080 sslmode=disable"
// 	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

// 	if err != nil {

// 		log.Fatal("Failed to connect to Database")
// 	}

// }

func main() {

	var arr = []string{"a", "b", "c", "a", "b"}
	var mapValue = Exercise.Frequency(arr)
	fmt.Println("Frequency of values:", mapValue)
	var map1 = map[string]int{
		"a": 1, "c": 2, "b": 3,
	}
	var newslice = Exercise.Sortedslicemap(map1)
	fmt.Printf("Sorted slice:%v\n", newslice)
	var map2 = map[string]int{
		"a": 1, "b": 2, "c": 3,
	}
	fmt.Printf("Values multiplied by 2 in map is:%v\n", Exercise.Newmap(map2))
	newSlice := Exercise.Mapslice(map2, []string{"a", "b", "d"})
	fmt.Printf("New Slice with keys in map are:%v\n", newSlice)
	fmt.Printf("New Reversed Map:%v\n", Exercise.Reversedmap(map2))

	//Exercise1

	slice := Exercise.Slice([]int{1, 2, 4, 4, 8, 3, 3, 3})
	fmt.Printf("New Slice with even number :%v\n", slice)

	slice1 := Exercise.Sortedslice([]string{"zebra", "monkey", "aardvark"})
	fmt.Printf("New sorted Slice:%v\n", slice1)

	newMap := Exercise.Alphabeticalfrequency("Hello")
	fmt.Printf("Frequency count of each character is %v\n", newMap)

	largestNo, secondLargestNo, sum, avg := Exercise.SumAvgLargest([]int{1, 2, 3, 4, 5})
	fmt.Printf("Largest No is :%d and second largest number is : %d and sum is: %d and avg is: %d\n", largestNo, secondLargestNo, sum, avg)

	slice2 := Exercise.Slicewithmaxletter([]string{"cat", "dog", "elephant", "lion"})
	fmt.Printf("Slice with more than 5 character %v\n", slice2)

	str := Exercise.Constantstr("hello")
	fmt.Printf("String with No Vowels:%s\n", str)

	slice3 := Exercise.Removeduplicate([]int{1, 2, 3, 2, 4, 5, 5})
	fmt.Printf("Slice with no duplicates :%v\n", slice3)

	check := Exercise.MatchString([]string{"cat", "dog", "elephant"}, "dog")
	fmt.Printf("String present in slice or nots: %t\n", check)

	occurence := Exercise.FirstOccurenceint([]int{1, 2, 3, 2, 4, 5}, 2)
	fmt.Printf("First Occurence of the given number at Index is:%d\n", occurence)

	totaloccurence := Exercise.TotalOccurence([]int{1, 2, 3, 2, 4, 5}, 2)
	fmt.Printf("Total Occurence of the given number at Index is:%d\n", totaloccurence)

	newSlice4 := Exercise.ReverseSlice([]int{1, 2, 3, 4, 5})
	fmt.Printf("Reversed slice is :%v\n", newSlice4)

	check1 := Exercise.CheckSorted([]int{1, 2, 3, 4, 5})
	fmt.Printf("Check the slice is sorted or not:%v\n", check1)

	stroccurence := Exercise.FirstOccurencestr([]string{"cat", "dog", "elephant"}, "dog")
	fmt.Printf("First Occurence of the given string at Index is:%d\n", stroccurence)

	slice4 := Exercise.TakestrFromletter([]string{"cat", "dog", "elephant"}, "e")
	fmt.Printf("All words start from the given letter is :%v\n", slice4)

	slice5 := Exercise.Checkpalindromestr([]string{"racecar", "level", "hello"})
	fmt.Printf("New Checked Palindrome Slice for string  is : %v\n", slice5)

	str2 := Exercise.ReversestrArray("This is a test")
	fmt.Printf("New Reversed words of String is: %s\n", str2)

	slice6 := Exercise.Prime([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Printf("New slice with prime number is: %v\n", slice6)

	slice7 := Exercise.Checkdivisibility([]int{1, 2, 3, 4, 5}, 2)
	fmt.Printf("New Slice which is divisible by given number is: %v\n", slice7)

	slice8 := Exercise.Checknumispreset([]int{1, 2, 3, 4, 5}, 3)
	fmt.Printf("New Slice which is divisible by given number is: %v\n", slice8)

	slice9 := Exercise.Slicewithminletter([]string{"cat", "dog", "elephant", "lion"}, 4)
	fmt.Printf(" Array with min letter present is:%v\n", slice9)

	str3 := string(Exercise.Reversewholestr("This is a test"))
	fmt.Printf("Reverse String is : %s\n", str3)

	slice10 := Exercise.Checkletterpresent([]string{"cat", "dog", "elephant"}, "e")
	fmt.Printf("Slice with given letter is present: %v\n", slice10)

	check2 := Exercise.Lcm([]int{2, 3, 5})
	fmt.Printf("Lcm is :%v\n", check2)

	slice11 := Exercise.Numberpalindrome([]int{1, 11, 121, 12321})
	fmt.Printf("Slice with all Palindrome is: %v\n", slice11)

	check3 := Exercise.CheckdescendingSorted([]int{5, 4, 3, 2, 1})
	fmt.Printf("Check the slice is sorted or not:%v\n", check3)

}
