package eval

import (
	"fmt"
	"strings"
)

func (v Var) String() string {
	return string(v)
}

func (l Literal) String() string {
	return fmt.Sprintf("%g", l)
}

func (u unary) String() string  {
	return fmt.Sprintf("(%c%s)", u.op, u.x.String())
}

func (b binary) String() string  {
	return fmt.Sprintf("(%s %c %s)", b.x.String(), b.op, b.y.String())
}

func (c call) String() string  {
	var args []string
	for _, arg := range c.args {
		args = append(args, arg.String())
	}
	return fmt.Sprintf("%s(%s)", c.fn, strings.Join(args, ", "))
}

func (m Minimum) String() string {
    var args []string
    for _, arg := range m.Args {
        args = append(args, arg.String())
    }
    return fmt.Sprintf("min(%s)", strings.Join(args, ", "))
}