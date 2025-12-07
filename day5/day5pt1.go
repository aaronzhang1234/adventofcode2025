package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func mainpt1(){
	file, _ := os.Open("DB.txt")
	scanner:= bufio.NewScanner(file)
	var db [][]int
	for scanner.Scan(){
	    memes := strings.Split(scanner.Text(), "-")
		min, _ := strconv.Atoi(memes[0])
		max, _ := strconv.Atoi(memes[1])
		db = append(db,[]int{min, max})
	}	

	stockFile, _ := os.Open("stocklist.txt")
	stockScanner:= bufio.NewScanner(stockFile)
	numFresh :=0
	for stockScanner.Scan(){
		stockListNum,_ := strconv.Atoi(stockScanner.Text())
		for i:=0; i<len(db); i++{
			if stockListNum>= db[i][0] && stockListNum<= db[i][1]{
				numFresh = numFresh +1
				break
			}
		}
	}
	fmt.Println(numFresh)

}