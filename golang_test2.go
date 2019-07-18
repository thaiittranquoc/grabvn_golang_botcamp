package main

import (
	"fmt"
	"strings"
	"strconv"
)

func eval(text string) string{
	list_text := strings.Split(text, " ")
	if len(list_text) > 3 {
		error := "Invalid"
		return error 
	}
	operator := list_text[1]
	f0, f2, err := get_val(list_text[0], list_text[2])
	if err != nil {
		error := "Invalid"
		return error
	}
	var result float64
	switch operator{
	    case "+":
	    	result = f0 + f2
	    case "-":
	    	result = f0 - f2
	    case "*":
	    	result = f0 * f2
	    case "/":
	    	if f2 == 0 {
	    		error := "Invalid"
	    		return error
	    	}
	    	result = f0 / f2
	    default:
	    	error := "Invalid"
	    	return error
	}
	result1 := fmt.Sprintf("%.2f", result)
	str := text + " = " + result1
	return str
}

func get_val(f0, f2 string) (r0, r1 float64, err error){
	r0, err = strconv.ParseFloat(f0, 64)
	if err != nil {
		return
	}
	r1, err = strconv.ParseFloat(f2, 64)
	if err != nil {
		return
	}
	return
}