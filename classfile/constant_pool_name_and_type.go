package classfile

/**
 * @author DingZhenJie
 * @date 2024/7/6
 * @desc: CONSTANT_NameAndType_info
 *	CONSTANT_Class_info和CONSTANT_NameAndType_info加在一起可以唯一确定一个字段或者方法
 *  为了支持重载在这之中要同时包含名称和描述符
 *  描述符,略有不同,是对应类型的一次转变
 * If the project has helped you, please make sure to give me a star
 **/
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}
