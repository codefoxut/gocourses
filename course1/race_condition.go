/* Here we are using two function IncrementByValue and Print to check functioning of goroutines.
Here we are sharing a variable x in both the functions. when we execute this function. 
There are four goroutines and their output comes like

---------
go routine  1
after adding  3 9
after adding  5 6
go routine  9
main  1
---------
go routine  9
go routine  1
main  1
after adding  5 9
after adding  3 4
----------

With these outputs we can notice that there is no fix order in goroutines for execution and
 if we try to access any shared variable, there will be unexpected results. 



*/

package main

import (
	"fmt"
	"time"
)

var x int = 0

func main() {
	x += 1
	go Print()
	go IncrementByValue(5)
	go Print()
	go IncrementByValue(3)
	fmt.Println("main ", x)
	time.Sleep(500 * time.Millisecond)
}

func Print() {
	fmt.Println("go routine ", x)
}

func IncrementByValue(val int) {
	x += val
	fmt.Println("after adding ", val, x)
}
