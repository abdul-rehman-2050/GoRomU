package main;

import "fmt"
import "os"
import "strings"

func main(){
	
	stringsToConvert := os.Args[1:]
	fmt.Println("Hello world!")
	
	if(len(stringsToConvert)<1){
		fmt.Println("Please Provide valid Arguments")
		return;
		
	}
	
	
	
	// Accept Command line arguments and convert them to tokenized strings
	
	for _,v := range stringsToConvert {
		parts := strings.Split(v, " ")
		for _,v2 := range parts{
		fmt.Println(v2)
		}
    }
	
	
	
	
	
}