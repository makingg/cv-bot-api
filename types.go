package main

type PartialDate struct {
	day, month uint8 // [0,255]
	year uint16 // [0, 65535]
}

type Socials struct {
	github string
	telegram string
	mail string
	tel string
}

type Achievement struct {
	title string
	link string
	date PartialDate
}

type Person struct {
	name string
	facts []string
	socials Socials
	achievements []Achievement
}
