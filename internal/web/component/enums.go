package component

import "github.com/orsinium-labs/enum"

type color enum.Member[string]
type size enum.Member[string]
type buttonType enum.Member[string]
type inputType enum.Member[string]

var (
	ColorBlack = color{"black"}
	ColorRed   = color{"red"}

	SizeSm = size{"sm"}
	SizeMd = size{"md"}
	SizeLg = size{"lg"}

	ButtonTypeSubmit = buttonType{"submit"}
	ButtonTypeButton = buttonType{"button"}
	ButtonTypeReset  = buttonType{"reset"}

	InputTypeText     = inputType{"text"}
	InputTypePassword = inputType{"password"}
	InputTypeEmail    = inputType{"email"}
	InputTypeNumber   = inputType{"number"}
	InputTypeTel      = inputType{"tel"}
	InputTypeUrl      = inputType{"url"}
)
