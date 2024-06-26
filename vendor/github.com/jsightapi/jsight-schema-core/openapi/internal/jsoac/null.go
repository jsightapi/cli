package jsoac

import (
	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/openapi/internal"
)

type Null struct {
	Example     *Example     `json:"example,omitempty"`
	Enum        *Enum        `json:"enum,omitempty"`
	Nullable    *Nullable    `json:"nullable,omitempty"`
	Description *Description `json:"description,omitempty"`
}

var _ Node = (*Null)(nil)

func newNull(astNode schema.ASTNode) *Null {
	return &Null{
		Example:     newExample(internal.StringNull, false),
		Enum:        newEnum(astNode),
		Nullable:    newNullable(astNode),
		Description: newDescription(astNode),
	}
}

func (n *Null) SetNodeDescription(s string) {
	n.Description = newDescriptionFromString(s)
}
