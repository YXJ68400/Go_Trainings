package main
import "fmt"

func main(){
  fmt.Println("Hello World")


	i := 6
	j := &i
	fmt.Println("Value of i:", i)
	fmt.Println("Address of i:", &i)
	fmt.Println("Value of j :", j)
	fmt.Println("Value pointed to by j:", *j)
	*j = 10
	fmt.Println("New value of i:", i)


	a := 10
	if a%2 == 0 {
		fmt.Println(a, "is even")
	} else {
		fmt.Println(a, "is odd")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}
