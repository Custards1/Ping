package interpreter

type PingType byte

type pinger interface{}
type PingPtr *pinger

//Pong is a Pointer to a PingObj
type Pong *PingObj

const (
	TypeNil            = PingType(iota)
	TypeNumber         = PingType(iota)
	TypeString         = PingType(iota)
	TypeBool           = PingType(iota)
	TypeStruct         = PingType(iota)
	TypeFunctionNative = PingType(iota)
	TypeFunctionC      = PingType(iota)
)

type PingObj struct {
	obj   PingPtr
	otype PingType
	arg   *PingObj
}

func PingNew(ptype PingType, data PingPtr) *PingObj {
	p := PingObj{data, ptype, nil}
	return &p
}
func (p *PingObj) SetNumber(i pinger) {
	p.otype = TypeNumber
	c := i
	p.obj = &c
}
func (p *PingObj) SetBool(i pinger) {
	p.otype = TypeBool
	c := i
	p.obj = &c
}
func (p *PingObj) SetNil() {
	p.otype = TypeNil
	var c PingPtr
	c = nil
	p.obj = c
}
func (p *PingObj) SetStruct(pa PingPtr) {
	p.otype = TypeStruct
	p.obj = pa
}
func (p *PingObj) SetString(pa PingPtr) {
	p.otype = TypeString
	p.obj = pa
}
func (p *PingObj) SetNativeFunction(pa PingPtr) {
	p.otype = TypeFunctionNative
	p.obj = pa
}
func (p *PingObj) SetCFunction(pa PingPtr) {
	p.otype = TypeFunctionC
	p.obj = pa
}
