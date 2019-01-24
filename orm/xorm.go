package orm

import "github.com/go-xorm/xorm"


type MyBean struct {
	mapxormengine map[string]*xorm.Engine
}

func NewBean() *MyBean {
	bean := new(MyBean)
	bean.mapxormengine = make(map[string]*xorm.Engine)

	return bean
}

// RegisterEngine 设置orm引擎
func (b *MyBean)RegisterEngine(key string, e *xorm.Engine) {
	b.mapxormengine[key] = e
}

// GetEngine 获取orm引擎
func (b *MyBean)GetEngine(keys ...string) (e *xorm.Engine) {
	if len(keys) == 0 {
		return b.mapxormengine["default"]
	} else {
		return b.mapxormengine[keys[0]]
	}
}