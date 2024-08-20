package graviola

const (
	CLS       = 0x00E0
	RET       = 0x00EE
	SYS       = 0x0000
	JP_ADD    = 0x1000
	CALL      = 0x2000
	SE_VX     = 0x3000
	SNE_VX    = 0x4000
	SE_VX_VY  = 0x5000
	LD_VX     = 0x6000
	ADD_VX    = 0x7000
	LD_VX_VY  = 0x8000
	OR        = 0x8001
	AND       = 0x8002
	XOR       = 0x8003
	ADD_VX_VY = 0x8004
	SUB       = 0x8005
	SHR       = 0x8006
	SUBN      = 0x8007
	SHL       = 0x800E
	SNE_VX_VY = 0x9000
	LD_I_ADD  = 0xA000
	JP_RV     = 0xB000
	RND       = 0xC000
	DRW       = 0xD000
	SKP       = 0xE09E
	SKNP      = 0xE0A1
	LD_VX_DT  = 0xF007
	LD_VX_K   = 0xF00A
	LD_DT_VX  = 0xF015
	LD_ST_VX  = 0xF018
	ADD_I_VX  = 0xF01E
	LD_F_VX   = 0xF029
	LD_B_VX   = 0xF033
	LD_I_VX   = 0xF055
	LD_VX_I   = 0xF065
)
