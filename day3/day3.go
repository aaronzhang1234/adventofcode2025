package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func main(){
	file, _ := os.Open("input.txt")
	scanner:= bufio.NewScanner(file)
	total := 0
	for scanner.Scan(){
		total = total + findHighestJolt(scanner.Text())
	}
	fmt.Println(total)
}
func findHighestJolt(bank string) int{
	max := 0
	for i:=0;i<len(bank)-1;i++{
		for y:=i+1; y<len(bank); y++{
			num, _ := strconv.Atoi(string(bank[i]) + string(bank[y]))
			if num > max{
				max = num
			}
		}	
	}
	return max
}