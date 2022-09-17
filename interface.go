package main

import (
	"fmt"
)

type phone struct {
	name         string
	model        string
	memory       string
	front_camera string
	back_camera  string
	version      float32
}

//PHONE ...
func (p phone) Display() {
	fmt.Printf("%s %s %s %s %s %f\n", p.name, p.model, p.memory, p.front_camera, p.back_camera, p.version)
}

type laptop struct {
	name string
	RAM  string
	VRAM string
	HDD  string
	SDD  string
	CPU  string
}

//LAPTOP ...
func (l laptop) Display() {
	fmt.Printf("%s %s %s %s %s %s\n", l.name, l.RAM, l.VRAM, l.HDD, l.SDD, l.CPU)
}

type book struct {
	title          string
	author         string
	ganre          string
	total_pages    int32
	isbn           string
	publashed_date int
}

//BOOK ...
func (b book) Display() {
	fmt.Printf("%s %s %s %d %s %d\n", b.title, b.author, b.ganre, b.total_pages, b.isbn, b.publashed_date)
}

type car struct {
	name       string
	color      string
	model      string
	high_speed int32
	year       string
}

func (c car) Display() {
	fmt.Printf("%s %s %s %d %s\n", c.name, c.color, c.model, c.high_speed, c.year)
}

type display interface {
	Display()
}

func main() {
	var I display
	phone1 := phone{
		name:         "readme 5",
		model:        "README",
		memory:       "32 GB",
		front_camera: "5MP",
		back_camera:  "20MP",
		version:      2.3,
	}
	I = phone1
	// I.Display()
	laptop1 := laptop{
		name: "acer",
		RAM:  "4GB",
		VRAM: "NIVIDEA 2GB",
		HDD:  "4GB",
		SDD:  "512GB",
		CPU:  "Core i5 1.6GHz, 3.4GHz",
	}
	I = laptop1
	switch I.(type) {
	case phone:
		fmt.Println("Phone: ")
	case laptop:
		fmt.Println("Laptop: ")
	case book:
		fmt.Println("Book: ")
	case car:
		fmt.Println("Car: ")
	default:
		fmt.Println("no devace: ")
	}
	I.Display()
}
