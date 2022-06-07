package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error){
	fmt.Print(prompt)
	input,err:=r.ReadString('\n')

	return strings.TrimSpace(input),err
}

func createbill() bill{
	reader:=bufio.NewReader(os.Stdin)

	// fmt.Print("Create a new bill name: ")
	// name,_:= reader.ReadString('\n')
	// name=strings.TrimSpace(name)
	name,_:=getInput("Create a new bill name: ",reader)
    
	b:=newbill(name)
	fmt.Println("Created the bill -",b.name)

	return b
}

func promptOptions(b bill){
	reader:= bufio.NewReader(os.Stdin)

	opt,_:=getInput("Choose Options a- add item, s - save bill, t- add tip: ",reader)
	// fmt.Println(opt)
	switch opt{
	case "a":
		item,_:=getInput("Item Name: ",reader)
		price,_:=getInput("Item Price: ",reader) // the input is in string always, will be required to convert to float later
		
		p,err:=strconv.ParseFloat(price,64)   //enter the number of bits in the end
		
		if err!=nil{
			fmt.Println("The price must be a number...")
			promptOptions(b)
		}
		b.addItem(item,p)
		fmt.Println("Item added -",item, p)
		promptOptions(b)
	case "t":
		tip,_:=getInput("Enter the tip amount ($): ",reader)
		
		t,err:=strconv.ParseFloat(tip,64)
		if err!=nil{
			fmt.Println("The tip must be a number...")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("Tip updated!")
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("You chose save the bill- ",b.name)
	default: fmt.Println("That was not a valid input...")
			promptOptions(b)
	}
}

func main(){
	mybill:=createbill()
	promptOptions(mybill)
	//fmt.Println(mybill)
}
