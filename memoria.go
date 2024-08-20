package graviola

type Memoria struct {
	offset  int
	setores [4096]byte
}

func NovaMemoria(offset int) *Memoria {
	m := new(Memoria)
	m.offset = offset
	return m
}

func (m *Memoria) Indice() int {
	return m.offset
}

func (m *Memoria) Adicionar(indice int, valor byte) {
	if indice < 512 || indice > 4095 {
		panic("endereco de memória inválido")
	}
	m.setores[indice-m.offset] = valor
}

func (m *Memoria) Obter(indice int) byte {
	if indice < 512 || indice > 4095 {
		panic("endereco de memória inválido")
	}
	return m.setores[indice-m.offset]
}
