package classfile

/**
 * @author DingZhenJie
 * @date 2024/7/6
 * @desc: Exceptions是变长属性，记录方法抛出的异常表
 * If the project has helped you, please make sure to give me a star
 **/
type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (self *ExceptionsAttribute) readInfo(reader *ClassReader) {
	self.exceptionIndexTable = reader.readUint16s()
}
func (self *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return self.exceptionIndexTable
}
