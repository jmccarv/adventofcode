package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type bits []byte

type packet struct {
	ver int
	id  int      // 4 == literal, anything else is an operator packet
	val uint64   // if it's a literal value packet
	sub []packet // operator packets contain sub packets
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	if !s.Scan() {
		panic("Invalid input")
	}

	// Convert to a more convenient format to parse
	packetData := explode(s.Text())

	// Now decode all packets
	packets := packetData.getPackets()

	fmt.Println("Part1", part1(packets))
	fmt.Println("Part2", packets[0].value())
}

func part1(packets []packet) (ret int) {
	for _, p := range packets {
		ret += p.ver + part1(p.sub)
	}
	return
}

func min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

func tf(b bool) uint64 {
	if b {
		return 1
	}
	return 0
}

func (p packet) value() (ret uint64) {
	switch p.id {
	case 4: // literal value
		ret = p.val
	case 0: // sum all sub packets
		for _, sub := range p.sub {
			ret += sub.value()
		}
	case 1: // product
		ret = 1
		for _, sub := range p.sub {
			ret *= sub.value()
		}
	case 2: // min
		ret = p.sub[0].value()
		for _, s := range p.sub[1:] {
			ret = min(ret, s.value())
		}
	case 3: // max
		ret = p.sub[0].value()
		for _, s := range p.sub[1:] {
			ret = max(ret, s.value())
		}
	case 5: // greater than
		ret = tf(p.sub[0].value() > p.sub[1].value())
	case 6: // less than
		ret = tf(p.sub[0].value() < p.sub[1].value())
	case 7:
		ret = tf(p.sub[0].value() == p.sub[1].value())
	}
	return
}

// each returned byte represents one bit from the incoming hexadecimal string
// so if the incoming string were 'A3' you'd get:
// [ 1 0 1 0 0 0 1 1 ]
func explode(hex string) (ret bits) {
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
			ret = append(ret, c)
			x <<= 1
		}
	}
	return
}

func (b *bits) getPackets() (ret []packet) {
	p := b.getNextPacket()
	for len(p) > 0 {
		ret = append(ret, p...)
		p = b.getNextPacket()
	}
	return
}

func (b *bits) getNextPacket() (ret []packet) {
	if len(*b) < 8 {
		// Ignore any trailing zeros in the encoded message
		return
	}

	p := packet{ver: int(b.next(3)), id: int(b.next(3))}
	switch p.id {
	case 4: // literal
		p.val = b.getLiteral()
		ret = append(ret, p)
	default: // operator packet, contains sub packets
		p.sub = append(p.sub, b.getSubPackets()...)
		ret = append(ret, p)
	}
	return
}

// Take the next x bits and convert them to an integer (our return value)
// Move our bits slice past those x bits
func (b *bits) next(x int) (ret uint64) {
	for _, i := range (*b)[0:x] {
		ret = (ret << 1) | uint64(i)
	}
	*b = (*b)[x:]
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

func (b *bits) getSubPackets() (ret []packet) {
	// length type ID (1 bit) immediately follows packet header
	switch b.next(1) {
	case 0: // next 15 bits are the total length in bits of the sub packets for this operator
		splen := int(b.next(15))
		sub := (*b)[0:splen]
		//fmt.Println("getOp() parsing packet(s)", sub)
		ret = append(ret, sub.getPackets()...)
		//fmt.Println("getOp, back", ret)
		b.next(splen)
	case 1: // next 11 bits are the number of sub packets contained by this packet
		nrSub := int(b.next(11))
		//fmt.Println("getOp() number sub packets", nrSub)
		for i := 0; i < nrSub; i++ {
			ret = append(ret, b.getNextPacket()...)
		}
	}
	return
}

func (b bits) String() string {
	ret := make([]byte, len(b))
	for i, x := range b {
		ret[i] = 48 + x
	}
	return string(ret)
}
