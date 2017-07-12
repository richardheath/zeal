Zeal Server
===========

Server API that zeal cli will use.

## Metadata

Provide API to grab metadata of package. Zeal will always get
metadata separately than package files.

## Minimum Files Needed

Provide API to grab package contents based on OS and environment filter. Server will send compressed file with contents that
satisfies the filter. This is useful for packages that have 
binaries for multiple OS and only one is needed for target OS
where package is installed.