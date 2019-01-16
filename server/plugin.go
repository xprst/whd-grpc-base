package server

import "wh-git.mingyuanyun.com/grpc-go/go-grpc-base/errors"

// Plugin 插件接口
type Plugin interface{}

type (
	// RegisterPlugin is .
	RegisterPlugin interface {
		Register(name string, rcvr interface{}, metadata string) error
		Unregister(name string) error
	}
)

// PluginContainer 插件管理接口
type PluginContainer interface {
	Add(plugin Plugin)    // 添加插件
	Remove(plugin Plugin) // 移除插件
	All() []Plugin        // 所有插件

	DoRegister(name string, rcvr interface{}, metadata string) error
}

// pluginContainer 实现 PluginContainer 接口.
type pluginContainer struct {
	plugins []Plugin
}

func (p *pluginContainer) Add(plugin Plugin) {
	p.plugins = append(p.plugins, plugin)
}

func (p *pluginContainer) Remove(plugin Plugin) {
	if p.plugins == nil {
		return
	}
	var plugins []Plugin
	for _, p := range p.plugins {
		if p != plugin {
			plugins = append(plugins, p)
		}
	}
	p.plugins = plugins
}

func (p *pluginContainer) All() []Plugin {
	return p.plugins
}

// DoRegister invokes DoRegister plugin.
func (p *pluginContainer) DoRegister(name string, rcvr interface{}, metadata string) error {
	var es []error
	for _, rp := range p.plugins {
		if plugin, ok := rp.(RegisterPlugin); ok {
			err := plugin.Register(name, rcvr, metadata)
			if err != nil {
				es = append(es, err)
			}
		}
	}
	if len(es) > 0 {
		return errors.NewMultiError(es)
	}
	return nil
}
