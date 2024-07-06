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

type LSTORE struct {
	base.Index8Instruction
}

func _lstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}
func (self *LSTORE) Execute(frame *rtda.Frame) {
	_lstore(frame, uint(self.Index))
}

type LSTORE0 struct {
	base.NoOperandsInstruction
}

func (self *LSTORE0) Execute(frame *rtda.Frame) {
	_lstore(frame, 0)
}

type LSTORE1 struct {
	base.NoOperandsInstruction
}

func (self *LSTORE1) Execute(frame *rtda.Frame) {
	_lstore(frame, 1)
}

type LSTORE2 struct {
	base.NoOperandsInstruction
}

func (self *LSTORE2) Execute(frame *rtda.Frame) {
	_lstore(frame, 2)
}

type LSTORE3 struct {
	base.NoOperandsInstruction
}

func (self *LSTORE3) Execute(frame *rtda.Frame) {
	_lstore(frame, 3)
}
