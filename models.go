package main

import "fmt"


type Customer struct {
	username string
	age uint16
	phoneNumber string
}

func (c Customer) str() string {
	return fmt.Sprintf("%s %d", c.username, c.age)
}

func (c *Customer) setPhoneNumber(number string) {
	c.phoneNumber = number
}
