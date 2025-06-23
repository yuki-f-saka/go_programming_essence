package main

import "fmt"

// interfaceについて
// メソッドを持つ型のインターフェースを定義できる

/*
Goでは型そのものがインターフェースSpeakerを実装していることを明示しなくてもいい設計をしている
これをダック・タイピングという
*/
type Speaker interface {
	Speak() error
}

var _ Speaker = (*Dog)(nil)

type Dog struct{}

func (d *Dog) Speak() error {
	fmt.Println("Bowwow")
	return nil
}

type Cat struct{}

func (c *Cat) Speak() error {
	fmt.Println("Nyan")
	return nil
}

// メソッドSpeakを持つ型DogとCatは値はSpeaker型の変数に代入できる
func DoSpeak(s Speaker) error {
	return s.Speak()
}

func main() {
	// dog := Dog{}
	// DoSpeak(&dog)

	cat := Cat{}
	DoSpeak(&cat)
}
