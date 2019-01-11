package main

import "fmt"

//Alright y'all today im gonna show you some pretty cool shit when it comes functions
//the first thing we are gonna look at is a function definition
//basically what it means is that you can define a function without actually writing what it means inside of it
//so let's check out what that looks like
type mystery func(a int) int

//so what the fuck does mystery do?  Well don't worry about that cause I wanna show something else too
//if you thought that was it then you'd be very in for a shock, cause you can actually pass a function into a function
//check this out
func passer(a func(a int) int) {
	fmt.Println(a(40)) //here im passing 40 into the function 'a' that was passed in
}

//shit's wild, right?

//so what does this look like in the main?

func main() {
	var fun mystery = func(a int) int {
		return a * 5
	}

	fmt.Println(fun(5))
	//why would you ever not want to define a function?
	//well maybe you don't know all the variables you are using yet
	//or maybe you just don't know how that'll affect your program until later
	//or maybe you just don't know how yet
	//regardless this is what we have to work with, so learn it

	var funner mystery = func(a int) int {
		return a + 5
	}

	//WHAT THE HELL MATTHEW YOU JUST DEFINED MYSTERY TO DO SOMETHING ELSE
	//FUCKING WATCH ME
	//JK so each variable can have its own definition of the 'mystery()' function
	//so 'fun' and 'funner' can have their own definition of the function
	//its kinda a form of overloading/polymorphism if you know what that is, if not its totally cool

	fmt.Println(funner(5))

	//and it worked, amazing right?

	//the other one is a little more basic so you get the gist
	//just make a new function kinda like below, assign it to a variable and pass it into the other function
	passie := func(b int) int {
		return b / 2
	}

	//done as so right here
	//prett easy just a cool mechanic
	passer(passie)

	//well that was cool and I hope y'all liked it
}
