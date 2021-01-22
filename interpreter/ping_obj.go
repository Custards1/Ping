package interpreter

type PingType byte

//*PingObj is a Pointer to a PingObj
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
	obj      interface{}
	otype    PingType
	arg      *PingObj
	hashid   *string
	refcount int
}

func PingNew(ptype PingType, data interface{}, hashid *string) *PingObj {
	p := PingObj{data, ptype, nil, hashid, 0}
	return &p
}
func (p *PingObj) SetNumber(i interface{}) {
	p.otype = TypeNumber
	c := i
	p.obj = &c
}
func (p *PingObj) SetBool(i interface{}) {
	p.otype = TypeBool
	c := i
	p.obj = &c
}
func (p *PingObj) SetNil() {
	p.otype = TypeNil
	var c interface{}
	c = nil
	p.obj = c
}
func (p *PingObj) SetStruct(pa interface{}) {
	p.otype = TypeStruct
	p.obj = pa
}
func (p *PingObj) SetString(pa interface{}) {
	p.otype = TypeString
	p.obj = pa
}
func (p *PingObj) SetNativeFunction(pa interface{}) {
	p.otype = TypeFunctionNative
	p.obj = pa
}
func (p *PingObj) SetCFunction(pa interface{}) {
	p.otype = TypeFunctionC
	p.obj = pa
}
func (p *PingObj) Dec(o *ObjectStore) {
	p.refcount -= 1
	if p.refcount <= 0 {
		o.Destroy(p.hashid)
	}
}
