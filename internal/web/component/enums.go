package component

import "github.com/orsinium-labs/enum"

type color enum.Member[string]
type size enum.Member[string]
type buttonType enum.Member[string]

var (
	ColorBlack = color{"black"}

	SizeSm = size{"sm"}
	SizeMd = size{"md"}
	SizeLg = size{"lg"}

	ButtonTypeSubmit = buttonType{"submit"}
	ButtonTypeButton = buttonType{"button"}
	ButtonTypeReset  = buttonType{"reset"}
)
