<div align="center">
  <h1>Events Service</h1>
  
  [![golang-services](https://github.com/janapc/event-tickets/actions/workflows/golang-services.yml/badge.svg?branch=main)](https://github.com/janapc/event-tickets/actions/workflows/golang-services.yml)

<a href="#description">Description</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#requirement">Requirement</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#run">Run</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#resources">Resources</a>

</div>

## Description

The API to manage events(concerts, shows, theater) using tokens to access(users api) routes.

## Requirement

This project your need:

- golang v1.22.1 [golang](https://go.dev/doc/install)

You must create a **.env** file with the same information as in **.env_examples**.

## Run

Run these commands in your terminal:

```sh
## install dependecies
❯ go mod tidy

## run this command to start api(localhost:3000):
❯ go run cmd/main.go

## this command run the tests:
❯ go test -v ./...

## run this command to update documents if necessary
❯ swag init -g cmd/main.go --output internal/infra/docs

```

API routes are in `http://localhost:3001/events/docs`

## Resources

- golang
- postgres
- swagger
- go-chi

<div align="center">

Made by Janapc 🤘 [Get in touch!](https://www.linkedin.com/in/janaina-pedrina/)

</div>
