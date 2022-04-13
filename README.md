# X32-MQTT

This application provides a link between the Behringer X32 Sound-mixer and MQTT
topics. It uses the OSC protocol to listen for changes on different OSC
addresses and then publishes the correct values to the MQTT broker.

Future us includes sending MQTT messages to control some mixers functions as
well.

## Configuration

Configuration is done using a config file. Either located in `.`
or `/etc/x32-mqtt`. The file should be named `x32-mqtt.y[a]ml`.

If no config file is found the default are loaded and an attempt to write a new
config file to current directory is made.

## Running

The program must be compiled in a golang environment. No external dependencies
are required to the final binary.

### Docker
A docker container can be built using the included Dockerfile like so:
```shell
docker build -t insektionen/x32-mqtt .
```