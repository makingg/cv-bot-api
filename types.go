package main

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

type PartialDate struct {
	year, month, day uint8
}

type person struct {
	name string
	facts []string
	socials Socials
	achievements []Achievement
}
