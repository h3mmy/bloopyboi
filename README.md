# BloopyBoi

BloopyBoi Redux

This is a golang rewrite of the python [bloopyboi](https://github.com/h3mmy/bloopyboi) with some additional functionality added, and more integrated into the k8s environment.
Designed with [bloopysphere](https://github.com/h3mmy/bloopysphere) in mind, but it should be adaptable for anyone's cluster.
Still a work in progress. Templates still alive below.

## Description

Bot that can be linked to a discord server. Is able to register slash commands and DMs. Includes a command to generate an inspirobot image.

Runs in a container, and comes with a helm chart for k8s integration (WIP)

## Authors and acknowledgment

Workflows initially templated from [bjw-s/helm-charts](https://github.com/bjw-s/helm-charts) and [onedr0p/containers](https://github.com/onedr0p/containers)

## License

Apache License 2.0

## Project status

Active-ish

## Developing

Local Testing

`docker build . -t bloopyboi_dev`

`docker run -v ~/projects/bloopyboii/config.yaml:/config.yaml bloopyboi_dev`
