/*
You can use output of other command as input to 
properties.
*/
test "vstest" "unittest" {
    solution = "../solution_name"
    test_dlls = "${build.msbuild.test_dlls}"
    failOnError = true
}