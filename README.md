[![CircleCI](https://circleci.com/gh/vleango/tech_writer/tree/master.svg?style=svg)](https://circleci.com/gh/vleango/tech_writer/tree/master)

# tech_writer_app

Learn React with Redux, Lambda with Go, and DynamoDB together

## Prerequisites

- Docker
- Docker-compose

## Get the App

```
$ git clone https://github.com/vleango/tech_writer.git
```

## Setup

1. Add an empty `.aws.env` file to `config/environments`

```
$ touch config/environemnts/.aws.env
```

2. Install node modules for `web-react`

```
$ docker-compose run --rm web-react yarn install
```

3. Build your lambda functions

```
$ docker-compose run --rm lambda-go make
```

4. Start App

```
$ docker-compose up
```
