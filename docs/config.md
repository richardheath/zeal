# Zeal Config Design

## name
Name of the package.

## description

## version
Package version.

## keep
Rule that zeal will use to determine if previously installed version needs 
to be uninstalled before installing new version of the package. 
If currently installed doesn’t need to be uninstalled then multiple version will be installed on machine.

> ### major
>Uninstall current version if they have same major version. Otherwise
multi version can exist as long as they have different major version.

> ### minor
> Uninstall current version if they have same major.minor version.
Otherwise multi version can exist as long as they have different major.minor 
version.

> ### patch
> Default setting. Patch is considered the last section of the version number.
This disallows multi version and force only 1 version per instance.

> ### all
> Allow any version.

> ### pN
> Packages not following semver can specify the version position rule.

> ### {{settingsKey}}
> Uninstall current version if they have the same settingsKey value. Otherwise
multi version can exist as long as they have different settingsKey value.
This can be used for custom indicator like {{customer}} which allows multiple
version per target customer for multi tenant deployment strategy.



## author

## destination
Destination path where files will be installed on the target server. Empty path will be relative to current directory.

## files
Files that will be included on the package and its path relative to destination folder.

> The key is the files to add on the package and value is the destination path.

## exclude
Array of regex rules on files to exclude from the package.

## scripts
Scripts are automatically executed based on event lifecycle of a project.

> Array is regular scripts
Object key must be one of the supported platform key and array for the script to use

## Events
* preinstall - runs before package is installed.
* installed - runs after package is installed.
* preuninstall - runs before package is uninstalled.
* uninstalled - runs after package is uninstalled.
* error - runs when install fails.
* prestart - runs before service start
* start - script to run to start service

## dependencies
key/value of package dependencies. Key is the package name and value is the version number needed. (*) Asterisk is used in version number to get latest on version section.
> 1.* - Get latest for major 1

> 1.0.* - Get latest for major 1 minor 0

>  \* - Get latest

## settings
Settings configuration. This allows package to have default value or description for package settings.

## metadata
Key/value metadata info about the package.

## override

Give capability to override config based on target OS or force override using install flag.