# OTA Provider

## Installation

This is the OTA Provider for github.com/flomon/espota.

A docker image is available in this repository under

- docker.pgk.github.com/flomon/ota-provider/ota-provider:{tag} (x86 version)
- docker.pgk.github.com/flomon/ota-provider/ota-provider-rpi:{tag} (arm7 version)

> to save the data add a volume for /app/data

The service can also be deployed by compiling the code yourself just run

```sh
go build src/main.go
```

and then run the created `main` binary from the root repository.

## Setup

On first Startup you will be asked to supply your github credentials (use a Personal Access Token) and the address of your MQTT Broker.

Then you can add repositories to track and the device to which the binary should be deployed to.

OTA Provider will on your behalf fetch the Github API every minute to check for new releases and download the first attachment and save it available under http://hostname/bin/{clientName}.bin

A binary is only downloaded when a release changes or a new one is added and only then the deployment to your device is triggered. You can also manually trigger a deployment of the saved binary to your device from the user interface.
