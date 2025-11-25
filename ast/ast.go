package ast

type SExpr interface{}

type Symbol struct {
	Value string
}

type Number struct {
	Value string
}

type List struct {
	Elements []SExpr
}

type Quote struct {
	Value SExpr
}
