# CLI Design
 
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

## repo

List supported repository types. 

usage: repo list

Commands available per repository type. Options provided will be passed to repository handler.

usage: repo repoType [get,set,delete,list] repo.name --options

--username

--password

--type

--artifact nameOfArtifactType

## config
Set zeal configuration.

usage: config [get,set,delete,list] configKey configValue

## query
Query zeal metadata.

## start
Start a package.

## stop
Stop a package.

## build

Build package.

## test

Test package.