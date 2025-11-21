install_go_dependencies() {
    echo "Installing go dependencies..."
    go get -tool github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
    go get -tool  go.uber.org/mock/mockgen@latest
    # go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.2
    go get -tool github.com/securego/gosec/v2/cmd/gosec@latest
    go get -tool github.com/google/wire/cmd/wire@latest
}

install_npm_dependencies() {
    npm install -g openapi-format@1.16.0
    npm install -g @redocly/cli@1.9.1
}

install_python_dependencies() {
    pip install argparse
    pip install sqlfluff
}

install_npm_dependencies
install_go_dependencies
install_python_dependencies