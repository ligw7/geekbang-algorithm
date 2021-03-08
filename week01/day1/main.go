package main

import "fmt"

func main() {
	n:=45
	m:=climbStairs(n)
	fmt.Print(m)
}

func climbStairs(n int) int {
	z:=-1
	if n==1 {
		z=1
	}
	if n==2{
		z=2
	}
	if n>2{
		x:=1
		y:=2
		for i:=3;i<n+1;i++{
			z=x+y
			x=y
			y=z
		}
	}
	return z
}