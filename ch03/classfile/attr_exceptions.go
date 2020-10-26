package classfile

type ExceptionsAttributes struct {
	exceptionIndexTable []uint16
}

func (self *ExceptionsAttributes) readInfo(reader *ClassReader) {
	self.exceptionIndexTable = reader.readUint16s()
}

func (self *ExceptionsAttributes) ExceptionIndexTable() []uint16 {
	return self.exceptionIndexTable
}