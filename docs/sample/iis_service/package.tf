package "options" {
    name = "samplePackage"
    description = "Sample Package"
    version = "1.0.0"
    author = "author"
}

package "content" "public" {
    source = "public/*"
    destination = "public"
    exclude = "*.temp"
}

package "content" "binaries" {
    source = "bin/release/*.dll"
    destination = "bin"
}

/* 
syntax: zeal [command] [options]
# if not package is defined use current zeal folder
zeal build
zeal unittest
zeal package

# if package is specified get the command form package
zeal install package_name
zeal uninstall package_name

commands are user defined. The standard is defined by implementers on how their package
is supposed to be used.
*/
