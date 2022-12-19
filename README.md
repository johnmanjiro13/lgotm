[![GitHub Workflow Status](https://github.com/johnmanjiro13/lgotm/actions/workflows/test.yaml/badge.svg?branch=main)](https://github.com/johnmanjiro13/lgotm/actions/workflows/test.yaml)
[![Release](https://img.shields.io/github/release/johnmanjiro13/lgotm.svg)](https://github.com/johnmanjiro13/lgotm/releases/latest)

# lgotm

lgotm is a command for generation LGTM image and generated image is copied to clipboard.

## Installation

### Homebrew

```shell
# tap and install
brew tap johnmanjiro13/tap
brew install lgotm

# install directory
brew install johnmanjiro13/tap/lgotm
```

### go get

```shell
go get github.com/johnmanjiro13/lgotm
```

### go install (requires Go1.16+)

```shell
go install github.com/johnmanjiro13/lgotm@latest
```

## Usage

### With Google Custom Search

```shell
lgotm query <query>
```

lgotm query searches images by [Google Custom Search API](https://developers.google.com/custom-search/v1/introduction). 
Therefore, you need to get api key and engine id of custom search api.

## Configuration

lgotm can be customized with a configuration file.
The location of the file is `$HOME/.config/lgotm/config` by default.

A default configuration file can be created with the `generate_config_file` sub sommand.

```shell
lgotm generate_config_file
```

Also, lgotm can be customized with environment variables instead of the file like below.

```shell
API_KEY=<your-api-key> ENGINE_ID=<your-engine-id> lgotm query <query>
```
