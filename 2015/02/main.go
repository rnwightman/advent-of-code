package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	paperRequired := 0
	ribbonRequired := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		p := parseCuboid(line)

		paperRequired += p.paperRequired()
		ribbonRequired += p.ribbonRequired()
	}

	fmt.Printf("Paper Required: %d square feet\n", paperRequired)
	fmt.Printf("Ribbon Required: %d linear feet\n", ribbonRequired)
}

func parseCuboid(s string) cuboid {
	measures := strings.Split(s, "x")
	length, _ := strconv.ParseInt(measures[0], 10, 32)
	width, _ := strconv.ParseInt(measures[1], 10, 32)
	height, _ := strconv.ParseInt(measures[2], 10, 32)

	return cuboid{int(length), int(width), int(height)}
}

type cuboid struct {
	length, width, height int
}

func (c *cuboid) paperRequired() int {
	slack := slices.Min(c.areas())
	area := c.area()

	return slack + area
}

func (c *cuboid) ribbonRequired() int {
	bow := c.volume()
	wrap := slices.Min(c.perimeters())

	log.Printf("ribbon for %v: bow=%d\twrap=%d", c, bow, wrap)

	return bow + wrap
}

func (c *cuboid) area() int {
	area := 0

	for _, f := range c.areas() {
		area = area + f
	}

	return area
}

func (c *cuboid) volume() int {
	return c.width * c.length * c.height
}

func (c *cuboid) areas() []int {
	return []int{
		c.width * c.length,
		c.width * c.length,

		c.width * c.height,
		c.width * c.height,

		c.height * c.length,
		c.height * c.length,
	}
}

func (c *cuboid) perimeters() []int {
	return []int{
		2*c.width + 2*c.height,

		2*c.height + 2*c.length,

		2*c.length + 2*c.width,
	}
}
