package classfile

/**
 * @author DingZhenJie
 * @date 2024/7/6
 * @desc: SourceFile属性
 * If the project has helped you, please make sure to give me a star
 **/
type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (self *SourceFileAttribute) readInfo(reader *ClassReader) {
	self.sourceFileIndex = reader.readUint16()
}
func (self *SourceFileAttribute) FileName() string {
	return self.cp.getUtf8(self.sourceFileIndex)
}
