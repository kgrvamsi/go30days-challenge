package main


import (
  "fmt"
  "runtime"

)


func sayMyName(name string){
	for i:=0; i<5; i++ {
		runtime.Gosched()
		fmt.Println("Oh na na what's my name? ", name)
	}
	
}




func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println("You've got" runtime.NumCPU() " CPU cores yo!")

	sayMyName("basic")

	go sayMyName("Gopher")

	go sayMyName("Rabbit")

	var input string
	fmt.Scanln(&input)
	fmt.Println("ok dude! the show is over, go home and eat a potato")
}
