package Games

import (
	"fmt"
	"math/rand"
	"time"
)

// Guessing game
func Systemnum() int {

	rand.Seed(time.Now().Unix())
	var secretnum int = rand.Intn(9)
	return secretnum
}
func CheckEnteredNum(num int, secretnum int, highestVal int) bool {
	var flag bool = false
	if num > highestVal || num < 0 {
		fmt.Println("enter the 1 cond")
		fmt.Println("You Entered the Invalid Number")
		fmt.Println("You Missed One Chance")
	} else if num > secretnum {
		fmt.Println("enter the 2 cond")
		fmt.Println("You Enter the higher no than secret no")
	} else if num < secretnum {
		fmt.Println("enter the 3 cond")
		fmt.Println("You Enter the lesser no than secret no")
	}
	if num == secretnum {
		fmt.Println("The Secret Number is", secretnum)
		fmt.Println("Congratulation!You Guessed the Number")
		fmt.Println("Game End !")
		flag = true
	}
	return flag
}
func Guess() {
	var value int = 9
	var secretnum = Systemnum()
	for i := 0; i <= value; i++ {

		if i == 9 {
			fmt.Println("Game End ! , You Missed All Chances")
		}
		var num int
		fmt.Println("Guess the Number")
		fmt.Println("Enter your", i+1, "Guess Between 0 to 9")
		_, err := fmt.Scanf("%v\n", &num)
		if err != nil {
			fmt.Println("ERROR occured during Input")
			return
		}
		var flag bool = CheckEnteredNum(num, secretnum, value)

		if !flag {
			continue
		} else {
			break
		}

	}

}
