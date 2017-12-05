package "settings" {
    name = "samplePackage"
    description = "Sample Package"
    version = "1.0.0"
    author = "author"
    updateRule = "minor"
}

package "content" "executables" {
    source = "${build.go.project.files}"
    destination = "/bin"
}
