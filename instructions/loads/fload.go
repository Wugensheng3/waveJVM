package loads

import (
	"waveJVM/instructions/base"
	"waveJVM/rtda"
)

/**
 * @author DingZhenJie
 * @date 2024/7/21
 * @desc: 把对应局部变量表上的float加载到stack中
 * If the project has helped you, please make sure to give me a star
 **/

type FLOAD struct {
	base.Index8Instruction
	//	仅仅存放该对象引用的索引位置
}

func (self *FLOAD) Execute(frame *rtda.Frame) {
	_fload(frame, self.Index)
}

type FLOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *FLOAD_0) Execute(frame *rtda.Frame) {
	_fload(frame, 0)
}

type FLOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *FLOAD_1) Execute(frame *rtda.Frame) {
	_fload(frame, 1)
}

type FLOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *FLOAD_2) Execute(frame *rtda.Frame) {
	_fload(frame, 2)
}

type FLOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *FLOAD_3) Execute(frame *rtda.Frame) {
	_fload(frame, 3)
}

func _fload(frame *rtda.Frame, index uint) {
	fval := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(fval)
}
