[![CircleCI](https://circleci.com/gh/vleango/famoustitle/tree/master.svg?style=svg)](https://circleci.com/gh/vleango/famoustitle/tree/master)

# famoustitle

Learn React with Redux, Lambda with Go, and DynamoDB together

## Prerequisites

- Docker
- Docker-compose

## Get the App

```
$ git clone https://github.com/vleango/famoustitle.git
```

## Setup

1. Add an empty `.aws.env` file to `config/environments`

```
$ touch config/environments/.aws.env
```

2. Install node modules for `web-react`

```
$ docker-compose run --rm web-react yarn install
```

3. Install node modules for `lambda-go`

```
$ docker-compose run --rm lambda-go npm i
```

4. Build your lambda functions

```
$ docker-compose run --rm lambda-go make
```

5. Start App

```
$ docker-compose up
```
