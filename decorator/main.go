package main

import "fmt"

type SetI interface {
	getString() string
}

type Set struct {
	str string
}

func (set Set) getString() string {
	return set.str
}

type ReverseSet struct {
	set SetI
}

func (reverseSet ReverseSet) getString() string {
	str := []rune(reverseSet.set.getString())

	for i := 0; i < len(str)/2; i++ {
		str[i], str[len(str)-i-1] = str[len(str)-i-1], str[i]
	}

	return string(str)
}

func main() {
	set := &Set{
		str: "table",
	}

	reverse := ReverseSet{
		set: set,
	}

	fmt.Printf("Set: %s\n", set.getString())
	fmt.Printf("ReverseSet: %s\n", reverse.getString())

	set.str = "топот"
	fmt.Printf("Set: %s\n", set.getString())
	fmt.Printf("ReverseSet: %s\n", reverse.getString())
}
