Zeal
====

With the explosion of different tools and languages, Zeal aims to
be the common denominator to simplify development, packaging,
installation and running of packages. Zeal is the abstraction 
point of package lifecyle using idempotent scripts.

This can be used to simplify DevOps work. It doesn't matter
what technology is used the projects can have similar pattern on
how it is built, tested, installed, etc. This can also simplify
CICD since the build job just needs to run the same command.

# Multi OS packaging

Only create/publish platform packages if all file source requirements
are satisfied on all commands.

For given sample below. Package for windows will only
be created when `bin/service_windows_amd64.exe` exist on
package. Same this with mac it will only create package
when `bin/service_darwin_amd64` exist on package.

```
install "contents" "service" {
    destination = "service"

    "windows/amd64" = {
        source = "bin/service_windows_amd64.exe"
        destination = "service.exe"
    }
    
    "darwin/amd64" = {
        source = "bin/service_darwin_amd64"
    }
}
```

add `--expectedPackages` flag as a check of what package
platforms are expected from current package command.

# Common Files

To save storage source files that are shared by multi platform are
stored on common zip package. These can still be overriden by files
on platform specific packages.

# Notes

Think about support for multi version
Finalize design support on multi OS
Do not support different files per OS on packaging

# Configuration

## Basic Config Structure

Each section of zeal config is composed of command, extension,
name, and options. Name is any text that means something to user 
which can be used to access output when using interpolation.

```
[command] [extension] [name] {
    [options]
}
```

## Terraform/HCL Syntax

Zeal will use HCL syntax for config files. Interpolation syntax
will also be supported.

## Configuration can be separated

Configs can be separated on multiple files.

# Commands

Package lifecyle is expressed as commands in Zeal. A package
lifecyle is composed of build, package, install, and run.

## build

This abstract building of package like compiling, 
linting etc. This step ensures that package files are in good 
condition. This step is usually executed before packaging but
some packages don't need builds like static files.

## test

## package

This abstracts creation of deployable package.

## install

Install configuration defines extentions used on how to install
a package. Install extensions can implicitly specify uninstall
instructions.

```
zeal install package-1.0.0
```

## uninstall

Usually install extensions will automatically specify uninstall
extensions to use. Extra uninstall can be specified for extra
clean up instructions like clearing logs.

## run

Abstraction on how package can be executed.

```
zeal package_name run
zeal package_name run #message hello
```

## stop

Stop running package.

# Extentions

Extensions are executables that will handle the options 
of zeal configuration.

## Extensible Commands

Allow custom commands.

## Extension Registry

Zeal cli will download extensions from registry. Registry
can point to local registry to control what extensions can
be used.

## Reserved extension names

- dependencies
- variables
- settings

## Extensions Uninstall

Install extensions are automatically called when uninstalled.
Ideally it should undo the setup that it does. Contents are
automatically removed.

## Metadata

### Output

Extensions can give output data that other extensions can use.

### Dependencies

Extensions can also specify package dependencies.

Explicit/Implicit Dependencies

Explicit dependencies are defined by [extension] dependencies

Implicit are difined by command extension

# Old Design

## Extensions Can Create Zeal Config

An extension can create another extension config. Sample
usecase is for install extension telling zeal on how to
uninstall by giving uninstall instructions using zeal config.

implicit configs.

## Single Source Of Truth

Zeal config is designed to be single source of truth within the lifecycle of the package.

## Settings

All commands that affects packages support settings. These are key value pairs
that can be used to replace sections in the config (Like scripts). Settings are
supplied as CLI parameters. Settings flags starts with # symbol followed by its value.

Sample:
zeal pack #environment dev #port 8080

## Locked Versions On Package Creation

Configuration specification supports gettings latest. When a package is created zeal will
get latest from repo and lock it in. This way the artifact dependencies doesn't change
when installed from environment to environment. When dependencies change package should be rebuilt with new dependencies. This way changes are fully traceable.

#### locked_dependencies
All dependencies are locked down when installed 