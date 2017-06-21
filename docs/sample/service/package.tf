package "settings" {
    name = "samplePackage"
    description = "Sample Package"
    version = "1.0.0"
    author = "author"
}

package "content" "executables" {
    source = "compiled/*"
    destination = ""
}
