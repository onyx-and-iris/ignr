![Windows](https://img.shields.io/badge/Windows-0078D6?style=for-the-badge&logo=windows&logoColor=white)
![Linux](https://img.shields.io/badge/Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black)
![macOS](https://img.shields.io/badge/mac%20os-000000?style=for-the-badge&logo=macos&logoColor=F0F0F0)

# ignr-cli

Simple no-frills .gitignore generator backed by the Github API.

![Selection Prompt](./img/selectionprompt.png)

## Install

```console
go install github.com/onyx-and-iris/ignr-cli@latest
```

## Authentication

You can run this tool without authenticating but requests will have a stricter rate limiting. 

If you prefer to authenticate you can pass a token in the following ways:

*Flag*

-   --token/-t: Github API Token

*Environment Variable*

```bash
#!/usr/bin/env bash

export GH_TOKEN=<API Token>
```

## Commands

### New

Trigger the selection prompt.

```console
ignr-cli new
```

The prompt filter can activated by pressing `/`:

![Prompt Filter](./img/promptfilter.png)

## Special Thanks

-   [Charm](https://github.com/charmbracelet) for their awesome CLI packages.
