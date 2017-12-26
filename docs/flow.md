only install command dependencies when actually needed

implement clean command where it uninstalls extensions

```
zeal extensions remove install
zeal remove install extensions
```

types of data to be stored?
 - file (always global)
 - package state (has separate db per package install)

How to handle update commands?
 - on update only update commands are called.

Event hooks is still important.
 - preInstall
 - postInstall
 - preStart
 - postStart
 - preUninstall
 - postUninstall

### NEW EVENT STYLE

Since injection of config is allowed. Event
hook can be pre/post of given key like

```
//For command specific event
pre install
post install

//For extension specific event
pre install.content
post install.content

//For specific config
pre install.content.service
post install.content.service
```

Extensions Initialization

```
./ext/path/extExecutable {RPC_PORT} {EXT_KEY}
```

RPC API

Use RPC API to set variables, hooks, log, and files information.

Every extension will have its own execution ID that will be used when
communicating with zeal.

Variable
- addVariable
(- removeVariable)
(- getVariable)

Files
- addFile
- removeFile
(- getFileInfo)
(- searchFiles)

Event
- registerToEvent(eventName)

- addDependency(event)

- addConfig(configOptions) // At end
- addConfigBefore(configOptions, configKey) - insert before config key
- addConfigAfter(configOptions, configKey)
- getConfigs() - Used to filter available configs

configKey can be command, command.extension, and command.extension.item
this way you can easily specify where if you dont care about
specific config

Log
- logInfo
- logDebug
- logError

Extensions RPC

processEvent
    - eventName
    - packageName
    - version
    - prevVersion (Only available when replacing an old version)
    - options (extension input based on config)


LOCAL EXTENSIONS HANDLER

package settings

package content

anycommand option

anycommand dependencies

ZEAL State

local state for init/build/test/pack

global state for install/uninstall/stop/start/verify