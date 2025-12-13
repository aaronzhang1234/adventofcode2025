package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	file, _ := os.Open("input.txt")
	scanner:= bufio.NewScanner(file)
	var beamLocations []int
	var nextBeamLocations []int
	for scanner.Scan(){
		text := scanner.Text()
		if beamLocations == nil{
			for i:=0; i<len(text); i++{
				if string(text[i])=="S"{
					beamLocations = append(beamLocations, 1)
				}else{
					beamLocations = append(beamLocations, 0)
				}	 
			}
		}else{
			nextBeamLocations = beamLocations
			for i:=0; i<len(text); i++{
				if string(text[i]) == "^" && beamLocations[i] != 0{
					nextBeamLocations[i-1] =+ nextBeamLocations[i-1]+beamLocations[i]
					nextBeamLocations[i+1] =+ nextBeamLocations[i+1]+beamLocations[i]
					nextBeamLocations[i] = 0
				}else{
					nextBeamLocations[i] = beamLocations[i]
				}
			}
			fmt.Println(nextBeamLocations)
			beamLocations = nextBeamLocations
		}
		fmt.Println(beamLocations)
	}	
	total := 0
	for x:=0; x<len(beamLocations); x++{
		total = total + beamLocations[x]
	}
	fmt.Println(total)
}
