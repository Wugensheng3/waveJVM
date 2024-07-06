package base

import "waveJVM/rtda"

/**
 * @author DingZhenJie
 * @date 2024/7/22
 * @desc: 实现跳转相关的函数
 * If the project has helped you, please make sure to give me a star
 **/
func Branch(frame *rtda.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
