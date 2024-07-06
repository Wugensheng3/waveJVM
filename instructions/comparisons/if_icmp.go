package comparisons

import (
	"waveJVM/instructions/base"
	"waveJVM/rtda"
)

/**
 * @author DingZhenJie
 * @date 2024/7/22
 * @desc: 关于int的比较,把两个int变量弹出,然后进行比较满足条件则跳转
 * 			那么我们不需要简单的icmp?没错
 * 下面是原代码:
 *public static void main(String[] args) {
 *      int a=0;
 *      int b=1;
 *      boolean c=a>b;
 * }
 *编译后的字节码:
 *0 iconst_0
 *1 istore_1
 *2 iconst_1
 *3 istore_2
 *4 iload_1
 *5 iload_2
 *6 if_icmple 13 (+7)
 *9 iconst_1
 *10 goto 14 (+4)
 *13 iconst_0
 *14 istore_3
 *15 return
 * 我们可以发现他使用的是if_icmple,然后如果成功的话(a<b)就跳到对应的13行,读取常量0,然后进行设置到对应的3
 * If the project has helped you, please make sure to give me a star
 **/
//=

type IF_ICMPEQ struct {
	base.BranchInstruction
}

func (self *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 == val2 {
		base.Branch(frame, self.Offset)
	}
}

//!=

type IF_ICMPNE struct {
	base.BranchInstruction
}

func (self *IF_ICMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 != val2 {
		base.Branch(frame, self.Offset)
	}
}

//<

type IF_ICMPLT struct {
	base.BranchInstruction
}

func (self *IF_ICMPLT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 < val2 {
		base.Branch(frame, self.Offset)
	}
}

//<=

type IF_ICMPLE struct{ base.BranchInstruction }

func (self *IF_ICMPLE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 <= val2 {
		base.Branch(frame, self.Offset)
	}
}

//>

type IF_ICMPGT struct {
	base.BranchInstruction
}

func (self *IF_ICMPGT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 > val2 {
		base.Branch(frame, self.Offset)
	}
}

//>=

type IF_ICMPGE struct {
	base.BranchInstruction
}

func (self *IF_ICMPGE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 >= val2 {
		base.Branch(frame, self.Offset)
	}
}
