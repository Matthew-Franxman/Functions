package main

import "fmt"

//Alright y'all today im gonna show you some pretty cool shit when it comes functions
//the first thing we are gonna look at is a function definition
//basicaly what it means is that you can define a function without actually writing what it means inside of it
//so lets check out what that looks like
type mystery func(a int) int

//so what the fuck does mystery do?  Well don't worry about that cause I wanna show something else too
//if you thought that was it then you'd be very in for a shock, cause yopu can actually pass a function into a function
//check this out
func passer(a func(a int) int) {
	fmt.Println(a(40)) //here im passing 40 into the function 'a' that was passed in
}

//shits wild, right?

//so what does this look like in the main?

func main() {
	var fun mystery = func(a int) int {
		return a * 5
	}

	fmt.Println(fun(5))
	//why would you ever not want to define a function?
	//well maybe you don't know all the variables you are using yet
	//or maybe you just don't know how that'll affect your program until later
	//or maybe you just don't kniw how yet
	//regardless this is what we have to work wtih, so learn it

	var funner mystery = func(a int) int {
		return a + 5
	}

	fmt.Println(funner(5))

	passie := func(b int) int {
		return b / 2
	}

	passer(passie)
}
