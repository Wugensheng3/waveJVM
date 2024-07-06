package classfile

/**
 * @author DingZhenJie
 * @date 2024/7/6
 * @desc: string字面量(注意可不是直接存string哦)
 ** If the project has helped you, please make sure to give me a star
 * */
type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
	//	这里存的index而不直接存字符串是因为之后要通过
	//  string这个来把他们加载到jvmString常量池
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}
func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}
