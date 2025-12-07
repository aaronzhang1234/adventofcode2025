package main

import (
	"cmp"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"slices"
)

func main(){
	file, _ := os.Open("DB.txt")
	scanner:= bufio.NewScanner(file)
	var db [][]int
	for scanner.Scan(){
	    memes := strings.Split(scanner.Text(), "-")
		min, _ := strconv.Atoi(memes[0])
		max, _ := strconv.Atoi(memes[1])
		db = append(db,[]int{min, max})
	}	
	slices.SortFunc(db, CompareRanges)
	for i :=0; i<len(db)-1;i++{
		a,b := db[i], db[i+1]
		if (a[0] == b[0]) || (b[0]>=a[0] && b[0]<=a[1]){
			db[i] =	combineRanges(a,b)
			db = append(db[:i+1], db[i+2:]...)
			i = i-1
		}
	}
	totalFresh := 0
	for i :=0; i<len(db);i++{
		totalFresh = totalFresh + db[i][1]-db[i][0] + 1
	}
	fmt.Println(totalFresh)
}
func CompareRanges(a,b[]int) int{
	if a[0] == b[0]{
		return cmp.Compare(a[1], b[1])
	}
	return cmp.Compare(a[0], b[0])
}

func combineRanges(a,b []int) []int{
	if a[0]==b[0] && a[1] == b[1]{ //same range
		return []int{a[0], a[1]}
	}
	if b[1] > a[1]{
		return []int{a[0],b[1]}
	}
	return []int{a[0], a[1]}
}