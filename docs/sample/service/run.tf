/*
Run service using zeal.
*/
run "daemon" "service" {
    executable = "${install.contents.service.file[0]}"
    params = ["--loglevel", "${var.log_level}"]
}

/*
This can be used to parameterized command like:

zeal run [package_name] --log_level "DEBUG"
*/
run "variable" "log_level" {
    default = "INFO"
}
