package bf_interpreter

type Machine struct {
	code string
	ip   int

	memory [30000]int
	dp     int
	buf    []byte

	input  io.Reader
	output io.Writer
}

func NewMachine(code string, in io.Reader, out io.Writer) *Machine {
	return &Machine{
		code:   code,
		buf:    make([]byte, 1),
		input:  in,
		output: out,
	}
}

func (m *Machine) Execute() {
	for m.ip < len(m.code) {
		instruction := m.code[m.ip]

		switch instruction {
		case '+':
			m.memory[m.dp]++
		case '-':
			m.memory[m.dp]--
		case '>':
			m.dp++
		case '<':
			m.dp--
		case ',':
			m.readChar()
		case '.':
			m.putChar()
		}

		m.ip++

	}
}

func (m *Machine) readChar() {
	n, err := m.input.Read(m.buf)
	if err != nil {
		panic(err)
	}
	if n != 1 {
		panic("wrong num bytes read")
	}

	m.memory[m.dp] = int(m.buf[0])
}

func (m *Machine) putChar() {
	m.buf[0] = byte(m.memory[m.dp])

	n, err := m.output.Write(m.buf)
	if err != nil {
		panic(err)
	}
	if n != 1 {
		panic("wrong num bytes written")
	}
}
