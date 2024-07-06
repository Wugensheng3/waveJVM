package base

/**
 * @author DingZhenJie
 * @date 2024/7/9
 * @desc: 用来从字节码中读取指定长度的字节,也就是提取字节码信息的作用
 * If the project has helped you, please make sure to give me a star
 **/

type BytecodeReader struct {
	code []byte
	pc   int
	//这里的pc不是程序计数器啦,而是确定我们当前读到了哪个字节
}

func (self *BytecodeReader) PC() int {
	return self.pc
}
func (self *BytecodeReader) SkipPadding() {
	for self.pc%4 != 0 { //当 self.pc%4 != 0 时，说明当前指针位置未对齐到 4 的倍数。
		self.ReadUint8() //跳一个字节&&将指针位置加 1
	}
}
func (self *BytecodeReader) ReadInt32s(n int32) []int32 {
	ints := make([]int32, n)
	for i := range ints {
		ints[i] = self.ReadInt32()
	}
	return ints
}
func (self *BytecodeReader) Reset(code []byte, pc int) {
	self.code = code
	self.pc = pc
}
func (self *BytecodeReader) ReadUint8() uint8 {
	i := self.code[self.pc]
	self.pc++
	return i
}
func (self *BytecodeReader) ReadInt8() int8 {
	return int8(self.ReadUint8())
}

func (self *BytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(self.ReadUint8())
	byte2 := uint16(self.ReadUint8())
	return (byte1 << 8) | byte2
}
func (self *BytecodeReader) ReadInt16() int16 {
	return int16(self.ReadUint16())
}
func (self *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(self.ReadUint8())
	byte2 := int32(self.ReadUint8())
	byte3 := int32(self.ReadUint8())
	byte4 := int32(self.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}
