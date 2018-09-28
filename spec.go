package nipap

type Spec map[string]interface{}

func (spec Spec) Set(key string, value interface{}) {
	spec[key] = value
}

func NewSpec() Spec {
	return make(Spec, 0)
}
