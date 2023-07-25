package stacks

const Golang = `
name: CI

on:
  push:
    branches:
      - main
  pull_request:
    branches:
      - main

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
    - uses: actions/checkout@v2
    - name: Setup Go
      uses: actions/setup-go@v2
      with:
        go-version: '1.16'
    - name: Build
      run: go build
    - name: Test
      run: go test ./...
`
