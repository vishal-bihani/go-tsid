package main

const (
	TSID_EPOCH  int64 = 1672531200000 // 2023-01-01T00:00:00.000Z
	TSID_BYTES  int8  = 8
	TSID_CHARS  int8  = 13
	RANDOM_BITS int8  = 22
	RANDOM_MASK int32 = 0x003fffff
)

var ALPHABET_UPPERCASE []rune = []rune("0123456789ABCDEFGHJKMNPQRSTVWXYZ")
var ALPHABET_LOWERCASE []rune = []rune("0123456789abcdefghjkmnpqrstvwxyz")
var ALPHABET_VALUES []int64

func init() {
	ALPHABET_VALUES = make([]int64, 128)
	for i := 0; i < len(ALPHABET_VALUES); i++ {
		ALPHABET_VALUES[i] = -1
	}

	// Numbers
	ALPHABET_VALUES['0'] = 0x00
	ALPHABET_VALUES['1'] = 0x01
	ALPHABET_VALUES['2'] = 0x02
	ALPHABET_VALUES['3'] = 0x03
	ALPHABET_VALUES['4'] = 0x04
	ALPHABET_VALUES['5'] = 0x05
	ALPHABET_VALUES['6'] = 0x06
	ALPHABET_VALUES['7'] = 0x07
	ALPHABET_VALUES['8'] = 0x08
	ALPHABET_VALUES['9'] = 0x09

	ALPHABET_VALUES['a'] = 0x0a
	ALPHABET_VALUES['b'] = 0x0b
	ALPHABET_VALUES['c'] = 0x0c
	ALPHABET_VALUES['d'] = 0x0d
	ALPHABET_VALUES['e'] = 0x0e
	ALPHABET_VALUES['f'] = 0x0f
	ALPHABET_VALUES['g'] = 0x10
	ALPHABET_VALUES['h'] = 0x11
	ALPHABET_VALUES['j'] = 0x12
	ALPHABET_VALUES['k'] = 0x13
	ALPHABET_VALUES['m'] = 0x14
	ALPHABET_VALUES['n'] = 0x15
	ALPHABET_VALUES['p'] = 0x16
	ALPHABET_VALUES['q'] = 0x17
	ALPHABET_VALUES['r'] = 0x18
	ALPHABET_VALUES['s'] = 0x19
	ALPHABET_VALUES['t'] = 0x1a
	ALPHABET_VALUES['v'] = 0x1b
	ALPHABET_VALUES['w'] = 0x1c
	ALPHABET_VALUES['x'] = 0x1d
	ALPHABET_VALUES['y'] = 0x1e
	ALPHABET_VALUES['z'] = 0x1f

	ALPHABET_VALUES['i'] = 0x01
	ALPHABET_VALUES['l'] = 0x01
	ALPHABET_VALUES['o'] = 0x00

	ALPHABET_VALUES['A'] = 0x0a
	ALPHABET_VALUES['B'] = 0x0b
	ALPHABET_VALUES['C'] = 0x0c
	ALPHABET_VALUES['D'] = 0x0d
	ALPHABET_VALUES['E'] = 0x0e
	ALPHABET_VALUES['F'] = 0x0f
	ALPHABET_VALUES['G'] = 0x10
	ALPHABET_VALUES['H'] = 0x11
	ALPHABET_VALUES['J'] = 0x12
	ALPHABET_VALUES['K'] = 0x13
	ALPHABET_VALUES['M'] = 0x14
	ALPHABET_VALUES['N'] = 0x15
	ALPHABET_VALUES['P'] = 0x16
	ALPHABET_VALUES['Q'] = 0x17
	ALPHABET_VALUES['R'] = 0x18
	ALPHABET_VALUES['S'] = 0x19
	ALPHABET_VALUES['T'] = 0x1a
	ALPHABET_VALUES['V'] = 0x1b
	ALPHABET_VALUES['W'] = 0x1c
	ALPHABET_VALUES['X'] = 0x1d
	ALPHABET_VALUES['Y'] = 0x1e
	ALPHABET_VALUES['Z'] = 0x1f

	ALPHABET_VALUES['I'] = 0x01
	ALPHABET_VALUES['L'] = 0x01
	ALPHABET_VALUES['O'] = 0x00
}

type tsid struct {
	number int64
}

func NewTsid(number int64) *tsid {
	return &tsid{
		number: number,
	}
}

func FromNumber(number int64) *tsid {
	return NewTsid(number)
}

func FromBytes(bytes []byte) *tsid {

	// TODO: Add validation

	var number int64 = 0

	number |= int64(bytes[0]&0xff) << 56
	number |= int64(bytes[1]&0xff) << 48
	number |= int64(bytes[2]&0xff) << 40
	number |= int64(bytes[3]&0xff) << 32
	number |= int64(bytes[4]&0xff) << 24
	number |= int64(bytes[5]&0xff) << 16
	number |= int64(bytes[6]&0xff) << 8
	number |= int64(bytes[7]) & 0xff

	return NewTsid(int64(number))
}

func FromString(str string) *tsid {
	arr := ToRuneArray(str)

	var number int64 = 0

	number |= ALPHABET_VALUES[arr[0]] << 60
	number |= ALPHABET_VALUES[arr[1]] << 55
	number |= ALPHABET_VALUES[arr[2]] << 50
	number |= ALPHABET_VALUES[arr[3]] << 45
	number |= ALPHABET_VALUES[arr[4]] << 40
	number |= ALPHABET_VALUES[arr[5]] << 35
	number |= ALPHABET_VALUES[arr[6]] << 30
	number |= ALPHABET_VALUES[arr[7]] << 25
	number |= ALPHABET_VALUES[arr[8]] << 20
	number |= ALPHABET_VALUES[arr[9]] << 15
	number |= ALPHABET_VALUES[arr[10]] << 10
	number |= ALPHABET_VALUES[arr[11]] << 5
	number |= ALPHABET_VALUES[arr[12]]

	return NewTsid(int64(number))
}

func ToRuneArray(str string) []rune {
	arr := []rune(str)

	if !IsValidRuneArray(arr) {
		return nil // TODO: Throw error
	}
	return arr
}

func IsValidRuneArray(arr []rune) bool {

	if arr == nil || len(arr) != int(TSID_CHARS) {
		return false
	}

	if (ALPHABET_VALUES[arr[0]] & 0b10000) != 0 {
		return false
	}

	for i := 0; i < len(arr); i++ {
		if ALPHABET_VALUES[arr[i]] == -1 {
			return false
		}
	}
	return true
}

func (t *tsid) ToLong() int64 {
	return t.number
}

func (t *tsid) ToBytes() []byte {
	bytes := make([]byte, TSID_BYTES)

	bytes[0] = byte(uint64(t.number) >> 56)
	bytes[1] = byte(uint64(t.number) >> 48)
	bytes[2] = byte(uint64(t.number) >> 40)
	bytes[3] = byte(uint64(t.number) >> 32)
	bytes[4] = byte(uint64(t.number) >> 24)
	bytes[5] = byte(uint64(t.number) >> 16)
	bytes[6] = byte(uint64(t.number) >> 8)
	bytes[7] = byte(t.number)

	return bytes
}

func (t *tsid) ToString() string {
	return t.ToStringWithAlphabets(ALPHABET_UPPERCASE)
}

func (t *tsid) ToStringWithAlphabets(alphabets []rune) string {
	chars := make([]rune, TSID_CHARS)

	chars[0] = alphabets[((uint64(t.number) >> 60) & 0b11111)]
	chars[1] = alphabets[((uint64(t.number) >> 55) & 0b11111)]
	chars[2] = alphabets[((uint64(t.number) >> 50) & 0b11111)]
	chars[3] = alphabets[((uint64(t.number) >> 45) & 0b11111)]
	chars[4] = alphabets[((uint64(t.number) >> 40) & 0b11111)]
	chars[5] = alphabets[((uint64(t.number) >> 35) & 0b11111)]
	chars[6] = alphabets[((uint64(t.number) >> 30) & 0b11111)]
	chars[7] = alphabets[((uint64(t.number) >> 25) & 0b11111)]
	chars[8] = alphabets[((uint64(t.number) >> 20) & 0b11111)]
	chars[9] = alphabets[((uint64(t.number) >> 15) & 0b11111)]
	chars[10] = alphabets[((uint64(t.number) >> 10) & 0b11111)]
	chars[11] = alphabets[((uint64(t.number) >> 5) & 0b11111)]
	chars[12] = alphabets[(uint64(t.number) & 0b11111)]

	return string(chars)
}
