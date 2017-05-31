/*
Here's how you can control which files to deploy based on
condition.
*/
install "contents" "dev" {
    condition = "${equal("dev", env.env_name)}"
    package = "./*"
    destination = "${var.target_path}"
}

// On prod don't include PDB files
install "contents" "prod" {
    condition = "${equal("prod", env.env_name)}"
    package = "./*"
    destination = "${var.target_path}"
    exclude = "*.pdb"
}

/* Use other command provider to setup other
requirements of your service.
*/
install "iis" "service_setup" {
    app_pool = "service"
    site_path = "topsite/subsite"
}

/*
This can be used to parameterized command like:

zeal install [package_name] target_path "c:\\other\\path"
*/
install "variable" "target_path" {
    default = "c:\\site\\"
}