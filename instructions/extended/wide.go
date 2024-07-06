package extended

import (
	"waveJVM/instructions/base"
	"waveJVM/instructions/loads"
	"waveJVM/instructions/math"
	"waveJVM/instructions/stores"
	"waveJVM/rtda"
)

/**
 * @author DingZhenJie
 * @date 2024/7/30
 * @desc: 我们索引只用uint8来表示,为了避免局部变量表的大小超过了256,于是产生了这个命令,用于改变其他指令的行为
 * If the project has helped you, please make sure to give me a star
 **/
type WIDE struct {
	modifiedInstruction base.Instruction //存放被该表的指令
}

func (self *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	//只关于我们那些需要涉及用到索引的命令
	switch opcode {
	case 0x15:
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = inst
	case 0x16: //lload
		inst := &loads.LLOAD{}
		inst.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = inst
	case 0x17: //fload
		inst := &loads.FLOAD{}
		inst.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = inst
	case 0x18: //dload
		inst := &loads.DLOAD{}
		inst.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = inst
	case 0x19: //aload
		inst := &loads.ALOAD{}
		inst.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = inst
	case 0x36: //istore
		inst := &stores.ISTORE{}
		inst.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = inst
	case 0x37: //lstore
		inst := &stores.LSTORE{}
		inst.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = inst
	case 0x38: //fstore
		inst := &stores.FSTORE{}
		inst.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = inst
	case 0x39: //dstore
		inst := &stores.DSTORE{}
		inst.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = inst
	case 0x3a: //astore
		inst := &stores.ASTORE{}
		inst.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = inst
	case 0x84:
		inst := &math.IINC{}
		inst.Index = uint(reader.ReadUint16())
		inst.Const = int32(reader.ReadInt16())
		self.modifiedInstruction = inst
	case 0xa9:
		panic("Unsupported opcode:0xa9!")
	}
}
func (self *WIDE) Execute(frame *rtda.Frame) {
	self.modifiedInstruction.Execute(frame)
}
