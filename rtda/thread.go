package rtda

/**
 * @author DingZhenJie
 * @date 2024/7/8
 * @desc: 线程相关内容,相当于线程中有栈,栈中有栈帧而每个栈帧约等于一个方法调用
 * If the project has helped you, please make sure to give me a star
 **/

type Thread struct {
	pc    int
	stack *Stack
}

func (self *Thread) NewFrame(maxLocals, maxStack uint) *Frame {
	return newFrame(self, maxLocals, maxStack)

}
func (self *Thread) PC() int {
	return self.pc
}
func (self *Thread) SetPC(pc int) {
	self.pc = pc
}
func NewThread() *Thread {
	return &Thread{
		//1024表示最多能容纳1024个栈帧
		stack: newStack(1024),
	}
}
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}
func (self *Thread) CurrentFrame() *Frame {
	//栈尾即是当前正在运行的帧
	return self.stack.top()
}
