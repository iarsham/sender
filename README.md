# Sender for new lines to telegram bot ✈

## Installation

Build the binary using Go's compiler

## Windows

#### 64 bits

```bash
env GOOS=windows GOARCH=amd64 go build .
```

## Linux

#### 64 bits

```bash
env GOOS=linux GOARCH=amd64 go build .
```

## macOS

#### 64 bits

```bash
env GOOS=darwin GOARCH=amd64 go build .
```

## Usage

Please keep in mind that the following file structure is required:

```bash
├── .env
```
The program will panic if it fails to load .env file

### Configuration instructions
in the .env file:
* **TOKEN - >** Add your bot token from botfather
* **TTL - >** Seconds for bot poller
* **FILE - >** filename for scanning
* **USERID - >** your userid from @userinfobot in telegram

### Example

**.env**

```
    TOKEN=43897394701:AAFR8SyAp2k474CjpcDwU0sd8Ssr6_Yk9-c
    TTL=10
    FILE=good.txt
    USERID=1241351234
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
