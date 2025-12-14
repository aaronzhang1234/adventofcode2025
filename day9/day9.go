package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sort"
)

type RedTile struct{
	x int
	y int
}

func main(){
	file, _ := os.Open("input.txt")
	scanner:= bufio.NewScanner(file)	
	var tiles []RedTile
	for scanner.Scan(){
		text := scanner.Text()
		xString, yString := strings.Split(text, ",")[0], strings.Split(text,",")[1]
		x,_ := strconv.Atoi(xString)
		y,_ := strconv.Atoi(yString)
		tiles = append(tiles, RedTile{x:x, y:y})
	}
	var areas []int
	for x:=0;x<len(tiles)-1;x++{
		for y:=x+1;y<len(tiles);y++{
			areas = append(areas, getAreas(tiles[x], tiles[y]))
		}
	}
	sort.Ints(areas[:])
	fmt.Println(areas[len(areas)-1])
}

func getAreas(tile1, tile2 RedTile) int{
	return (int(math.Abs(float64(tile1.x-tile2.x)))+1)*
			(int(math.Abs(float64(tile1.y-tile2.y)))+1)
}