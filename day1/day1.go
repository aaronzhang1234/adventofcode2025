package main

import(
	"bufio"
	"fmt"
	"os"
	"strconv"
	"math"
)


func main(){
	file, _ := os.Open("input.txt")
	scanner:= bufio.NewScanner(file)

	loc := 50
	numZeros := 0
	zerosHit := 0	

	for scanner.Scan(){
		var hitZeros int
		loc, hitZeros = moveLock(loc, scanner.Text())
		if loc == 0{
			numZeros = numZeros + 1
		}
		zerosHit = zerosHit + hitZeros
	}
	fmt.Println(numZeros)
	fmt.Println(zerosHit)
}

func moveLock(loc int, turn string) (int, int){
	direction := string(turn[0])
	notches, _ := strconv.Atoi(turn[1:])

	fullTurns := int(math.Floor(float64(notches) / 100))
	remainder := notches % 100

	if remainder == 0{
		return loc, fullTurns
	}
	
	loc = convertL(loc, direction)

	end := loc + remainder
	if end == 0 || end == 100{
		return 0, fullTurns+1
	}
	if end > 100{
		return convertL(end - 100 ,direction), fullTurns+ 1
	}
	return convertL(end, direction), fullTurns
}

func convertL(num int, direction string) int{
	if direction=="L"{
		return (100 - num) % 100
	}
	return num
}