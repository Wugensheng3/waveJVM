package comparisons

import (
	"waveJVM/instructions/base"
	"waveJVM/rtda"
)

/**
 * @author DingZhenJie
 * @date 2024/7/22
 * @desc: 比较指令可以分为两类：一类将比较结果推入操作数栈顶，一类根据比较结果跳转(如if-else)
 * 			本部分为第一类,用于比较long变量值
 * 			实现上也就是把我们对应的字节码转为go来执行而已啦
 * If the project has helped you, please make sure to give me a star
 **/

type LCMP struct {
	base.NoOperandsInstruction
}

func (self *LCMP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}
