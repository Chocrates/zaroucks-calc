# Zaroucks Calculator
This small go program will output the amounts of nutrients needed to build a Spirulina growth medium for a specified volume of water

## Usage
At the moment you need to install [golang](https://go.dev/doc/install) as a pre-requisite

```bash
go run . -o <volume, in liters>
```

```bash
Usage:
  zaroucks-calc [OPTIONS]

Application Options:
  -o, --volume= The volume of liquid in liters, that will be used to created the media.  Note: The volume will increase with the nutrients so choose a value
                smaller than the size of your bioreactor

Help Options:
  -h, --help    Show this help message

```
