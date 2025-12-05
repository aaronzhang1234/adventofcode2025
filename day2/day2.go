package main

import(
	"fmt"
	"strings"
	"strconv"
)


func main(){
	//input := "222220-222224"
	input := "503950-597501,73731-100184,79705998-79873916,2927-3723,35155-50130,52-82,1139-1671,4338572-4506716,1991-2782,1314489-1387708,8810810-8984381,762581-829383,214957-358445,9947038-10058264,4848455367-4848568745,615004-637022,5827946-5911222,840544-1026063,19-46,372804-419902,486-681,815-1117,3928-5400,28219352-28336512,6200009-6404247,174-261,151131150-151188124,19323-26217,429923-458519,5151467682-5151580012,9354640427-9354772901,262-475,100251-151187,5407-9794,8484808500-8484902312,86-129,2-18"
	ranges := strings.Split(input, ",")
	total:=0
	for _, value := range ranges{
		min,_ := strconv.Atoi(strings.Split(value, "-")[0])
		max,_ := strconv.Atoi(strings.Split(value, "-")[1])
		for i := min; i<=max; i++{
			stringNum := strconv.Itoa(i)
			for x:=1; x<=len(stringNum)/2; x++{
				if (len(stringNum)) % x == 0{			
					isValid := true
					numsIn := len(stringNum)/x
					numer := 0
					for y:=0; y<numsIn-1; y++{
						if stringNum[numer:numer+x] != stringNum[numer+x: numer+(x*2)]{					
							isValid = false
							break
						}
						numer = numer + x
					}
					if isValid{
						stringInt,_ := strconv.Atoi(stringNum)
						total = total + stringInt
					}
				}
			}
		}
	}
	fmt.Println(total)
}