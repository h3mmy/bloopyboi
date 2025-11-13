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

Once I have settled on an architectural approach, this will be moved to a contributing file

To generate ent schema

`go generate ent`

Local Testing (Manual)

`docker build . -t bloopyboi_dev`

`docker run -v ~/projects/bloopyboii/config.yaml:/config.yaml bloopyboi_dev`

## AI notes

This is part of a very limited set of repositories where I will try some AI agent experiments. Anyone who knows me, knows that I have a large amount of distrust for the current landscape. I've been working with ML for many years and understand capabilities and limits. Any code that makes use of AI agents still needs to abide by the code of conduct and have the human author sign-off. Since LLM use has proliferated to the detriment of everything, I will personally scrutinize every PR extra. I intend to summarize my observations in a different document in the future.

So far, the only worthwhile tool I've found within free usage is Google's Jules. Tasks mostly have to be limited to fixing linting errors and very simple bugs. It's not good with much else. Copilot with Anthropic Claude used up the monthly limit of tokens on simply fixing six linting errors (which it failed to complete as well)
