package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	file, _ := os.Open("input.txt")
	scanner:= bufio.NewScanner(file)
	var beamLocations []string
	split :=0
	for scanner.Scan(){
		text := scanner.Text()
		if beamLocations == nil{
			for i:=0; i<len(text); i++{
				if string(text[i])=="S"{
					beamLocations = append(beamLocations, "|")
				}else{
					beamLocations = append(beamLocations, ".")
				}	 
			}
		}else{
			for i:=0; i<len(text); i++{
				if string(text[i]) == "^" && beamLocations[i] == "|"{
					beamLocations[i-1] = "|"
					beamLocations[i+1] = "|"
					beamLocations[i] = "."
					split = split + 1
				}
			}
		}
		fmt.Println(beamLocations)
	}	
	fmt.Println(split)
}
