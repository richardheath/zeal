Zeal
====

Zeal config is designed to facilitate lifecycle of packages.

## Package Server

Zeal server store metadata and actual files separately so they can be capture separately.

State file can be stored locally or stored on a state server. This is
useful to keep state of serverless deployments?

## Configuration can be separated

Configs can be separated on multiple files.

## Reserved extension names

- dependencies
- variables
- settings

## Package Contents/Use

Package items are stored on package. You can pick and choose 
items needed for install.

## Limit commands
- build
- package
- install
    - uninstall: implicit command based on install
- run
    - stop: implicit command based on run
    - Can be used to run the package. It can be a service, test, etc.

## Dependencies

Explicit/Implicit Dependencies

Explicit dependencies are defined by [extension] dependencies

Implicit are difined by command extension

# Old Design


## Single Source Of Truth

Zeal config is designed to be single source of truth within the lifecycle of the package.

## Package Lifecycle

Zeal can be used to automate full lifecyle of package which includes building,
deployment, and runtime. Zeal is focused on deployment. Build and runtime events
can be used to facilitate automation.

* Build
* Test
* Package
* Install
* Uninstall
* Start
* Stop

## Settings

All commands that affects packages support settings. These are key value pairs
that can be used to replace sections in the config (Like scripts). Settings are
supplied as CLI parameters. Settings flags starts with # symbol followed by its value.

Sample:
zeal pack #environment dev #port 8080

## Extendable Supported Repo

Zeal will support multiple repo types. Each repo will have it's own config that
is outside of zeal. Repo extensions follows contract that zeal expects.

## Locked Versions On Package Creation

Configuration specification supports gettings latest. When a package is created zeal will
get latest from repo and lock it in. This way the artifact dependencies doesn't change
when installed from environment to environment. When dependencies change package should be rebuilt with new dependencies. This way changes are fully traceable.

## Dependencies

Dependencies config accepts multiple dependency types but zeal only recognize packages
dependency type. Different types are up to users needs. One use case is having service
dependencies. This metadata can be used by deployment orchestrator to ensure services
that the package depends on are up and running.

## Settings

### Package

#### name
Package name.
#### locked_dependencies
All dependencies are locked down when installed 