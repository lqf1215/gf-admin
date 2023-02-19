package dept

import (
	"github.com/SupenBysz/gf-admin-community/utility/enum"
)

type StateEnum enum.IEnumCode[int]

type state struct {
	Normal   StateEnum
	Disabled StateEnum
}

var State = state{
	Normal:   enum.New[StateEnum](0, "正常"),
	Disabled: enum.New[StateEnum](1, "停用"),
}

func (e *state) New(code int, description string) StateEnum {
	if (code&State.Normal.Code()) == State.Normal.Code() ||
		(code&State.Disabled.Code()) == State.Disabled.Code() {
		return enum.New[StateEnum](code, description)
	}
	panic("Dept.State.New: error")
}
