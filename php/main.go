package main

import (
	"fmt"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr/binary"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php7"
	"github.com/z7zmey/php-parser/visitor"
	"github.com/z7zmey/php-parser/walker"
	"os"
	"strconv"
)

func main() {
	src := []byte(`<? echo "Hello world";echo 2*3+4+6*7;`)

	parser := php7.NewParser(src, "7.4")
	parser.Parse()

	for _, e := range parser.GetErrors() {
		fmt.Println(e)
	}

	rootNode := parser.GetRootNode()
	v := visitor.Dumper{
		Writer:    os.Stdout,
		Indent:    "",
	}
	rootNode.Walk(&v)

	rootNode.Walk(&Interpreter{})
}

type Interpreter struct {
	Stack []interface{}
}

// EnterNode is invoked at every node in hierarchy
func (d *Interpreter) EnterNode(w walker.Walkable) bool {
	switch n := w.(type) {
	case *node.Root:
	case *stmt.Echo:
		n.Exprs[0].Walk(d)
		value := d.pop()
		switch v := value.(type) {
		case string:
			fmt.Println(v[1:len(v)-1])
		default:
			fmt.Println(v)
		}
		return false
	case *scalar.Lnumber:
		i, _ := strconv.Atoi(n.Value)
		d.push(i)
	case *scalar.String:
		d.push(n.Value)
	case *binary.Plus:
		n.Left.Walk(d)
		l := d.pop()
		n.Right.Walk(d)
		r := d.pop()
		d.push(l.(int) + r.(int))
		return false
	case *binary.Mul:
		n.Left.Walk(d)
		l := d.pop()
		n.Right.Walk(d)
		r := d.pop()
		d.push(l.(int) * r.(int))
		return false
	}

	return true
}

func (d *Interpreter) pop() interface{} {
	value := d.Stack[len(d.Stack)-1]
	d.Stack = d.Stack[:len(d.Stack)-1]
	return value
}

func (d *Interpreter) push(value interface{}) {
	d.Stack = append(d.Stack, value)
}

func (d *Interpreter) LeaveNode(n walker.Walkable) {}
func (d *Interpreter) EnterChildNode(key string, w walker.Walkable) {}
func (d *Interpreter) LeaveChildNode(key string, w walker.Walkable) {}
func (d *Interpreter) EnterChildList(key string, w walker.Walkable) {}
func (d *Interpreter) LeaveChildList(key string, w walker.Walkable) {}
