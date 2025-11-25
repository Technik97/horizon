package eval

type Value interface{}

type NumberVal float64
type SymbolVal string

// A built-in function like +, -, *
type BuiltinFunc func(args []Value) (Value, error)

type Env struct {
	store map[string]Value
	outer *Env
}

func New() *Env {
	return &Env{store: make(map[string]Value)}
}

func (e *Env) Get(name string) (Value, bool) {
	v, ok := e.store[name]
	if !ok && e.outer != nil {
		return e.outer.Get(name)
	}

	return v, ok
}

func (e *Env) Set(name string, val Value) Value {
	e.store[name] = val
	return val
}
