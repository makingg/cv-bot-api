package main

import (
	"fmt"
)

func main() {
	socials := Socials{github: "github.com/evermake"}
	achs := []Achievement{
		Achievement {
			"winner of icpc",
			"~/",
			PartialDate{1, 6, 2025},
		},
	}
	p := Person{"Vlad", []string{"clever", "stylish"}, socials, achs}
	fmt.Println(p)
}
