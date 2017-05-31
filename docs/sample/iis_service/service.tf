/*
Sample on how zeal can be extended based on your own architecture.
This explains that this service depends on two other services
at runtime. Data can be used to create third party dependency
graph. For advanced implementation this can also be used to
automate spin up of dependencies by orchestration tool.
*/
service "dependency" "service1" {
    version = "1.0.0"
}

service "dependency" "service2" {
    version = ">=2.0.0"
}