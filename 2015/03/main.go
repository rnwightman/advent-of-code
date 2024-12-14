package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	raw, _ := reader.ReadString('\n')
	movements := parseMovements(raw)

	deliveredTo := make(map[coord]int)

	rxDelivery := make(chan coord, 10)
	txs := make([]chan movement, 0)

	// send out the santas
	const numSantas = 2
	for i := 0; i < numSantas; i += 1 {
		tx := make(chan movement)
		txs = append(txs, tx)

		go deliver(tx, rxDelivery)
	}
	deliveredTo[coord{0, 0}] = numSantas

	// Send out the instructions
	for i, m := range movements {
		txIndex := i % len(txs)

		log.Printf("Instructing %d to %v", txIndex, m)
		txs[txIndex] <- m

		d := <-rxDelivery

		count, ok := deliveredTo[d]
		if !ok {
			count = 0
		}
		deliveredTo[d] = count + 1
	}

	for _, tx := range txs {
		close(tx)
	}

	locs := len(deliveredTo)

	fmt.Printf("Delivered to %d houses\n", locs)
}

func deliver(instruction chan movement, delivered chan coord) {
	loc := coord{0, 0}
	log.Printf("Setting out %v\n", loc)

	for m := range instruction {
		log.Printf("Instructed to %v\n", m)

		loc = loc.move(m)

		log.Printf("Reporting delivery to %v\n", loc)
		delivered <- loc
	}
}

func parseMovements(s string) []movement {
	result := make([]movement, 0, len(s))

	for _, c := range s {
		var m movement
		switch c {
		case '^':
			m = north
		case '>':
			m = east
		case 'v':
			m = south
		case '<':
			m = west

		default:
			continue
		}

		result = append(result, m)
	}

	return result
}

func track(c coord, movements []movement) map[coord]int {
	deliveries := make(map[coord]int)

	// record starting
	deliveries[c] = 1

	for _, m := range movements {
		c = c.move(m)

		d, ok := deliveries[c]
		if !ok {
			d = 0
		}

		deliveries[c] = d + 1
	}

	return deliveries
}

type coord struct {
	x, y int
}

func (c coord) move(m movement) coord {
	return coord{
		x: c.x + m.xDist,
		y: c.y + m.yDist,
	}
}

type movement struct {
	xDist, yDist int
}

var (
	north = movement{0, 1}
	east  = movement{1, 0}
	south = movement{0, -1}
	west  = movement{-1, 0}
)
