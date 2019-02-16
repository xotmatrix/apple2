package core6502

import "fmt"

const (
	regA    = 0
	regX    = 1
	regY    = 2
	regP    = 4
	regSP   = 5
	regPC   = 6 // 2 bytes
	regNone = -1
)

const (
	flagN uint8 = 1 << 7
	flagV uint8 = 1 << 6
	flag5 uint8 = 1 << 5
	flagB uint8 = 1 << 4
	flagD uint8 = 1 << 3
	flagI uint8 = 1 << 2
	flagZ uint8 = 1 << 1
	flagC uint8 = 1 << 0
)

type Registers struct {
	data [8]uint8
}

func (r *Registers) getRegister(i int) uint8 { return r.data[i] }

func (r *Registers) getA() uint8  { return r.data[regA] }
func (r *Registers) getX() uint8  { return r.data[regX] }
func (r *Registers) getY() uint8  { return r.data[regY] }
func (r *Registers) getP() uint8  { return r.data[regP] }
func (r *Registers) getSP() uint8 { return r.data[regSP] }

func (r *Registers) setRegister(i int, v uint8) {
	r.data[i] = v
}
func (r *Registers) setA(v uint8)  { r.setRegister(regA, v) }
func (r *Registers) setX(v uint8)  { r.setRegister(regX, v) }
func (r *Registers) setY(v uint8)  { r.setRegister(regY, v) }
func (r *Registers) setP(v uint8)  { r.setRegister(regP, v) }
func (r *Registers) setSP(v uint8) { r.setRegister(regSP, v) }

func (r *Registers) getPC() uint16 {
	return uint16(r.data[regPC])*256 + uint16(r.data[regPC+1])
}

func (r *Registers) setPC(v uint16) {
	r.data[regPC] = uint8(v >> 8)
	r.data[regPC+1] = uint8(v)
}

func (r *Registers) getFlagBit(i uint8) uint8 {
	if r.getFlag(i) {
		return 1
	}
	return 0
}

func (r *Registers) getFlag(i uint8) bool {
	return (r.data[regP] & i) != 0
}

func (r *Registers) setFlag(i uint8) {
	r.data[regP] |= i
}

func (r *Registers) clearFlag(i uint8) {
	r.data[regP] &^= i
}

func (r *Registers) updateFlag(i uint8, v bool) {
	if v {
		r.setFlag(i)
	} else {
		r.clearFlag(i)
	}
}

func (r *Registers) updateFlagZN(t uint8) {
	r.updateFlag(flagZ, t == 0)
	r.updateFlag(flagN, t >= (1<<7))
}

func (r Registers) String() string {
	return fmt.Sprintf("A: %#02x, X: %#02x, Y: %#02x, SP: %#02x, PC: %#04x, P: %#02x, (NV-BDIZC): %08b",
		r.getA(), r.getX(), r.getY(), r.getSP(), r.getPC(), r.getP(), r.getP())
}
