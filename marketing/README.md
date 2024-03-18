<div align="center">
  <h1>Marketing API</h1>

<a href="#description">Description</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#requirement">Requirement</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#usage">Usage</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#resources">Resources</a>

</div>

## Description

Api to manage marketing.

## Requirement

To this project your need:

- nodejs v20.6.0 [Nodejs](https://nodejs.org/en/download)

You must create a **.env** file with the same information as in **.env_examples**

## Usage

Run this commands in your terminal:

```sh
## install dependecies
❯ npm i

## run migrate
❯ npm run prisma-migrate-dev

## run this command to start api(localhost:3003):
❯ npm run start

## run this command to start consumer queue:
❯ npm run start:queue

```

API routes are in `http://localhost:3003/leads/docs`

## Resources

- typecript
- express
- mysql
- swagger
- prisma
- rabbitmq

<div align="center">

Made by Janapc 🤘 [Get in touch!](https://www.linkedin.com/in/janaina-pedrina/)

</div>
