<div align="center">
  <h1>Clients Service</h1>

[![golang-services](https://github.com/janapc/event-tickets/actions/workflows/golang-services.yml/badge.svg?branch=main)](https://github.com/janapc/event-tickets/actions/workflows/golang-services.yml)

<a href="#description">Description</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#requirement">Requirement</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#usage">Usage</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#resources">Resources</a>

</div>

## Description

this API to manage clients.
consume payment queue for creating customers.
produce messages to the client_created queue and send_ticket after the client is created.

## Requirement

This project your need:

- golang v1.22.1 [golang](https://go.dev/doc/install)

You must create a **.env** file with the same information as in **.env_examples**

## Usage

Run this commands in your terminal:

```sh
## install dependecies
❯ go mod tidy

## run this command to start api(localhost:3004):
❯ go run cmd/main.go

```

API routes are in `http://localhost:3004/clients/docs`

## Resources

- golang
- postgres
- swagger
- fiber
- rabbitmq

<div align="center">

Made by Janapc 🤘 [Get in touch!](https://www.linkedin.com/in/janaina-pedrina/)

</div>
