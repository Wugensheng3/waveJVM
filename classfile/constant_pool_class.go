package classfile

/**
 * @author DingZhenJie
 * @date 2024/7/6
 * @desc: 常量池之类信息:主要包括类名索引,会被类和超类索引,以及接口表中的接口索引所指向
 * If the project has helped you, please make sure to give me a star
 **/
type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}
func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
