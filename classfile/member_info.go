package classfile

/**
 * @author DingZhenJie
 * @date 2024/7/6
 * @desc: 统一的struct用来表示类,字段,方法
 ** If the project has helped you, please make sure to give me a star
 * */
type MemberInfo struct {
	cp              ConstantPool //常量池指针,因为每个class文件都有自己的常量池,保存后之后可以方便读取
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

// 读取字段表或方法表
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

// 读取字段或方法数据
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}
func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}
