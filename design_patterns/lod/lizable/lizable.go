package lizable

//有依赖关系的类之间，尽量只依赖必要的接口

type Serializable interface {
	serialize(object interface{}) string
}

type Deserializer interface {
	deserializer(object interface{}) string
}

type Serialization struct {

}

func (self *Serialization) serialize(object interface{}) string {
	return ""
}

func (self *Serialization) deserializer(object interface{}) string {
	return ""
}