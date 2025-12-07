package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"slices"
)


func main(){
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
			total = total + createNumMatrix(lastBorder, i, mathMatrix)
			lastBorder = i+1
		}
	}
	total = total + createNumMatrix(lastBorder, largestSize, mathMatrix)
	fmt.Println(total)
}
func isBorder(i int, matrix [][]string) bool{
	isBorder := true
	for x:=0; x<len(matrix);x++{
		if matrix[x][i] != " "{
			isBorder = false
			break
		}
	}
	return isBorder
}
func createNumMatrix(start, end int, matrix [][]string) int {
	var oufMatrix [][]string
	for x:=0 ; x<len(matrix); x++{
		array := []string{}
		for i:=start;i<end;i++{
			array = append(array, matrix[x][i])
		}
		oufMatrix = append(oufMatrix, array)
	}
	cephMatrix := switchMatrix(oufMatrix)
	total:=1
	for x:=0; x<len(cephMatrix); x++{
		if slices.Contains(oufMatrix[len(oufMatrix)-1], "*"){
			total = total * cephMatrix[x]
		}
		if slices.Contains(oufMatrix[len(oufMatrix)-1], "+"){
			total = total + cephMatrix[x]
		}
	}
	if slices.Contains(oufMatrix[len(oufMatrix)-1], "+"){
		return total-1
	}
	return total
}
func switchMatrix(matrix [][]string) []int{
	var endMatrix []int
	for x:=len(matrix[0])-1;x>=0;x--{
		array := []string{}
		for y:=0; y<len(matrix)-1; y++{
			array = append(array, (matrix[y][x]))
		}
		num,_ := strconv.Atoi(strings.ReplaceAll(strings.Join(array,""), " ", ""))
		endMatrix = append(endMatrix, num)
	}
	return endMatrix
}