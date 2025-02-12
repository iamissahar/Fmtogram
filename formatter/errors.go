package formatter

import "github.com/l1qwie/Fmtogram/fmerrors"

func code01() error {
	err := new(fmerrors.FME)
	err.Code = 1
	err.String = "there isnt a place to put a data in"
	return err
}

func code3() error {
	err := new(fmerrors.FME)
	err.Code = 3
	err.String = "media holder is full"
	return err
}

func code5() error {
	err := new(fmerrors.FME)
	err.Code = 5
	err.String = "incorrect data in an input slice"
	return err
}

func code07() error {
	err := new(fmerrors.FME)
	err.Code = 7
	err.String = "method isn't available"
	return err
}

func code10() error {
	err := new(fmerrors.FME)
	err.Code = 10
	err.String = "the data is already present"
	return err
}

func code12() error {
	err := new(fmerrors.FME)
	err.Code = 12
	err.String = "incorrect type of file"
	return err
}

func code20() error {
	err := new(fmerrors.FME)
	err.Code = 20
	err.String = "incorrect input data"
	return err
}

func code21() error {
	err := new(fmerrors.FME)
	err.Code = 21
	err.String = "missed required data"
	return err
}

func code22() error {
	err := new(fmerrors.FME)
	err.Code = 22
	err.String = "got an error from Telegram"
	return err
}

func code25() error {
	err := new(fmerrors.FME)
	err.Code = 25
	err.String = "incompatible data"
	return err
}

func code54() error {
	err := new(fmerrors.FME)
	err.Code = 54
	err.String = "isn't allowed to be united with others"
	return err
}
