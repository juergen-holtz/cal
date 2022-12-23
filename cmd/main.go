package main

import (
	"fmt"

	"github.com/juergen-holtz/cal"
)

func main() {
	c := cal.NewCalendar(cal.ForMonth(11), cal.ForYear(2023))
	fmt.Println(c)
}
