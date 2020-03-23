package main
import (
 "fmt"
)

func main() {
 //reading an string
 var name string
 fmt.Print("What is your Name? ")
 fmt.Scan(&name)
 fmt.Println("Entered Name is :", name)

}