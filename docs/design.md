# Refactored Features

## Package Server

Zeal server store metadata and actual files separately so they can be capture separately.

State file can be stored locally or stored on a state server. This is
useful to keep state of short lived servers?

## Use terraform's variable syntax

Can also specify command specific variables

## Configuration can be separated

Configs can be separated on multiple files.

## Reserved options

- dependencies
- variables

## Package Contents/Use

Package items are stored on package. You can pick and choose 
items needed for install.

## Limit commands
build
install
    uninstall is determined by this
package
run
    more generic. Can be used to run test/service

## Dependencies

Explicit/Implicit Dependencies

Explicit dependencies are defined by [command] dependencies [type_provider]

Implicit are difined by command provider

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

## Split Package Per Target OS

Ability to split package per OS. Useful for compiled languages that have different binaries per target OS.

possible repo structure so this can be supported:
package_name/release/patch/@target/package.tar.gz

## Dependencies

Dependencies config accepts multiple dependency types but zeal only recognize packages
dependency type. Different types are up to users needs. One use case is having service
dependencies. This metadata can be used by deployment orchestrator to ensure services
that the package depends on are up and running.

# Design

## Zeal Path

data - zeal metadata - do not manually modify 

logs - regular log files

temp - packages tar.gz files / extracted / script files are stored.

plugins - zeal plugins for supported repo.