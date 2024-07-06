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

type FSTORE struct {
	base.Index8Instruction
}

func _fstore(frame *rtda.Frame, index uint) {
	fval := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, fval)
}
func (self *FSTORE) Execute(frame *rtda.Frame) {
	_fstore(frame, uint(self.Index))
}

type FSTORE0 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE0) Execute(frame *rtda.Frame) {
	_fstore(frame, 0)
}

type FSTORE1 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE1) Execute(frame *rtda.Frame) {
	_fstore(frame, 1)
}

type FSTORE2 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE2) Execute(frame *rtda.Frame) {
	_fstore(frame, 2)
}

type FSTORE3 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE3) Execute(frame *rtda.Frame) {
	_fstore(frame, 3)
}
