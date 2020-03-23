package main


import (
	"io/ioutil"
    "fmt"
    "os"
)

func main() {

	var initialize string

	fmt.Printf("Please select one of the Following:\n1:begin new file\n2:add to existing file\n->:")
	fmt.Scan(&initialize)

if(initialize == "1"){
		mydata := []byte("List of Current Shiny Pokemon in Pokemon: Shield\n")

   	 	// the WriteFile method returns an error if unsuccessful
    	err := ioutil.WriteFile("pokeFile.data", mydata, 0777)
   		// handle this error
    	if err != nil {
        	// print it out
        	fmt.Println(err)
    	}
    }
    	data, err := ioutil.ReadFile("pokeFile.data")
    	if err != nil {
        	fmt.Println(err)
    	}

    	fmt.Print(string(data))
	

    //--------------------------------------------------------------------
	if(initialize == "2"){
    	f, err := os.OpenFile("pokeFile.data", os.O_APPEND|os.O_WRONLY, 0600)
    	if err != nil {
        	panic(err)
    	}
    	defer f.Close()

    	var name string
    	fmt.Printf("\nEnter name of Shiny to be Added\n")
    	fmt.Scan(&name)

    	if _, err = f.WriteString(name + "\n"); err != nil {
        	panic(err)
    	}

    	data, err = ioutil.ReadFile("pokeFile.data")
    	if err != nil {
        	fmt.Println(err)
    	}

    	fmt.Print(string(data))
	}
}