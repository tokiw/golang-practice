package eval

import (
	"fmt"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprint(float64(l))
}

func (u unary) String() string {
	return string(u.op) + u.x.String()
}

func (b binary) String() string {
	return b.x.String() + string(b.op) + b.y.String()
}

func (c call) String() string {
	str := c.fn + "("

	for i, arg := range c.args {
		str += arg.String()

		if i != len(c.args)-1 {
			str += ", "
		}
	}

	return str + ")"
}

func (m min) String() string {
	s := "min("
	for i, arg := range m.args {
		s += arg.String()
		if i != len(m.args)-1 {
			s += ", "
		}
	}
	s += ")"
	return s
}
