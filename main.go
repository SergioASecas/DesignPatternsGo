package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	Value string
}

var instance *Singleton

var once sync.Once

func GetInstance() *Singleton {

	/*if instance == nil {
		instance = &Singleton{}
	}*/

	once.Do(func() {
		instance = &Singleton{Value: "Soy una instancia única"}
	})

	return instance
}

func main() {

	s1 := GetInstance()
	fmt.Println("Valor de la instancia s1:", s1.Value)

	s2 := GetInstance()
	fmt.Println("Valor de la instancia s2:", s2.Value)

	if s1 == s2 {
		fmt.Println("s1 y s2 son la misma instancia.")
	} else {
		fmt.Println("s1 y s2 son diferentes instancias.")
	}
}
