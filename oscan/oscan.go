package oscan

type tLex = int
type tName = [NameLen]uint16

const (
	NameLen   = 31
	lexNone   = 0
	lexName   = 1
	lexNum    = 2
	lexMODULE = 3
	lexIMPORT = 4
	lexBEGIN  = 5
	lexEND    = 6
	lexCONST  = 7
	lexVAR    = 8
	lexWHILE  = 9
	lexDO     = 10
	lexIF     = 11
	lexTHEN   = 12
	lexELSIF  = 13
	lexELSE   = 14
	lexMult   = 15
	lexDIV    = 16
	lexMOD    = 17
	lexPlus   = 18
	lexMinus  = 19
	lexEQ     = 20
	lexNE     = 21
	lexLT     = 22
	lexLE     = 23
	lexGT     = 24
	lexGE     = 25
	lexDot    = 26
	lexComma  = 27
	lexColon  = 28
	lexSemi   = 29
	lexAss    = 30
	lexLpar   = 31
	lexRpar   = 32
	lexEOT    = 33
)

var (
	Lex    tLex
	Name   tName
	Num    int
	LexPos int
)

func InitScan() {

}

func NextLex() {

}
