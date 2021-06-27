package main

import (
"fmt"
)
func isPrime(number int ) bool{
	for i:=2;i<number;i++{
		if number%i ==0{
			return false
		}
	}
	return true
}
func main()  {

	primes := make([]int,0)
	fmt.Println("Enter a number ->")
	var number int
	fmt.Scanf("%d",&number)

	for i:=2;i<number;i++ {
		if isPrime(i)==true{
			primes = append(primes,i)
		}
	}

	fmt.Println("There are ",len(primes)," prime number up to ",number )
	fmt.Println("This numbers are ",primes)
}