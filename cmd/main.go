package main

import "graviola"

func main() {
	chip8 := graviola.Novo()
	defer chip8.Encerrar()
	chip8.CarregarROM("../roms/ibm.ch8")
	chip8.Executar()
}
