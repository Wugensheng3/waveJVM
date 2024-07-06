package classfile

/**
 * @author DingZhenJie
 * @date 2024/7/6
 * @desc: LineNumberTable属性表存放方法的行号信息
 * If the project has helped you, please make sure to give me a star
 **/
type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}
type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	self.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range self.lineNumberTable {
		self.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}
