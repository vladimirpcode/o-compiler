package main

import (
	"fmt"
	"o-compiler/otext"
	"o-compiler/opars"

)

func main(){
	fmt.Println("Компилятор языка О")
	Init();
	opars.Compile();
	Done();

}

func Init(){
	otext.ResetText()
}

func Done(){
	otext.CloseText()
}