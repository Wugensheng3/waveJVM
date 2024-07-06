package base

import "waveJVM/rtda"

/**
 * @author DingZhenJie
 * @date 2024/7/9
 * @desc: 最基础的指令(相当于一个父类)
 * If the project has helped you, please make sure to give me a star
 **/

type Instruction interface {
	//从字节码中提取操作数
	FetchOperands(reader *BytecodeReader)
	//执行指令逻辑
	Execute(frame *rtda.Frame)
}

// NoOperandsInstruction表示没有操作数的指令，所以没有定义任何字段。
type NoOperandsInstruction struct{}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {

}

// 跳转指令
type BranchInstruction struct {
	//跳转偏移量
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	//跳转指令的offset是16
	self.Offset = int(reader.ReadInt16())
}

// 这个作为需要用到索引的指令的父类,且索引是单字节操作数也就是8位
// 这里的索引主要是针对通过索引来存取局部变量表

type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

// 一些指令需要访问运行时常量池,而运行时常量池的索引是由2字节操作数给的,所以读取16
type Index16Instruction struct {
	Index uint
}
