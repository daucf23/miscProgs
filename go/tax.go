package main

import( "fmt"  
		"math" 	
	  )


func main() {
	
	var salary float64
	var fedTax float64
	var menu string

	fmt.Printf("Type the number of the Tax Bracket your salary falls in.\n")
	fmt.Printf("Bracket 1: ($0-$9875)\n")
	fmt.Printf("Bracket 2: ($9876-$40125)\n")
	fmt.Printf("Bracket 3: ($40126-$85525)\n")
	fmt.Printf("Bracket 4: ($85526-$163300)\n")
	fmt.Printf("Bracket 5: ($163301-$207350)\n")
	fmt.Printf("Bracket 6: ($207351-$518400)\n")
	fmt.Printf("Bracket 7: (>$518401)\n")
	fmt.Printf("Selection:")
	fmt.Scan(&menu)
	
	



	switch{
		case menu == "1" : 
			fmt.Println("Please Enter Your Salary:")
			var temp float64
			fmt.Scan(&temp)
			salary = temp
			fedTax = salary*.10
			fedTax = Round(fedTax, .01)
			fmt.Println("\nHere is your estimated Federal Income Tax: $", fedTax)
			break
			
		case menu == "2" : 
			fmt.Println("Please Enter Your Salary:")
			var temp float64
			fmt.Scan(&temp)
			salary = temp
			fedTax = (salary-9875)*.12 + 987.50
			fedTax = Round(fedTax, .01)
			fmt.Println("\nHere is your estimated Federal Income Tax: $", fedTax)
			break
			
		case menu == "3" : 
			fmt.Println("Please Enter Your Salary:")
			var temp float64
			fmt.Scan(&temp)
			salary = temp
			fedTax = (salary-40125)*.22 + 4617.50
			fedTax = Round(fedTax, .01)
			fmt.Println("\nHere is your estimated Federal Income Tax: $", fedTax)
			break
			
		case menu == "4" : 
			fmt.Println("Please Enter Your Salary:")
			var temp float64
			fmt.Scan(&temp)
			salary = temp
			fedTax = (salary-85525)*.24 + 14605.50
			fedTax = Round(fedTax, .01)
			fmt.Println("\nHere is your estimated Federal Income Tax: $", fedTax)
			break
			
		case menu == "5" : 
			fmt.Println("Please Enter Your Salary:")
			var temp float64
			fmt.Scan(&temp)
			salary = temp
			fedTax = (salary-163300)*.32 + 33271.50
			fedTax = Round(fedTax, .01)
			fmt.Println("\nHere is your estimated Federal Income Tax: $", fedTax)
			break
			
		case  menu == "6" : 
			fmt.Println("Please Enter Your Salary:")
			var temp float64
			fmt.Scan(&temp)
			salary = temp
			fedTax = (salary-207350)*.35 + 47367.50
			fedTax = Round(fedTax, .01)
			fmt.Println("\nHere is your estimated Federal Income Tax: $", fedTax)
			break
			
		case menu == "7" : 
			fmt.Println("Please Enter Your Salary:")
			var temp float64
			fmt.Scan(&temp)
			salary = temp
			fedTax = (salary-518400)*.37 + 156235
			fedTax = Round(fedTax, .01)
			fmt.Println("\nHere is your estimated Federal Income Tax: $", fedTax)
			break
			
		}

}

func Round(x, unit float64) float64 {
    return math.Round(x/unit) * unit
}