/*
Here's how you can control which files to deploy based on
os and environment.

os is standard Go OS result.
environment is from ZEAL_environment environment variable.
*/
install "contents" "windows" {
    os = "windows"
    package = "windows/service.exe"
    // Can be omitted default is empty
    destination = "" 
}

install "contents" "darwin" {
    os = "darwin"
    package = "./darwin/service"
}

install "contents" "linux" {
    os = "linux"
    package = "./linux/service"
}

// Make executable accessible globally.
install "bin" "service" {
    name = "service_name"
    target_path = "${first("service.*")}"
}

install "dependencies" {
    package1 = {
        version = "1.0.0"
    }
    package2 = {
        version = "2.0.0"
    }
}
