/*  This program is designed to show an example of some of the different 
	capabilities that Go offers, from datatypes to different methods of 	
	outputing strings, integers and other data.
*/

package main

import ("fmt"  		//this is standard
		"math"		//math libs
		"strings"	//string libs
		"errors")	//error libs

/*gloabal variable declaration*/
var global int = 50;

/* STRUCT */
type Books struct {
   title string
   author string
   subject string
   book_id int
}

/* recursive funtion test*/
func factorial(i int)int {
   if(i <= 1) {
      return 1
   }
   return i * factorial(i - 1)
}

func fibonaci(i int) (ret int) {
   if i == 0 {
      return 0
   }
   if i == 1 {
      return 1
   }
   return fibonaci(i-1) + fibonaci(i-2)
}

/*defining interfaces*/
/* define an interface */
type Shape interface {
   area() float64
}

/* define a circle */
type Circle struct {
   x,y,radius float64
}

/* define a rectangle */
type Rectangle struct {
   width, height float64
}

/* define a method for circle (implementation of Shape.area())*/
func(circle Circle) area() float64 {
   return math.Pi * circle.radius * circle.radius
}

/* define a method for rectangle (implementation of Shape.area())*/
func(rect Rectangle) area() float64 {
   return rect.width * rect.height
}

/* define a method for shape */
func getArea(shape Shape) float64 {
   return shape.area()
}


func main() {

	var x float64									//declaring a float

	x = 69.0
	y := 20											//declaring integer 

	const LENGTH int = 10							//declaring constants
	const WIDTH int = 10

	var z, a = 21, "foo" 							//multiple declarations in a single line (type int and String)

	fmt.Println(x,a,z) 								//print line
	fmt.Printf("x is of type of %T\n", x) 			//prints the datatype of variable x
	fmt.Printf("y is of type %T\n", y) 				//print format
	fmt.Println("hello world\a")

	var area int
	area = LENGTH * WIDTH							//performing math with constants and 
													//storing it in int variable
	
	fmt.Printf("value of the area is: %d\n", area)	//displaying variable that has been calculated

/********************************************************************************/
	   fmt.Printf("\n\n---------------------\n\n") //spacing

	if(y < z){										// basic if statement
		fmt.Printf("y: %d is less than x: %d\n" , y, z)
	}
/********************************************************************************/
	   fmt.Printf("\n\n---------------------\n\n") //spacing

	 /* example of switch statements in go */
   var grade string = "B"
   var marks int = 90

   switch marks {
      case 90: grade = "A"
      case 80: grade = "B"
      case 50,60,70 : grade = "C"
      default: grade = "D"  
   }
   switch {
      case grade == "A" :
         fmt.Printf("Excellent!\n" )     
      case grade == "B", grade == "C" :
         fmt.Printf("Well done\n" )      
      case grade == "D" :
         fmt.Printf("You passed\n" )      
      case grade == "F":
         fmt.Printf("Better try again\n" )
      default:
         fmt.Printf("Invalid grade\n" );
   }
   fmt.Printf("Your grade is  %s\n", grade ); 


   /* Example of select statement which is similar to a switch statement*/
     var c1, c2, c3 chan int
   var i1, i2 int
   select {
      case i1 = <-c1:
         fmt.Printf("received ", i1, " from c1\n")
      case c2 <- i2:
         fmt.Printf("sent ", i2, " to c2\n")
      case i3, ok := (<-c3):  // same as: i3, ok := <-c3
         if ok {
            fmt.Printf("received ", i3, " from c3\n")
         } else {
            fmt.Printf("c3 is closed\n")
         }
      default:
         fmt.Printf("no communication\n")
   }   
/********************************************************************************/
      fmt.Printf("\n\n---------------------\n\n") //spacing

													    //infinite loop statement, but i put a stop
	for true  {
       fmt.Printf("This loop will run forever.\n")
       if(y == 20){										//this kills the loop
       	break
       }
   }
/********************************************************************************/
      fmt.Printf("\n\n---------------------\n\n") //spacing

    /* this section is focused on an example of a function */
    /* calling a function to get max value between 20 and 21*/
   var ret int
   ret = max(y, z)

   fmt.Printf( "Max value is : %d\n", ret )


   /* multiple variables returned from a function*/
   c, d := swap("Almeida II", "David")
   fmt.Println(c, d)
/********************************************************************************/
      fmt.Printf("\n\n---------------------\n\n") //spacing

   /* local variable declaration */
   var f, g, h int 

   /* actual initialization */
   f = 10
   g = 20
   h = f + g

   fmt.Printf ("value of f = %d, g = %d and h = %d\n", f, g, h)
   fmt.Printf ("Global variable value is : %d\n", global)


  /********************************************************************************/
   /* This section deals with strings*/
   fmt.Printf("\n\n---------------------\n\n") //spacing
   var greeting = "Hello World!"

   fmt.Printf("normal string: ")
   fmt.Printf("%s", greeting)
   fmt.Printf("\n")
   fmt.Printf("hex bytes: ")

   for i := 0; i < len(greeting); i++ {
       fmt.Printf("%x ", greeting[i])
   }
   
   fmt.Printf("\n")
   const sampleText = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98" 
   
   /*q flag escapes unprintable characters, with + flag it escapses non-ascii 
   characters as well to make output unambigous */
   fmt.Printf("quoted string: ")
   fmt.Printf("%+q", sampleText)
   fmt.Printf("\n") 
	

   /* String length*/
   fmt.Printf("\nString Length is: ")
   fmt.Println(len(greeting))

   /*String Concatenation*/
   greetings :=  []string{"Hello","world!"}   
   fmt.Println(strings.Join(greetings, " "))
   /********************************************************************************/
   fmt.Printf("\n\n---------------------\n\n") //spacing

   /* Declaring and using Arrays*/
   var n [10]int /* n is an array of 10 integers */
   var i,j int

   /* initialize elements of array n to 0 */         
   for i = 0; i < 10; i++ {
      n[i] = i + 100 /* set element at location i to i + 100 */
   }
   
   /* output each array element's value */
   for j = 0; j < 10; j++ {
      fmt.Printf("Element[%d] = %d\n", j, n[j] )
   }
 /********************************************************************************/
   /*Working with pointers*/
   fmt.Printf("\n\n---------------------\n\n") //spacing

   var ip *int      /* pointer variable declaration */

   ip = &z  /* store address of a in pointer variable*/

   fmt.Printf("Address of a variable: %x\n", &z  )

   /* address stored in pointer variable */
   fmt.Printf("Address stored in ip variable: %x\n", ip )

   /* access the value using the pointer */
   fmt.Printf("Value of *ip variable: %d\n\n", *ip )

   var  ptr *int

   fmt.Printf("The value of ptr is : %x\n", ptr  )

   /********************************************************************************/
   fmt.Printf("\n\n---------------------\n\n") //spacing
   /* Usitng and accessing structures*/
   var Book1 Books    /* Declare Book1 of type Book */
   var Book2 Books    /* Declare Book2 of type Book */
 
   /* book 1 specification */
   Book1.title = "Go Programming"
   Book1.author = "Mahesh Kumar"
   Book1.subject = "Go Programming Tutorial"
   Book1.book_id = 6495407

   /* book 2 specification */
   Book2.title = "Telecom Billing"
   Book2.author = "Zara Ali"
   Book2.subject = "Telecom Billing Tutorial"
   Book2.book_id = 6495700
 
   /* print Book1 info */
   fmt.Printf( "Book 1 title : %s\n", Book1.title)
   fmt.Printf( "Book 1 author : %s\n", Book1.author)
   fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
   fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)

   /* print Book2 info */
   fmt.Printf( "Book 2 title : %s\n", Book2.title)
   fmt.Printf( "Book 2 author : %s\n", Book2.author)
   fmt.Printf( "Book 2 subject : %s\n", Book2.subject)
   fmt.Printf( "Book 2 book_id : %d\n", Book2.book_id)

   fmt.Printf("\n\n")

/* print Book1 info */
   printBook(Book1)

   /* print Book2 info */
   printBook(Book2)

   fmt.Printf("\n\n")


   /* pointers to structures*/
    /* print Book1 info */
   printBook2(&Book1)

   /* print Book2 info */
   printBook2(&Book2)

 /********************************************************************************/
   fmt.Printf("\n\n---------------------\n\n") //spacing
   /* go slices (portions of an array that are accessed specifically)*/

   var numbers = make([]int,3,5)
   printSlice(numbers)

    /********************************************************************************/
   fmt.Printf("\n\n\n\n") //spacing
   /* go range is a keyword, has special purpose*/

   /* create a slice */
    numbers2 := []int{0,1,2,3,4,5,6,7,8} 
   
   /* print the numbers */
   for i:= range numbers2 {
      fmt.Println("Slice item",i,"is",numbers2[i])
   }
   
   /* create a map*/
   countryCapitalMap := map[string] string {"France":"Paris","Italy":"Rome","Japan":"Tokyo"}
   
   /* print map using keys*/
   for country := range countryCapitalMap {
      fmt.Println("Capital of",country,"is",countryCapitalMap[country])
   }
   
   /* print map using key-value*/
   for country,capital := range countryCapitalMap {
      fmt.Println("Capital of",country,"is",capital)
   }


    /********************************************************************************/
   fmt.Printf("\n\n---------------------\n\n") //spacing
   /* Using maps*/

   var countryCapitalMap2 map[string]string
   /* create a map*/
   countryCapitalMap2 = make(map[string]string)
   
   /* insert key-value pairs in the map*/
   countryCapitalMap2["France"] = "Paris"
   countryCapitalMap2["Italy"] = "Rome"
   countryCapitalMap2["Japan"] = "Tokyo"
   countryCapitalMap2["India"] = "New Delhi"
   
   /* print map using keys*/
   for country := range countryCapitalMap2 {
      fmt.Println("Capital of",country,"is",countryCapitalMap2[country])
   }
   
   /* test if entry is present in the map or not*/
   capital, ok := countryCapitalMap2["United States"]
   
   /* if ok is true, entry is present otherwise entry is absent*/
   if(ok){
      fmt.Println("Capital of United States is", capital)  
   } else {
      fmt.Println("Capital of United States is not present") 
   }


    /********************************************************************************/
   fmt.Printf("\n\n---------------------\n\n") //spacing
   /* Using recursion*/

   /*factoring*/
   var l int = 15
   fmt.Printf("Factorial of %d is %d\n\n", l, factorial(l))

   /*fibonacci*/
   var m int
   for m = 0; m < 10; m++ {
      fmt.Printf("%d ", fibonaci(m))
   }

   /********************************************************************************/
   fmt.Printf("\n\n---------------------\n\n") //spacing

   /*type casting*/
   var sum int = 17
   var count int = 5
   var mean float32
   
   mean = float32(sum)/float32(count)
   fmt.Printf("Value of mean : %f\n",mean)

/********************************************************************************/
   fmt.Printf("\n\n---------------------\n\n") //spacing
   /* using interfaces*/

   circle := Circle{x:0,y:0,radius:5}
   rectangle := Rectangle {width:10, height:5}
   
   fmt.Printf("Circle area: %f\n",getArea(circle))
   fmt.Printf("Rectangle area: %f\n",getArea(rectangle))

   /********************************************************************************/
   fmt.Printf("\n\n---------------------\n\n") //spacing
   /* error handling */
	result, err:= Sqrt(-1)

   if err != nil {
      fmt.Println(err)
   } else {
      fmt.Println(result)
   }
   
   result, err = Sqrt(9)

   if err != nil {
      fmt.Println(err)
   } else {
      fmt.Println(result)
   }



}


func max(num1, num2 int) int {
   /* local variable declaration */
   var result int

   if (num1 > num2) {
      result = num1
   } else {
      result = num2
   }
   return result 
}

func swap(x, y string) (string, string) {
   return y, x
}

/* Using a struct with functions*/

func printBook( book Books ) {
   fmt.Printf( "Book title : %s\n", book.title);
   fmt.Printf( "Book author : %s\n", book.author);
   fmt.Printf( "Book subject : %s\n", book.subject);
   fmt.Printf( "Book book_id : %d\n", book.book_id);
}

/* unsing pointers from a struct to a function*/
func printBook2( book *Books ) {
   fmt.Printf( "Book title : %s\n", book.title);
   fmt.Printf( "Book author : %s\n", book.author);
   fmt.Printf( "Book subject : %s\n", book.subject);
   fmt.Printf( "Book book_id : %d\n", book.book_id);
}


/* slicing function for array slice*/
func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

/* function for error handling test*/
func Sqrt(value float64)(float64, error) {
   if(value < 0){
      return 0, errors.New("Math: negative number passed to Sqrt")
   }
   return math.Sqrt(value), nil
}

