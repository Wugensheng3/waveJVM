package stores

import (
	"waveJVM/instructions/base"
	"waveJVM/rtda"
)

/**
 * @author DingZhenJie
 * @date 2024/7/21
 * @desc: 实现的功能与加载指令相反(把变量从操作数栈顶弹出,存入局部变量表)
 * 			同样数字后缀仅仅是直接放到那个数字的下标处
 * If the project has helped you, please make sure to give me a star
 **/

type DSTORE struct {
	base.Index8Instruction
}

func _dstore(frame *rtda.Frame, index uint) {
	dval := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, dval)
}
func (self *DSTORE) Execute(frame *rtda.Frame) {
	_dstore(frame, uint(self.Index))
}

type DSTORE0 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE0) Execute(frame *rtda.Frame) {
	_dstore(frame, 0)
}

type DSTORE1 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE1) Execute(frame *rtda.Frame) {
	_dstore(frame, 1)
}

type DSTORE2 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE2) Execute(frame *rtda.Frame) {
	_dstore(frame, 2)
}

type DSTORE3 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE3) Execute(frame *rtda.Frame) {
	_dstore(frame, 3)
}
