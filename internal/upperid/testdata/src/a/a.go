package a

func f() {
	var ID int  // IDになってるのでOK
	println(ID) // IDになってるのでOK

	var Id int  // want "Id must be ID"
	println(Id) // want "Id must be ID"
}
