package model

type Plasticine struct {
	A string
	B string
	C string
	D string
	E string
	F string
	G string
	H string
	I string
	J string
	K string `gorm:"type:varchar(500)"`
	L string
	M string
	N string
	O string
	P string
	Q string
	R string
	S string
	T string
	U string
	V string
	W string
	X string
	Y string
	Z string
}

var ColumnNameMaping []string

func (p Plasticine)TableName() string {
	return "plasticine"
}
