package main

// NewHogeというコンストラクタ関数を用意してあげるのが一般的な習慣
func NewUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}
