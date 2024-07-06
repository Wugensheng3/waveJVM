package rtda

/**
 * @author DingZhenJie
 * @date 2024/7/8
 * @desc: 栈帧
 * If the project has helped you, please make sure to give me a star
 **/

type Frame struct {
	//相当于next,只不过为了适合栈,改成了lower
	lower     *Frame
	localVars LocalVars
	//保存操作数栈指针
	operandStack *OperandStack
	//region 为了实现跳转指令而添加
	thread *Thread
	nextPC int
	//endregion
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
func (self *Frame) Thread() *Thread {
	return self.thread
}
func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}
func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}
