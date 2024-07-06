package math

import (
	"waveJVM/instructions/base"
	"waveJVM/rtda"
)

/**
 * @author DingZhenJie
 * @date 2024/7/22
 * @desc: 专门给int变量增加常量的命令
 * If the project has helped you, please make sure to give me a star
 **/
//这里就懒于继承了Index8Operation了,那个主要是针对只有Index的情况

type IINC struct {
	Index uint
	Const int32
}

func (self *IINC) FetchOperands(reader *base.BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
	self.Const = int32(reader.ReadInt8())
}

//注释:这里犯了点迷糊,感想之前的自己弄了篇笔记,导致i++与++i的区别
//主要是在编译器的编译成class文件上,与我们jvm关系不大

func (self *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(self.Index)
	val += self.Const
	localVars.SetInt(self.Index, val)
}
