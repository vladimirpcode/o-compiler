package otext

/*Драйвер исходного текста*/
//новый коммент
//еще один коммент
const (
	chSpace = ' '
	chTab   = byte(9)
	chEOL   = byte(10)
	chEOT   = byte(0)
)

var (
	Ch   byte
	Line int
	Pos  int
)

func ResetText() {

}

func CloseText() {

}

func NextCh() {

}
