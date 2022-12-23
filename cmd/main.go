package main

import (
	"fmt"

	"github.com/juergen-holtz/cal"
)

func main() {
	c, _ := cal.NewCalendar()
	fmt.Println(c)
}
