package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func mainpt1(){
	file, _ := os.Open("input.txt")
	scanner:= bufio.NewScanner(file)
	var forkMatrix [][]string
	for scanner.Scan(){
		lineSlice := strings.Split(scanner.Text(),"")
		forkMatrix = append(forkMatrix, lineSlice)
	}
	fmt.Println(forkMatrix)
	cleanable := 0
	for i:=0;i<len(forkMatrix);i++{
		for x:=0; x<len(forkMatrix[i]); x++{
			if forkMatrix[i][x] == "@"{
				if isCleanablept1(forkMatrix, i ,x){
					cleanable = cleanable +1	
				}
			}
		}
	}
	fmt.Println(cleanable)
}
func isCleanablept1(matrix [][]string, i int, x int) bool {
	validSpots:= make(map[string][]int)

	validSpots["upper-left"] = []int{i-1, x-1}
	validSpots["upper-upper"] = []int{i-1, x}
	validSpots["upper-right"] = []int{i-1, x+1}
	validSpots["left"] = []int{i, x-1}
	validSpots["right"] = []int{i, x+1}
	validSpots["lower-left"] = []int{i+1, x-1}
	validSpots["lower-lower"] = []int{i+1, x}
	validSpots["lower-right"] = []int{i+1, x+1}

	if i-1 <0{ // top row
		delete(validSpots, "upper-left")
		delete(validSpots, "upper-upper")
		delete(validSpots, "upper-right")
	}
	if x -1 <0 { // left-most column
		delete(validSpots, "upper-left")
		delete(validSpots, "left")
		delete(validSpots, "lower-left")
	}
	if x+1 > len(matrix[i])-1{ //right-most column
		delete(validSpots, "upper-right")
		delete(validSpots, "right")
		delete(validSpots, "lower-right")	
	}
	if i+1 > len(matrix)-1{ // bottom row	}
		delete(validSpots, "lower-left")
		delete(validSpots, "lower-lower")
		delete(validSpots, "lower-right")
	}
	numRolls := 0
	for _,value := range validSpots{
		if matrix[value[0]][value[1]] == "@"{
			numRolls = numRolls+1
		}
	}
	return numRolls<=3
}