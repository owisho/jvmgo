package classfile

import (
	"fmt"
	"unicode/utf16"
)

type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	
}