package math

import (
	"waveJVM/instructions/base"
	"waveJVM/rtda"
)

/**
 * @author DingZhenJie
 * @date 2024/7/22
 * @desc: 位移指令(左移,右移(有符号右移(有符号补符号位),无符号右移(无符号的话,会补0)))
 *			有符号又称算术右移(其实现上是:保留符号位,同时新增的位以符号位为准)
 * @summary: 无论有无符号右移其实都一样,我们这里的实现基本上是利用了go语言,go的话,如果是本身是uint那么他就会是无符号右移,如果是int的话,那么他就会是有符号右移了
 * 			然后需要注意的一点是:偏移量上,必须是无符号整数
 * If the project has helped you, please make sure to give me a star
 **/
//int 左移

type ISHL struct {
	base.NoOperandsInstruction
}

func (self *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	//v2代表左移多少位
	v2 := stack.PopInt()
	//v1是要进行位移操作的变量
	v1 := stack.PopInt()
	//int 只有32位,我们顶多移32位,之后再移就是全0了,所以我们相当于一个mod操作
	//剩下最后的5个1即可
	//这里待定,左移=32的话,这里类似不变动的情况,也就是不移动
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}

//int 有符号右移

type ISHR struct {
	base.NoOperandsInstruction
}

func (self *ISHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 >> s
	stack.PushInt(result)
}

//int 无符号右移

type IUSHR struct {
	base.NoOperandsInstruction
}

func (self *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32((uint32(v1) >> s))
	stack.PushInt(result)
}

//long 左移

type LSHL struct {
	base.NoOperandsInstruction
}

func (self *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := (v1) << s
	stack.PushLong(result)
}

//long 有符号右移

type LSHR struct {
	base.NoOperandsInstruction
}

func (self *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	//由于long有64位所以取3f,6位
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}

//long 无符号右移

type LUSHR struct {
	base.NoOperandsInstruction
}

func (self *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	//由于long有64位所以取3f,6位
	s := uint32(v2) & 0x3f
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}
