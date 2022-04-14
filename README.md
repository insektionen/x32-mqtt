# X32-MQTT

This application provides a link between the Behringer X32 Sound-mixer and MQTT
topics. It uses the OSC protocol to listen for changes on different OSC
addresses and then publishes the correct values to the MQTT broker.

## Configuration

Configuration is done using a config file. Either located in `.`
or `/etc/x32-mqtt`. The file should be named `x32-mqtt.y[a]ml`.

If no config file is found the default are loaded and an attempt to write a new
config file to current directory is made.

## Running

The program must be compiled in a golang environment. No external dependencies
are required to the final binary.

## Topics and payload

The topics used by the software follows the same syntax as the X32 mixers
OSC addresses. Each address is reflected to its own topic with retention to
allow user interfaces to load the current state from the topic.

Each mixer address is also mirrored to a "set" topic. The config file can
specify a prefix for all MQTT topics, used for both reads and writes (set).

Example:

| Topic                          | Payload                                        | Description                   |
|--------------------------------|------------------------------------------------|-------------------------------|
| `<prefix>/ch/02/config/name`   | `{"value": "Channel 02", "type": "string" }`   | Report a new channel name     |
| `<prefix>/set/ch/02/config/name` | `{"value": "Mychannel", "type": "string" }`    | Set a new name for channel 02 |

Since OSC can send several values on the same address depending on the
configuration object each payload consists of a JSON encoded array of objects.
Each object contains information about the value and the binary format used for
the value.

Example:

```json
[
  {
    "value": 0.75,
    "type": "float32"
  }
]
```

*Note*: The type is always needed when using a set topic (or the software does
not know how to send the data).

OSC addresses for the X32 mixer have been documented unofficially
in [this document][1].

# Docker

A docker container can be built using the included Dockerfile like so:

```shell
docker build -t insektionen/x32-mqtt .
```

[1]: https://wiki.munichmakerlab.de/images/1/17/UNOFFICIAL_X32_OSC_REMOTE_PROTOCOL_%281%29.pdf
