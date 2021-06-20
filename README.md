<p align="center">
  <h1 align="center">KeyPrompter 📽️</h1>
</p>

<p align="center">
  <a href="https://i.imgur.com/hw1oNl6.gif" title="KeyPrompter"><img alt="KeyPrompter" src="https://i.imgur.com/hw1oNl6.gif" width="500"></a>
</p>

<p align="center">
  <a href="https://github.com/HiDeoo/KeyPrompter/actions/workflows/keyprompter.yml"><img alt="Integration Status" src="https://github.com/HiDeoo/KeyPrompter/actions/workflows/keyprompter.yml/badge.svg"></a>
  <a href="https://github.com/HiDeoo/KeyPrompter/blob/master/LICENSE"><img alt="License" src="https://badgen.now.sh/badge/license/MIT/blue"></a>
  <br /><br />
</p>

**Broadcast your shortcuts…**

## Motivations

I could not find a macOS application to broadcast keyboard shortcuts to a web page so it can be easily used with various live streaming softwares so I decided to build my own.

KeyPrompter is a Go application monitoring keyboard shortcuts and also serving locally a React application which can be used to visualize the pressed shortcuts (_only shortcuts are visible meaning typing text won't display anything_).

## Installation

1. [Download the latest version](https://github.com/HiDeoo/KeyPrompter/releases) matching your system architecture.
1. Extract the downloaded archive.
1. Optionally, move the `keyprompter` application in your `PATH` (e.g. `/usr/local/bin/`)
1. Ensure the `keyprompter` application can be executed (e.g. `chmod +x keyprompter`).

_Note: the application has been developed and tested only on macOS with a QWERTY layout and even tho it should technically work on other systems, it will probably require some adjustments and testing._

## Usage

You can quickly starts the application using the following command:

```plaintext
$ keyprompter
You can now view the KeyPrompter UI in the browser: http://localhost:8484.
```

Various options are available:

| Option     | Default | Description                                    |
| ---------- | ------- | ---------------------------------------------- |
| `-p VALUE` | 8484    | Port used to run the web UI.                   |
| `-c VALUE` | _None_  | Path to the optional client configuration file |
| `-v`       |         | Display the application version.               |
| `-h`       |         | Display the application options.               |

## Configuration

Using the `-c` option, an optional path to a client configuration file can be passed down to the application.

The configuration is a [TOML file](https://toml.io/) where any value can be configured individually. Here is an example with the default values for each options:

```toml
# The maximum number of shortcuts to display at the same time.
Count = 5

# The duration in seconds during when a shortcut is visible on screen.
Duration = 5

# The font size in pixels used when displaying a shortcut.
FontSize = 18

# The font color used for a shortcut text (any supported CSS value).
FontColor = "white"

# The background color for a shortcut (any supported CSS value).
BgColor = "rgba(0, 0, 0, 0.6)"
```

## Contribute

1. [Fork](https://help.github.com/articles/fork-a-repo) & [clone](https://help.github.com/articles/cloning-a-repository) this repository.
1. Make sure your Go version is at least `1.16`.
1. Build & run the development version using `go run .`.

## License

Licensed under the MIT License, Copyright © HiDeoo.

See [LICENSE](https://github.com/HiDeoo/KeyPrompter/blob/master/LICENSE) for more information.
