package classfile

/**
 * @author DingZhenJie
 * @date 2024/7/6
 * @desc:ConstantValue是定长属性
 * If the project has helped you, please make sure to give me a star
 **/
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}
func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}
