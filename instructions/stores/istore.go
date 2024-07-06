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

type ISTORE struct {
	base.Index8Instruction
}

func _istore(frame *rtda.Frame, index uint) {
	ival := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, ival)
}
func (self *ISTORE) Execute(frame *rtda.Frame) {
	_istore(frame, uint(self.Index))
}

type ISTORE0 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE0) Execute(frame *rtda.Frame) {
	_istore(frame, 0)
}

type ISTORE1 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE1) Execute(frame *rtda.Frame) {
	_istore(frame, 1)
}

type ISTORE2 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE2) Execute(frame *rtda.Frame) {
	_istore(frame, 2)
}

type ISTORE3 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE3) Execute(frame *rtda.Frame) {
	_istore(frame, 3)
}
