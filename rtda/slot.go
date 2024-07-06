package rtda

/**
 * @author DingZhenJie
 * @date 2024/7/8
 * @desc: 之所以采用slot,是因为go对于直接使用uintptr指针会在没有其他地方引用这个实例时,会被当垃圾回收
 * If the project has helped you, please make sure to give me a star
 **/

type Slot struct {
	num int32
	ref *Object
}
