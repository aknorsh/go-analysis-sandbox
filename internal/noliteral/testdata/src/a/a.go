package a

import "x"

func OK() {
	_ = x.Exec(x.Hoge) // consts の値を渡しているのでOK

	arg := "hoge"
	_ = x.Exec(arg) // 変数の値を渡しているのでOK

	_ = x.OtherMethod("hoge") // 別の関数なのでOK
	st := x.St{}
	_ = st.Exec("hoge") // メソッドなのでOK
}

func NG() {
	_ = x.Exec("Hoge") // want "Use x.Hoge instead of string literal"
}