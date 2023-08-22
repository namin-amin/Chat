# Chat

## Simple Chat app using Golang and Svellte

## RoadMap

- [x] Basic SSE chat protocol
- [ ] User Auth and TLS
- [ ] Groups and Direct Message
- [ ] Encrypted P2P

# ScreenShots

![screenshot1](docs/Demo.gif)

## Running And Compiling

### Requirements

- [Go](https://go.dev/) SDK need to be installed
- [NodeJs](https://nodejs.org/en)

### Run

    make run

### Build Binary

    make compile

### Without make

    cd ui3/
    yarn build
    cd ..
    go build .

## Configuration

- supports configuring **PORT** and **URL**
- can configure using commandline or config.yaml
- commandline

        ./Chat port=3000 url="0.0.0.0"

- config.yaml

```yml
    port: 3000
    url: 0.0.0.0
```

Login to *Url:port/ui/*

Default config is <http://localhost:3000/ui/>

## 💀 Some Known Bugs

- Double message sending on first login
- UI unpolished
- Security needs polishing
- Many Bugs and HouseKeeping
