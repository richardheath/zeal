# Features

## Settings

Commands pack, publish, install, uninstall, and start accepts settings flags. 
These flags uses # symbol like #key followed by the value. Sample:

zeal pack #customer someCustomer

customer settings on the config file will be replaced with someCustomer.

settings specified at earlier levels can be overriden at later levels.
? Find a better way to explain this :)

## Extendable Supported Repo

Design repo download/publish structure to support multiple repository types.
Repository handlers are separate cli app that implements the standard that zeal specifies. 

## Locked Versions On Package Creation

Configuration specification supports gettings latest. When a package is created zeal will
get latest from repo and lock it in. This way the artifact dependencies doesn't change
when installed from environment to environment.

## Split Package Per Target OS

Useful for compiled languages that have different binaries per target OS.

# Design

## Zeal Path

data - zeal metadata - do not manually modify 

logs - regular log files

temp - packages tar.gz files / extracted / script files are stored.

plugins - zeal plugins for supported repo.