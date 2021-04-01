package main

import (
	"github.com/ansoni/termination"
	"math/rand"
	"strconv"
  "time"
	"fmt"
)

type Fish struct {
  kind string
  xPosition int
  yPosition int
}

func fishMovement(t *termination.Termination, e *termination.Entity, position termination.Position) termination.Position {
	data, _ := e.Data.(Fish)
	if data.xPosition < position.X {
		position.X-=1
	} else if data.xPosition > position.X {
		position.X+=1
	}

	if data.yPosition < position.Y {
		position.Y-=1
	} else if data.yPosition > position.Y {
		position.Y+=1
	}
	return position
}

func (f *Fish) fish(term *termination.Termination) {
	rand.Seed(time.Now().UnixNano())
	// Ukuran kolam (mengatur titik pergerakan)
	min := 5
	max := 30

	fish := term.NewEntity(termination.Position{f.xPosition,f.yPosition,0})
  fish.MovementCallback = fishMovement
	fish.Data = f

	for {
      data,_ := fish.Data.(Fish)
      data.xPosition = rand.Intn(max - min + 1) + min
      data.yPosition = rand.Intn(max - min + 1) + min
			f.xPosition = data.xPosition
			f.yPosition = data.yPosition
			s := fmt.Sprintf("X:%s Y:%s fish %s", strconv.Itoa(f.xPosition), strconv.Itoa(f.yPosition), f.kind)
			fish.Shape = termination.Shape {
			    "default": []string {
			    	s,
			    },
			}
      fish.Data = data
	}
}

func main()  {
	term := termination.New()
	term.FramesPerSecond = 2 // Atur kecepatan
	defer term.Close()
	gurame := Fish{ kind: "#GURAME", xPosition: 10, yPosition: 10 }
	bawal := Fish{ kind: "#BAWAL", xPosition: 20, yPosition: 10 }
	nila := Fish{ kind: "#NILA", xPosition: 20, yPosition: 20 }
	go gurame.fish(term)
	go bawal.fish(term)
	go nila.fish(term)
	go term.Animate()
	fmt.Scanln()
}
