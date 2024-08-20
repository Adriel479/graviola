package graviola

import (
	"log"
	"os"

	gc "github.com/rthornton128/goncurses"
)

type Chip8 struct {
	processador *Processador
	memoria     *Memoria
	monitor     *Monitor
}

func Novo() *Chip8 {
	memoria := NovaMemoria(512)
	monitor := NovoMonitor()
	return &Chip8{
		memoria: memoria,
		monitor: monitor,
		processador: NovoProcessador(
			memoria,
			monitor,
		),
	}
}

func (c *Chip8) CarregarROM(caminho string) {
	dados, erro := os.ReadFile(caminho)
	if erro != nil {
		log.Println("[ERROR] erro ao carregar ROM:", erro)
		os.Exit(1)
	}
	indice := 512
	for i := range dados {
		c.memoria.Adicionar(indice, dados[i])
		indice++
	}
}

func (c *Chip8) Executar() {
	c.processador.Executar()
}

func (c *Chip8) Encerrar() {
	gc.End()
}
