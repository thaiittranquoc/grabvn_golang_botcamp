package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
   scanner := bufio.NewScanner(os.Stdin)
   fmt.Print(">")
   for scanner.Scan(){
   	    text := scanner.Text()
   	    temp := eval(text)
   	    fmt.Println(temp)
   	    fmt.Print(">")
    }
}