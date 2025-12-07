package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func mainpt1(){
	file, _ := os.Open("input.txt")
	scanner:= bufio.NewScanner(file)
	var mathMatrix [][]string
	largestSize := 0 
	for scanner.Scan(){
		lineSlice := strings.Split(scanner.Text(),"")
		mathMatrix = append(mathMatrix, lineSlice)
		if len(lineSlice) > largestSize{
			largestSize = len(lineSlice)
		}
	}
	lastBorder := 0
	total := 0
	for i:=0;i<len(mathMatrix[0]);i++{
		if isBorder(i, mathMatrix){	
			total = total + createNumMatrixpt1(lastBorder, i, mathMatrix)
			lastBorder = i
		}
	}
	total = total + createNumMatrixpt1(lastBorder, largestSize, mathMatrix)
	fmt.Println(total)
}
func isBorderpt1(i int, matrix [][]string) bool{
	isBorder := true
	for x:=0; x<len(matrix);x++{
		if matrix[x][i] != " "{
			isBorder = false
			break
		}
	}
	return isBorder
}
func createNumMatrixpt1(start, end int, matrix [][]string) int {
	endMatrix := []string{}
	for x:=0 ; x<len(matrix); x++{
		array := []string{}
		for i:=start;i<end;i++{
			array = append(array, matrix[x][i])
		}
		endMatrix = append(endMatrix, strings.ReplaceAll(strings.Join(array,""), " ", ""))
	}
	total := 1
	for x:=0; x<len(endMatrix)-1; x++{
		num,_ := strconv.Atoi(endMatrix[x])
		if endMatrix[len(endMatrix)-1] == "*"{
			total = total * num 
		}
		if endMatrix[len(endMatrix)-1] == "+"{
			total = total + num 
		}
	}
	if endMatrix[len(endMatrix)-1] == "+"{
		return total-1
	}
	return total
}