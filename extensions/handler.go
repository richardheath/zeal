package extensions

import (
	"github.com/richardheath/zeal/config"
)

// LocalExtension Extensions that doesn't use RPC 
type LocalExtension interface {
	Execute(item config.Item) error
}

// ExtensionInstaller Install extensions required by config items
type ExtensionInstaller interface {
	InstallExtensions(items []config.Item) error
}

// ExtensionHandler Handler for extension execution and result.
type ExtensionHandler interface {
	Init(metadataHandler interface{})
	RegisterLocalExtension(localExtension LocalExtension)
	ExecuteConfig(item config.Item) error
}

/*
this will handle extensions rpc

local extensions - extensions that are implemented locally

rpc extensions - extensions that uses rpc to communicate

handle events
*/