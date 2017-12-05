/*
Here's how you can control which files to deploy based on
os and environment.

os is standard Go OS result.
environment is from ZEAL_environment environment variable.
*/
install "contents" "service" {
    destination = "service"

    "windows/amd64" = {
        source = "bin/service_windows_amd64.exe"
        destination = "service.exe"
    }
    
    "darwin/amd64" = {
        source = "bin/service_darwin_amd64"
    }

    "linux/amd64" = {
        source = "bin/service_linux_amd64"
    }
}


// Make executable accessible globally.
install "bin" "service" {
    name = "service_name"
    target_path = "${first("service.*")}"
}

// Install dependencies on install
install "dependencies" {
    package1 = {
        version = "1.0.0"
    }
    package2 = {
        version = "2.0.0"
    }
}

install "variable" "installpath" {
    default = "/bin"
}

install "variable" "scope" {
    default = "global"
    description = "control scope where it will be installed"
}
