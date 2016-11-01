# Design

## pack
Config file to package. If config file is not supplied, then zeal will use zeal.json file as default.

usage: pack [configPath]

## publish
Package and publish a given config file. If a tar.gz file is specified then it will skip packaging.

usage: publish [configPath|tar.gz] [options]

options:

--repo - Repo name where artifact will be published. Will use default when option is not specified.
 
## install
Install a package and packages that it depends on.

usage: install \<name\>@\<version\> [options]

options: 

--repo - Force search of package to only use specified repo. Otherwise will use order of repo in config.
 
## uninstall
Uninstall a package from the machine.

usage: uninstall \<name\>@[\<version\>|settings] [options]

options:

--force - Force the files to be removed even if scripts from package is failing.

--all - Uninstall all installed versions on the machine.

## Settings Feature

Commands pack, publish, install, uninstall, and start accepts settings flags. 
These flags uses # symbol like #key followed by the value. Sample:

zeal pack #customer someCustomer

customer settings on the config file will be replaced with someCustomer.

settings specified at earlier levels can be overriden at later levels.
? Find a better way to explain this :)

## Repo

## Extensions
Extend command line to support other artifact creation like state?

## Repo
repo [get,set,delete,list] repo.name --options

--username

--password

--type

--artifact nameOfArtifactType

## Config
Add [target] value

Add or update if it doesn’t exist.
 
Set [target] value

Don’t need just use ADD
 
Remove [target] value
  