/*
Components implicitly create normalized dependencies section.

build "msbuild" "dependencies" {
    "type": {"name":"version"}
}

*/

build "go" "project" {
    cross_compile = true
}

test "go_test" "unittest" {}