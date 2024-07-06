package constants

import (
	"waveJVM/instructions/base"
	"waveJVM/rtda"
)

/**
 * @author DingZhenJie
 * @date 2024/7/9
 * @desc: 什么都不做的指令
 * If the project has helped you, please make sure to give me a star
 **/

type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {

}
