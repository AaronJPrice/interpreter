package object

import "fmt"

type Object interface {
	Type() ObjectType
	Inspect() string
}

type ObjectType string

const (
	NULL_OBJ    ObjectType = "NULL"
	BOOLEAN_OBJ ObjectType = "BOOLEAN"
	INTEGER_OBJ ObjectType = "INTEGER"
)

type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
