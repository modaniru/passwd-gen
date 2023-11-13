package main

import (
	"fmt"
	"log"
	"math/rand"
)

var lettersLow = "abcdefghijklmnopqrstuvwxyz"
var lettersUp = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var numbers = "1234567890"
var symbols = "@#$%&*!~"

func main(){
	var length int
	fmt.Print("Password len: ")
	_, err := fmt.Scanln(&length)
	if err != nil{
		log.Fatal(err.Error())
	}

	alphabet := lettersLow
	var op int
	fmt.Print("include uppercase? (0 - NO, 1 - YES)")
	_, err = fmt.Scanln(&op)
	if err != nil{
		log.Fatal(err.Error())
	}

	if op != 0{
		alphabet += lettersUp
	}

	fmt.Print("include numbers? (0 - NO, 1 - YES)")
	_, err = fmt.Scanln(&op)
	if err != nil{
		log.Fatal(err.Error())
	}

	if op != 0{
		alphabet += numbers
	}

	fmt.Print("include special symbols? (0 - NO, 1 - YES)")
	_, err = fmt.Scanln(&op)
	if err != nil{
		log.Fatal(err.Error())
	}

	if op != 0{
		alphabet += symbols
	}

	res := make([]byte, length)
	for i := 0; i < length; i++{
		index := rand.Intn(len(alphabet))
		res[i] = alphabet[index]
	}
	fmt.Println(string(res))
}