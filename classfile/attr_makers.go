package classfile

/**
 * @author DingZhenJie
 * @date 2024/7/6
 * @desc: 一些属性
 * If the project has helped you, please make sure to give me a star
 **/
type DeprecatedAttribute struct{ MarkerAttribute }
type SyntheticAttribute struct{ MarkerAttribute }
type MarkerAttribute struct{}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	//这两个属性没有数据,所以readInfo方法是空的
}
