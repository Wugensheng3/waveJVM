package classfile

import "encoding/binary"

/**
 * @author DingZhenJie
 * @date 2024/7/6
 * @desc: 由于class文件类似于结构体,且包含u1,u2,u4四种类型,所以定义此结构体来方便读取
 ** If the project has helped you, please make sure to give me a star
 * */
type ClassReader struct {
	data []byte
}

// jvm中定义了u1,u2,u4四种类型来表示1,2,4字节无符号整数
// 解析u1
func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

// 解析u2
func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:] //这里之所以是跳过前两个byte的切片是因为获得了u2长度的信息,然后我们直接对
	return val
}

// 解析u4
func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}
func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

// 首先n第一次调用获取到后面表项的长度
// 然后不断遍历即封装成s
func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}
func (self *ClassReader) readBytes(length uint32) []byte {
	bytes := self.data[:length]
	self.data = self.data[length:]
	return bytes
}
