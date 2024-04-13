package x

const Hoge = "hoge"

func Exec(s string) string {
	return s
}

func OtherMethod (s string) string {
	return s
}

type St struct {}
func (st *St) Exec(s string) string {
	return s
}