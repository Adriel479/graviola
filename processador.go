package graviola

import "time"

type Processador struct {
	contadorPrograma    int
	registradorPilha    int
	registradorI        int
	memoria             *Memoria
	monitor             *Monitor
	registradorVariavel [16]byte
}

type Instrucao struct {
	cmd int
	x   byte
	y   byte
	n   byte
	kk  byte
	nnn int
}

func NovoProcessador(memoria *Memoria, monitor *Monitor) *Processador {
	return &Processador{
		memoria:          memoria,
		monitor:          monitor,
		contadorPrograma: memoria.Indice(),
		registradorPilha: 4095 - 100,
	}
}

/*func (p *Processador) push(item int) {
	p.memoria.Adicionar(p.registradorPilha, byte(item>>8))
	p.memoria.Adicionar(p.registradorPilha-1, byte(item&0b0000000011111111))
	p.registradorPilha -= 2
}*/

func (p *Processador) pop() int {
	parteA := int(p.memoria.Obter(p.registradorPilha))
	parteB := int(p.memoria.Obter(p.registradorPilha + 1))
	p.registradorPilha += 2
	return parteA<<8 | parteB
}
func (p *Processador) buscarInstrucao() int {
	parteA := int(p.memoria.Obter(p.contadorPrograma))
	parteB := int(p.memoria.Obter(p.contadorPrograma + 1))
	p.contadorPrograma += 2
	return int(parteA)<<8 | parteB
}

func (p *Processador) decodificar(entrada int) *Instrucao {
	instrucao := new(Instrucao)
	instrucao.x = byte(entrada & 0b0000111100000000 >> 8)
	instrucao.y = byte(entrada & 0b0000000011110000 >> 4)
	instrucao.n = byte(entrada & 0b0000000000001111)
	instrucao.kk = byte(entrada & 0b0000000011111111)
	instrucao.nnn = entrada & 0b0000111111111111
	if entrada == 0x00e0 || entrada == 0x00ee {
		return instrucao
	}
	operacao := entrada & 0b1111000000000000
	switch operacao {
	case 0x1000:
		fallthrough
	case 0x2000:
		fallthrough
	case 0x3000:
		fallthrough
	case 0x4000:
		fallthrough
	case 0x6000:
		fallthrough
	case 0x7000:
		fallthrough
	case 0xA000:
		fallthrough
	case 0xB000:
		fallthrough
	case 0xC000:
		fallthrough
	case 0xD000:
		instrucao.cmd = 0b1111000000000000
	case 0x5000:
		fallthrough
	case 0x8000:
		fallthrough
	case 0x9000:
		instrucao.cmd = 0b1111000000001111
	case 0xE000:
		fallthrough
	case 0xF000:
		instrucao.cmd = 0b1111000011111111
	}
	instrucao.cmd = entrada & instrucao.cmd
	return instrucao
}

func (p *Processador) desenhar(instr *Instrucao) {
	const padrao = byte(0b10000000)
	var (
		col = int(p.registradorVariavel[instr.x] & byte(p.monitor.Comprimento()-1))
		lin = int(p.registradorVariavel[instr.y] & byte(p.monitor.Altura()-1))
	)

	for i := 0; i < int(instr.n); i++ {
		sprite := p.memoria.Obter(p.registradorI + i)
		for j := 0; j < 8; j++ {
			if col+j > p.monitor.Comprimento()-1 {
				break
			}
			resultado := p.monitor.Pixel(lin, col+j) ^ ((sprite & (padrao >> j)) >> (7 - j))
			p.registradorVariavel[0xf] = resultado
			p.monitor.Desenhar(lin, col+j, resultado)
		}
		lin++
		if lin > p.monitor.Altura()-1 {
			break
		}
	}
}

func (p *Processador) Executar() {
	for {
		instrucao := p.decodificar(p.buscarInstrucao())
		switch instrucao.cmd {
		case CLS:
			p.monitor.Limpar()
		case RET:
			p.contadorPrograma = p.pop()
		case JP_ADD:
			p.contadorPrograma = instrucao.nnn
		case ADD_VX:
			if p.registradorVariavel[instrucao.x]+instrucao.kk <= 255 {
				p.registradorVariavel[instrucao.x] += instrucao.kk
			}
		case LD_VX:
			p.registradorVariavel[instrucao.x] = instrucao.kk
		case LD_I_ADD:
			p.registradorI = instrucao.nnn
		case DRW:
			p.desenhar(instrucao)
			p.monitor.Renderizar()
		default:
		}
		time.Sleep(time.Millisecond * 3)
	}
}
