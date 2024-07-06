package stack

import (
	"waveJVM/instructions/base"
	"waveJVM/rtda"
)

/**
 * @author DingZhenJie
 * @date 2024/7/22
 * @desc: pop指令把栈顶变量弹出
 * @distinction: 与我们在operand_stack.go中定义的一系列方法区别在于
 * 				那里的popLong等等情况是需要获取到对应的弹出元素的
 * If the project has helped you, please make sure to give me a star
 **/

//仅仅负责弹出占用一个slot的变量

type POP struct {
	base.NoOperandsInstruction
}

func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

//负责弹出两个slot的变量

type POP2 struct {
	base.NoOperandsInstruction
}

func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
