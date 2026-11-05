# TibiaData API assets

[![assets-generator](https://github.com/tibiadata/tibiadata-api-assets/actions/workflows/generate.yml/badge.svg)](https://github.com/tibiadata/tibiadata-api-assets/actions/workflows/generate.yml)
[![code-coverage](https://codecov.io/gh/tibiadata/tibiadata-api-assets/graph/badge.svg?token=ZEWYWLUZ9Y)](https://codecov.io/gh/tibiadata/tibiadata-api-assets)
[![GitHub go.mod version](https://img.shields.io/github/go-mod/go-version/tibiadata/tibiadata-api-assets?logo=go)](https://github.com/tibiadata/tibiadata-api-assets/blob/main/go.mod)
[![GitHub license](https://img.shields.io/github/license/tibiadata/tibiadata-api-assets)](https://github.com/tibiadata/tibiadata-api-assets/blob/main/LICENSE)

JSON assets that are almost static for the TibiaData API.

This repo contains tooling that generates the assets json files used by the [tibiadata-api-go](https://github.com/tibiadata/tibiadata-api-go) image.

The generated file can be accessed on [assets.tibiadata.com](https://assets.tibiadata.com) in a minified or pretty version.

## What's inside

The generation will be generating new json files on a scheduled interval, providing updates JSON data files used by the container.

There are some details missing from the upstream servers called on by [tibiadata-api-go](https://github.com/tibiadata/tibiadata-api-go) application, so therefore this data provides information.

Example of data that is missing is a list of valid worlds that can have houses/guildhalls or mapping for every houseid for us to provide the town and housetype.

Assets that are available in the files are worlds, towns, houses, creatures and spells.

## General information

Tibia is a registered trademark of [CipSoft GmbH](https://www.cipsoft.com/en/). Tibia and all products related to Tibia are copyright by [CipSoft GmbH](https://www.cipsoft.com/en/).

## Credits

- Authors: [Tobias Lindberg](https://github.com/tobiasehlert) – [List of contributors](https://github.com/tibiadata/tibiadata-api-assets/graphs/contributors)
- Distributed under [MIT License](LICENSE)
