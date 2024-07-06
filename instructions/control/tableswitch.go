package control

import (
	"waveJVM/instructions/base"
	"waveJVM/rtda"
)

/**
 * @author DingZhenJie
 * @date 2024/7/29
 * @desc: 用来实现java的Switch-case语句的的方式之一,主要是用在case:0,case1:等连续的情况下,可以直接用索引表来实现
 * If the project has helped you, please make sure to give me a star
 **/

type TABLE_SWITCH struct {
	defaultOffset int32   //记录默认情况下执行跳转所需的字节码偏移量
	low           int32   //case的取值范围min
	high          int32   //case取值范围max
	jumpOffsets   []int32 //table对应的索引表,存放了high-low+1个int值,对应各个case,所需跳转的偏移量
}

func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	//在 JVM 中，tableswitch 指令后面的 padding 是为了确保 defaultOffset 在字节码中的地址是 4 的倍数。这是因为 JVM 使用了 4 字节对齐的策略来提高内存访问的效率和处理器执行速度。
	//之所以需要跳过padding是因为tableswitch指令操作码后面由0~3字节的padding
	//为了保证defaultOffset在字节码中的地址是4的倍数
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetsCount := self.high - self.low + 1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}
func (self *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= self.low && index <= self.high {
		offset = int(self.jumpOffsets[index-self.low])
	} else {
		offset = int(self.defaultOffset)
	}
	base.Branch(frame, offset)
}
