package main

import (
	"fmt"
	"os"
)

//defining the custom user type using struct
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

//function to make new bills
func newbill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

//----- There can be functions associated with struct objects
//--- Can be used like mybill.format(), this is called a receiver function

// Receiver funtion

func (b *bill) format() string { 
	fs := "Bill breakdown\n"
	var total float64 = 0
	// List items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total+=v
	}
	total+=b.tip

	// Adding the tip to the formated string
	fs+=fmt.Sprintf("%-25v ...$%v\n","tip:",b.tip)

	//Total bill value
	fs+= fmt.Sprintf("%-25v ...$%0.2f\n","total:",total)  // -25 means the variable value is of 25 char long, if shorter then filled with empty spaces
	return fs

}

// Receiver function to update the tip
func (b *bill) updateTip(tip float64){
	b.tip=tip        // structs are dereferenced automatically in Go, otherwise (*b).tip=tip
}

// Receiver function for add items to the menu
func (b bill) addItem(name string, price float64){
	b.items[name]=price
}

//Receiver Funtion to save the bills 
func (b *bill) save(){  // whenever we are saving the data we are saving it in slice of bytes format
	data:= []byte(b.format()) // we gett a string and it is converted into slice of bytes and then is being saved to the folder as shown below

	err:= os.WriteFile("saved_bills/"+b.name+".txt",data,0644)// 3 arguments, path, data and then 0644 is the permissions for the fils
	if err!=nil{
		panic(err)
	}
	fmt.Println("Bill was saved to file")
}