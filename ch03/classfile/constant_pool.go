package classfile

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUnit16())
	cp:= make([]ConstantInfo, cpCount)

	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader,cp)
		switch cp[i].(type){
		case *ConstantLongInfo,*ConstantDoubleInfo:
			i++
		}
	}
}

func (self *ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index];cpInfo!=nil{
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

func (self *ConstantPool) getNameAndType(index uint16) (string,string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name,_type
}

func (self *ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

func (self *ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}

type ConstantInfo interface{
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader,cp ConstantPool) ConstantInfo {
	tag:=reader.readUnit8()
	c:=newConstantInfo(tag,cp)
	c.readInfo(reader)
	return c
}

func newConstantInfo(tag uint8,cp ConstantPool) ConstantInfo{
	switch tag{
	case CONSTANT_Integer: return &ConstantIntegerInfo{}
	case CONSTANT_FLoat: return &ConstantFloatInfo{}
	}
}