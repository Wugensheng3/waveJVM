package comparisons

import (
	"waveJVM/instructions/base"
	"waveJVM/rtda"
)

/**
 * @author DingZhenJie
 * @date 2024/7/22
 * @desc: 第一类(如此概念不清楚,看lcmp),用于比较float变量
 *			与long不同的点在于浮点数计算会产生NaN,也就是无法比较,下面做出具体补充
 * @supplement:NaN通常出现在浮点数运算中的一些特殊情况下，例如除以0的结果或者对负数开平方根。当这些操作无法产生一个合法的浮点数时，Java会返回NaN。
 * 				不过最终实现上java的底层代码isNaN应该是利用计算机的性质,判断当前的值不等于他本身,来作为NaN,最终会落到一个值上去
 *				而我们通常是直接利用go的sdk来写代码,所以我们需要特别判断一下NaN的情况,例如说要默认给他true,还是false?
 *				所以我们有下面两个指令来针对这一情况
 * If the project has helped you, please make sure to give me a star
 **/

type FCMPG struct {
	base.NoOperandsInstruction
}

func (self *FCMPG) Execute(frame *rtda.Frame) {
	_fcmp(frame, true)
}

type FCMPL struct {
	base.NoOperandsInstruction
}

func (self *FCMPL) Execute(frame *rtda.Frame) {
	_fcmp(frame, false)
}
func _fcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
