package classfile

/**
 * @author DingZhenJie
 * @date 2024/7/6
 * @desc: UnparsedAttribute结构体
 * If the project has helped you, please make sure to give me a star
 **/
type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (self *UnparsedAttribute) readInfo(reader *ClassReader) {
	self.info = reader.readBytes(self.length)
}
