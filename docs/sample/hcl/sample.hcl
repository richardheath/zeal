properties {
    name = "samplePackage"
    description = "Sample Package"
    version = "1.0.0"
    author = "author"
}

/* 
syntax: command "source tool"
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
build "msbuild" {
    solution = "sample.sln",
    other = "potential_build_options"
}

unittest "vstest" {

}



packager "blob" {
    files = {
        "public/*": "",
        "bin/release/*.dll": "{#message}/samplePackage/{#version}/bin"
    }
}


