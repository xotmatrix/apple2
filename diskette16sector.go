package apple2

/*
See:
	"Beneath Apple DOS" https://fabiensanglard.net/fd_proxy/prince_of_persia/Beneath%20Apple%20DOS.pdf
	https://github.com/TomHarte/CLK/wiki/Apple-GCR-disk-encoding
*/

type diskette16sector struct {
	nib      *fileNib
	position int
}

func newDisquette16Sector(f *fileNib) *diskette16sector {
	var d diskette16sector
	d.nib = f
	return &d
}

func (d *diskette16sector) powerOn(cycle uint64) {
	// Not used
}
func (d *diskette16sector) powerOff(_ uint64) {
	// Not used
}

func (d *diskette16sector) read(quarterTrack int, cycle uint64) uint8 {
	track := d.nib.track[quarterTrack/stepsPerTrack]
	value := track[d.position]
	d.position = (d.position + 1) % nibBytesPerTrack
	//fmt.Printf("%v, %v, %v, %x\n", 0, 0, d.position, uint8(value))
	return value
}

func (d *diskette16sector) write(quarterTrack int, value uint8, _ uint64) {
	track := quarterTrack / stepsPerTrack
	d.nib.track[track][d.position] = value
	d.position = (d.position + 1) % nibBytesPerTrack
}
