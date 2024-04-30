package eval

import "bitbucket.org/hurricanecommerce/dev-days/2024-05-09/src/eval/object"

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)
