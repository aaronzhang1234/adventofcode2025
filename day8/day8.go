package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
	"sort"
)

type Junction struct{
	x float64 
	y float64
	z float64
	circuit int
}

type kv struct {
	Key   string
	Value float64
}

func main(){
	file, _ := os.Open("input.txt")
	scanner:= bufio.NewScanner(file)
	var junctions []*Junction
	for scanner.Scan(){
		text := scanner.Text()
		locs := strings.Split(text, ",")
		x,_ := strconv.ParseFloat(locs[0], 64)
		y,_ := strconv.ParseFloat(locs[1], 64)
		z,_ := strconv.ParseFloat(locs[2], 64)
		junction := Junction{x:x, y:y,z:z, circuit: -1}
		junctions = append(junctions, &junction)
	}
	distances := make(map[string]float64)
	circuits := make(map[int][]*Junction)
    var ss []kv
	for i:=0;i<len(junctions)-1;i++{
		for x:=i+1; x<len(junctions);x++{
			key := fmt.Sprintf("%d-%d", i,x)
			value := compareJunctionLengths(junctions[i], junctions[x])
			distances[key] = value 
			ss = append(ss, kv{key, value})
		}
	}
	sort.Slice(ss, func(i, j int) bool {
        return ss[i].Value < ss[j].Value
    })

	for num, kv := range ss {
		if num==1000{
			break
		}
		fmt.Println("Starting new connection")
		juncLos := strings.Split(kv.Key, "-")
		firstJuncLoc,_ := strconv.Atoi(juncLos[0])
		secondJuncLoc,_ := strconv.Atoi(juncLos[1])
		firstJunc, secondJunc := junctions[firstJuncLoc], junctions[secondJuncLoc]
		fmt.Println(firstJunc)
		fmt.Println(secondJunc)
		fmt.Println(circuits)
		if firstJunc.circuit==-1 &&  secondJunc.circuit==-1{
			circuits[len(circuits)+1] = []*Junction{firstJunc, secondJunc}
			firstJunc.circuit, secondJunc.circuit = len(circuits), len(circuits)
		}else if firstJunc.circuit==-1 && secondJunc.circuit!=-1{
			firstJunc.circuit = secondJunc.circuit
			circuits[secondJunc.circuit] = append(circuits[secondJunc.circuit], firstJunc)
		}else if secondJunc.circuit==-1 && firstJunc.circuit!=-1{
			secondJunc.circuit = firstJunc.circuit
			circuits[firstJunc.circuit] = append(circuits[firstJunc.circuit], secondJunc)
		}else if firstJunc.circuit !=-1 && secondJunc.circuit!=-1 && firstJunc.circuit!=secondJunc.circuit{
			secondCircuitSlice := circuits[secondJunc.circuit] 
			toRemoveCircuit := secondJunc.circuit
			for _,o:= range circuits[secondJunc.circuit]{
				o.circuit = firstJunc.circuit
				fmt.Println(o)
			}
			circuits[firstJunc.circuit] = append(circuits[firstJunc.circuit], secondCircuitSlice...)
			circuits[toRemoveCircuit] = nil
		}
		fmt.Println(firstJunc)
		fmt.Println(secondJunc)
		fmt.Println(circuits)
    }
	var lengths []int
	for _,v := range circuits{
		lengths = append(lengths, len(v))
	}
	sort.Ints(lengths[:])
	fmt.Println(lengths)


}

func compareJunctionLengths(junc1, junc2 *Junction) float64 {
	return math.Sqrt(math.Pow(junc1.x-junc2.x,2) + math.Pow(junc1.y-junc2.y,2) + math.Pow(junc1.z-junc2.z,2))
}