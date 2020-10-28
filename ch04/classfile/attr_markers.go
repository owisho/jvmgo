package classfile

type DeprecatedAttribute struct {
	MarkerAttribute
}

type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct {
}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	//read nothing 因为这两个属性都没有数据，所以readerInfo()方法是空的
}
