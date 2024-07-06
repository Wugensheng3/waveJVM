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

type ASTORE struct {
	base.Index8Instruction
}

func _astore(frame *rtda.Frame, index uint) {
	ref := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, ref)
}
func (self *ASTORE) Execute(frame *rtda.Frame) {
	_astore(frame, uint(self.Index))
}

type ASTORE0 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE0) Execute(frame *rtda.Frame) {
	_astore(frame, 0)
}

type ASTORE1 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE1) Execute(frame *rtda.Frame) {
	_astore(frame, 1)
}

type ASTORE2 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE2) Execute(frame *rtda.Frame) {
	_astore(frame, 2)
}

type ASTORE3 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE3) Execute(frame *rtda.Frame) {
	_astore(frame, 3)
}
