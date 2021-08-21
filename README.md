[![Release](https://img.shields.io/github/release/johnmanjiro13/lgotm.svg)](https://github.com/johnmanjiro13/lgotm/releases/latest)

# lgotm

lgotm is a command for generation LGTM image and generated image is copied to clipboard.

lgotm uses [Google Custom Search API](https://developers.google.com/custom-search/v1/introduction). 
Therefore, you need to get api key and engine id of custom search api.

## Installation

### go get

```
go get github.com/johnmanjiro13/lgotm
```

### go install (requires Go1.16+)

```
go install github.com/johnmanjiro13/lgotm@latest
```

## Usage

```
lgotm <query>
```

## Configuration

lgotm can be customized with a configuration file.
The location of the file is `$HOME/.config/lgotm/config` by default.

A default configuration file can be created with the `generate_config_file` sub sommand.

```
lgotm generate_config_file
```

Also, lgotm can be customized with environment variables instead of the file like below.

```
API_KEY=<your-api-key> ENGINE_ID=<your-engine-id> lgotm <query>
```

## Available Options

| Option | Description |
| --- | --- |
| `-h`, `--help` | Show help |
| `--config` | Set config file path |
