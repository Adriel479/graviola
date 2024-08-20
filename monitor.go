package graviola

import (
	"log"

	gc "github.com/rthornton128/goncurses"
)

const (
	col = 64
	lin = 32
)

type Monitor struct {
	buffer [lin][col]byte
	stdscr *gc.Window
}

func NovoMonitor() *Monitor {
	m := &Monitor{}
	stdscr, err := gc.Init()
	if err != nil {
		log.Fatal(err)
	}
	m.stdscr = stdscr
	return m
}

func (m *Monitor) Limpar() {
	//fmt.Print("\033[H\033[2J")
	m.stdscr.Erase()
	for i := 0; i < lin; i++ {
		for j := 0; j < col; j++ {
			m.buffer[i][j] = 0
		}
	}
}

func (m *Monitor) Comprimento() int {
	return col
}

func (m *Monitor) Altura() int {
	return lin
}

func (m *Monitor) Pixel(lin, col int) byte {
	return m.buffer[lin][col]
}

func (m *Monitor) Desenhar(lin, col int, valor byte) {
	m.buffer[lin][col] = valor
	if m.buffer[lin][col] == 1 {
		m.stdscr.MovePrintf(lin, col, "@")
	} else {
		m.stdscr.MovePrintf(lin, col, " ")
	}
}

func (m *Monitor) Renderizar() {
	m.stdscr.Refresh()
	/*for i := 0; i < lin; i++ {
		for j := 0; j < col; j++ {
			if m.buffer[i][j] == 0 {
				fmt.Printf(" ")
			} else {
				fmt.Print("@")
			}
		}
		fmt.Println()
	}*/
}
