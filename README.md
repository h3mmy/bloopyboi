# BloopyBoi

BloopyBoi Redux

This is a golang rewrite/continuation of the python [bloopyboi](https://github.com/h3mmy/bloopyboi) with some additional functionality added.
Designed with [bloopysphere](https://github.com/h3mmy/bloopysphere) in mind, but once more developed it should be adaptable for anyone's cluster.
This is still a work in progress, so if you have suggestions for future functionality, open a discussion!

## Description

Bot that can be linked to a discord server. Is able to register slash commands and DMs. Includes a command to generate an inspirobot image.

Runs in a container.

Note: The Helm Chart is Deprecated currently. 

## Authors and acknowledgment

Workflows initially templated from [bjw-s/helm-charts](https://github.com/bjw-s/helm-charts) and [onedr0p/containers](https://github.com/onedr0p/containers) so I could butcher them here.

## License

Apache License 2.0

## Project status

Active-ish

## Developing

Once I have settled on an architectural approach, this will be moved to the [Contributing.md] file

Local Testing (Manual)

`docker build . -t bloopyboi_dev`

`docker run -v ~/projects/bloopyboii/config.yaml:/config.yaml bloopyboi_dev`
