package loads

import (
	"waveJVM/instructions/base"
	"waveJVM/rtda"
)

/**
 * @author DingZhenJie
 * @date 2024/7/21
 * @desc: 把对应局部变量表上的double加载到stack中
 * If the project has helped you, please make sure to give me a star
 **/

type DLOAD struct {
	base.Index8Instruction
	//	仅仅存放该对象引用的索引位置
}

func (self *DLOAD) Execute(frame *rtda.Frame) {
	_dload(frame, self.Index)
}

type DLOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *DLOAD_0) Execute(frame *rtda.Frame) {
	_dload(frame, 0)
}

type DLOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *DLOAD_1) Execute(frame *rtda.Frame) {
	_dload(frame, 1)
}

type DLOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *DLOAD_2) Execute(frame *rtda.Frame) {
	_dload(frame, 2)
}

type DLOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *DLOAD_3) Execute(frame *rtda.Frame) {
	_dload(frame, 3)
}

func _dload(frame *rtda.Frame, index uint) {
	dval := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(dval)
}
