package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type bits struct {
	b []byte
}

type packet struct {
	ver int
	id  int    // 4 == literal, anything else is an op packet
	val uint64 // if it's a literal value packet
	//packets []packet // if it's an op packet, it contains other packets
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	if !s.Scan() {
		panic("Invalid input")
	}
	fmt.Println(s.Text())
	packetData := explode(s.Text())
	//fmt.Println(packetData)

	p := packetData.getPackets()
	fmt.Println(p)
	p1 := 0
	for _, x := range p {
		p1 += x.ver
	}
	fmt.Println("Part1", p1)
}

// each returned byte represents one bit from the incoming hexadecimal string
// so if the incoming string were 'A3' you'd get:
// [ 1 0 1 0 0 0 1 1 ]
func explode(hex string) (ret bits) {
	//ret = &bits{}
	for i := 0; i < len(hex); i += 2 {
		x, err := strconv.ParseUint(hex[i:i+2], 16, 8)
		if err != nil {
			panic(fmt.Sprintf("Invalid input: %v", err))
		}
		for b := 0; b < 8; b++ {
			var c byte = 1
			if x&0x80 == 0 {
				c = 0
			}
			ret.b = append(ret.b, c)
			x <<= 1
		}
	}
	return
}

func (b *bits) getPackets() (ret []packet) {
	for {
		p := b.getNextPacket()
		if len(p) == 0 {
			break
		}
		ret = append(ret, p...)
	}

	return
}

func (b *bits) getNextPacket() (ret []packet) {
	fmt.Println("getNextPacket()", b)

	var p packet
	if len(b.b) < 8 {
		return
	}
	p.ver = int(b.next(3))
	p.id = int(b.next(3))

	switch p.id {
	case 4:
		fmt.Println("getNextPacket(): literal", p)
		p.val = b.getLiteral()
		ret = append(ret, p)
	default:
		fmt.Println("getNextPacket(): operator", p)
		ret = append(ret, p)
		ret = append(ret, b.getOperator()...)
	}
	fmt.Println("getNextPacket()", ret)

	return
}

func (b *bits) next(x int) (ret uint64) {
	for _, i := range b.b[0:x] {
		ret = (ret << 1) | uint64(i)
	}
	b.b = b.b[x:]
	return
}

func (b *bits) getLiteral() (ret uint64) {
	// Literals are represented by groups of 5 bits
	// bit#1 == 0 flags the final 4 bits of the literal value
	for {
		ret <<= 4

		val := b.next(5)
		ret |= (val & 0x0F)

		if val&0x10 == 0 {
			break
		}
	}
	return
}

func (b *bits) getOperator() (ret []packet) {
	// lenght type ID (1 bit) immediately follows packet header
	switch b.next(1) {
	case 0:
		// next 15 bits are the total length in bits of the sub packets for this operator
		splen := int(b.next(15))
		sub := bits{b: b.b[0:splen]}
		fmt.Println("getOp() parsing packet(s)", sub)
		ret = append(ret, sub.getPackets()...)
		fmt.Println("getOp, back", ret)
		b.next(splen)
	case 1:
		// next 11 bits are the number of sub packets contained by this packet
		nrSub := int(b.next(11))
		fmt.Println("getOp() number sub packets", nrSub)
		for i := 0; i < nrSub; i++ {
			ret = append(ret, b.getNextPacket()...)
		}
	}
	return
}

func (b bits) String() string {
	ret := make([]byte, len(b.b))
	for i, x := range b.b {
		ret[i] = 48 + x
	}
	return string(ret)
}

func (p packet) String() string {
	return fmt.Sprintf("p{ver %d id %d val %d}", p.ver, p.id, p.val)
}
