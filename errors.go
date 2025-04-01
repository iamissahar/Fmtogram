package fmtogram

import (
	"fmt"
	"runtime/debug"
)

type FME struct {
	Code   int
	String string
	Stack  string
}

func (f *FME) Error() string {
	f.Stack = string(debug.Stack())
	return fmt.Sprintf("[ERROR:%d] message: %s\nStack: %s", f.Code, f.String, f.Stack)
}

// func (f *FME) Error() string {
// 	return fmt.Sprintf("%d", f.Code)
// }

func TelegramErrors(code int, dis string) error {
	return fmt.Errorf("[ERROR:%d] %s", code, dis)
}

func code01() error {
	err := new(FME)
	err.Code = 1
	err.String = "there isnt a place to put a data in"
	return err
}

func code3() error {
	err := new(FME)
	err.Code = 3
	err.String = "media holder is full"
	return err
}

func code5() error {
	err := new(FME)
	err.Code = 5
	err.String = "incorrect data in an input slice"
	return err
}

func code07() error {
	err := new(FME)
	err.Code = 7
	err.String = "method isn't available"
	return err
}

func code10() error {
	err := new(FME)
	err.Code = 10
	err.String = "the data is already present"
	return err
}

func code12() error {
	err := new(FME)
	err.Code = 12
	err.String = "incorrect type of file"
	return err
}

func code20() error {
	err := new(FME)
	err.Code = 20
	err.String = "incorrect input data"
	return err
}

func code21() error {
	err := new(FME)
	err.Code = 21
	err.String = "missed required data"
	return err
}

func code22() error {
	err := new(FME)
	err.Code = 22
	err.String = "got an error from Telegram"
	return err
}

func code25() error {
	err := new(FME)
	err.Code = 25
	err.String = "incompatible data"
	return err
}

func code54() error {
	err := new(FME)
	err.Code = 54
	err.String = "isn't allowed to be united with others"
	return err
}
