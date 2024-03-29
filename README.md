<div align="center">
  <img src=".github/images/snab_logo.svg" width="200px" height="200px" />

  <h1>SnaB - Snake Basket</h1>

  <p>
    <a href="https://github.com/msiebeneicher/snab/actions/workflows/pre-commit.yml">
      <img src="https://github.com/msiebeneicher/snab/actions/workflows/pre-commit.yml/badge.svg" />
    </a>
    <a href="https://github.com/msiebeneicher/snab/actions/workflows/go-ci.yml">
      <img src="https://github.com/msiebeneicher/snab/actions/workflows/go-ci.yml/badge.svg" />
    </a>
  </p>

  <p>
    SnaB is a shell runner, inspired by  <a href="https://taskfile.dev/">task</a> and <a href="https://cobra.dev/">Cobra</a>, and aims to enable you to bundle shell script and commands like a powerful modern CLI application
  </p>

  <p>
    <img src=".github/images/snab_demo.gif" height="350px" />
  </p>

  <hr />
</div>

**Features:**

- Easy installation: just download a single binary, add to $PATH and you're done!
- Simulate own application binary: install and uninstall a starting script of your bundle with simple commands
- Tab completion: Usable tab autocompletion for your bundle
- Generate docs: generate automatically docs for your bundle with a simple command

**Content:**
- [Installation](#installation)
  - [Get The Binary](#get-the-binary)
- [Usage](#usage)
  - [Getting started](#getting-started)
  - [Snabfile](#snabfile)
  - [Command directory](#command-directory)
  - [Sub-Commands and grouping](#sub-commands-and-grouping)
  - [Flags](#flags)
  - [Variables](#variables)
  - [Forwarding CLI arguments to commands](#forwarding-cli-arguments-to-commands)
- [API Reference](#api-reference)
  - [CLI](#cli)
    - [SnaB Commands](#snab-commands)
    - [Global Flags](#global-flags)
    - [ENV](#env)
  - [Schema](#schema)
    - [Snabfile](#snabfile-1)
    - [Description](#description)
    - [Task](#task)
    - [Flag](#flag)
    - [Example](#example)
- [Know issues](#know-issues)
  - [Flag parsing errors when hyphen is used](#flag-parsing-errors-when-hyphen-is-used)


## Installation

### Get The Binary

You can download the binary from the [releases page](https://github.com/msiebeneicher/snab/releases) on GitHub and add to your $PATH.

The checksums.txt file contains the SHA-256 checksum for each file.

## Usage

### Getting started

Create a file called `.snab.yml` in the root of your project. The cmds attribute should contain the commands of a task. Each task reflects and command your final cli app. The example below allows provide a command `foo` which will execute the `foo.sh` script in the same directory.

```yaml
schema_version: '0.2'

name: boa
version: '1.0.0'

description:
  short: 'Boa Example - SnaB - bundle shell script to a powerful modern CLI applications'
  long: 'Boa Example - SnaB - enable you to bundle shell script to a powerful modern CLI applications'
  example: boa /foo/bar/example

tasks:
  foo:
    description:
      short: 'foo script'
      long: 'execute awesome foo script'
      example: 'boa foo [--verbose] PATH'
    cmds:
      - ./foo.sh
```

Running the command is as simple as running:

```sh
snab foo --help
```

### Snabfile

By default SnaB will search for a `.snab.yml` snabfile in the current working directory.
You can set also set a path to your snabfile globally by using the `--snabfile ./path/to/.snab.yml` or by the env var `SNABFILE`.

### Command directory

By default, commands will be executed in the current working directory. But you can easily make the commands run in another folder, informing dir:

```yaml
schema_version: '0.2'

# [...]

tasks:
  foo:
    dir: foobar
    cmds:
      - ./bar.sh
```

### Sub-Commands and grouping

You can add one or more commands to a parent command by setting the parent name in your task config:

```yaml
schema_version: '0.2'

# [...]

tasks:
  foo:
    description:
      example: 'dummy foo [--verbose]'
    cmds:
      - echo "I am foo"

  bar:
    description:
      short: bar category

  foobar:
    parent: bar
    description:
      example: 'dummy bar foobar [--verbose]'
    cmds:
      - echo "I am bar/foobar"
```

### Flags

You are able to define flags for your commands which is important for generating the help and docs.
The defined flags can be used in your commands by [Go's template engine](https://golang.org/pkg/text/template/) and are available under the defined name of the flag.

You can choose between `string` and `boolean` flags.

```yaml
schema_version: '0.2'

# [...]

tasks:
  bar:
    flags:
      - name: optionA
        shorthand: o
        usage: My option to handle something
        value: default-bar-value
        type: string
      - name: optionB
        usage: Second option without default value
      - name: optionC
        usage: Boolean option to handle something
        value: false
        type: bool
    cmds:
      - ./bar.sh --optionA="{{ .optionA }}"
      - ./bar.sh --optionA="{{ .optionA }}"{{ if .optionB }} --optionB="{{ .optionB }}"{{ end }}
      - ./bar.sh --optionA="{{ .optionA }}"{{ if .optionC }} --isOptionC{{ end }}
```

### Variables

You have access to some default variables, which can be used in your commands:

| Variable    | Type   | Description                    |
| ----------- | ------ | ------------------------------ |
| snabfileDir | string | Path to your snabfile location |

```yaml
schema_version: '0.2'

# [...]

tasks:
  foo:
    description:
      example: 'exec foo.sh'
    cmds:
      - "{{ .snabfileDir }}/bin/foo.sh"
```

### Forwarding CLI arguments to commands

By default all arguments will forward to the commands of your task.

The snab command `snab foo my-first-argument my-second-argument` from the example will execute `./foo.sh my-first-argument my-second-argument` under the hood.

## API Reference

### CLI

SnaB command line tool has the following syntax:

```sh
snab [--flags] [commands...] [CLI_ARGS...]
```

If you installed you bundle you can execute your app by the defined name in your snabfile:

```yaml
schema_version: '0.2'

name: boa

# [...]
```

```sh
boa [--flags] [commands...] [CLI_ARGS...]
```

#### SnaB Commands

SnaB provide some default subcommands: `<snab|app> origin --help`

| Command       | Description                                 |
| ------------- | ------------------------------------------- |
| docs generate | Generate markdown docs                      |
| install       | Install snab basket under `/usr/local/bin`  |
| uninstall     | Uninstall snab basket from `/usr/local/bin` |

#### Global Flags

| Short | Flag       | Type   | Default     | Description                       |
| ----- | ---------- | -------| ----------- | --------------------------------- |
|       | --snabfile | string | working dir | Path to your snabfile (hidden)    |
|       | --trace    | bool   | false       | set verbose output to trace level |
| -v    | --verbose  | bool   | false       | set verbose output to debug level |

#### ENV

| ENV      | Default | Description           |
| -------- | ------- | --------------------- |
| SNABFILE |         | Path to your snabfile |

### Schema

#### Snabfile

| Attribute      | Type                        | Default | Description |
| -------------- | ----------------------------| ------- | ----------- |
| schema_version | string                      | 0.2     | Version of the snabfile. The current version is `0.2` |
| name           | string                      |         | Name ouf your snab bundle |
| version        | string                      |         | Version ouf your snab bundle |
| description    | [Description](#description) |         | Description struct |
| tasks          | [map[string]Task](#task)    |         | A set of tasks. The key also reflects the final command name. |

#### Description

| Attribute | Type   | Default | Description |
| --------- | ------ | ------- | ----------- |
| short     | string |         | Short description shown in the 'help' output |
| long      | string |         | Short message shown in the 'help this-command' output |
| example   | string |         | Example of how to use the command |

#### Task

| Attribute   | Type                        | Default             | Description                                    |
| ----------- | --------------------------- | --------------------| ---------------------------------------------- |
| parent      | string                      |                     | Add task commands to this parent command.      |
| description | [Description](#description) |                     | Description struct                             |
| dir         | string                      | _working directory_ | The directory in which this command should run |
| flags       | [[]Flag](#flag)             |                     | Array of Flag structs                          |
| cmds        | []string                    |                     | Array of commands to be execute                |

#### Flag

| Attribute | Type   | Default | Description |
| --------- | ------ | ------- | ----------- |
| name      | string |         | Name of the flag. Used in 'help' output and can be used as go-template var in cmds strings. |
| shorthand | string |         | Shorthand letter that can be used after a single dash |
| usage     | string |         | Usage string used in 'help' output |
| value     | string |         | Default value of the flag |
| type      | string | string  | Type of the Flag value. Possible are the "bool" and "string" |

#### Example

Full `.snab.yml` example:

```yaml
schema_version: '0.2'

name: dummy
version: '0.1.0'

description:
  short: Example app
  long: Example app to provide shell script collection under /usr/local/bin/dummy
  example: dummy --help

tasks:
  # simple command
  foo:
    description:
      short: foo script
      long: echo awesome foo script name
      example: dummy foo [--verbose]
    cmds:
      - echo "I am foo"

  # command with flag
  bar:
    description:
      short: bar script
    flags:
      - name: name
        shorthand: n
        usage: Your name to say hello
        value: Max
        type: string
    cmds:
      -  echo "Hello {{ .Name }}! I am bar."

  # category without own cmds execution
  cat:
    description:
      short: A category to add commands as subcommands
      long:  Only a category for further subcommands

  # subcommand unter `cat`
  foobar:
    parent: cat
    description:
      short: foobar sub command
      long: foobar command under cat
      example: dummy cat foobar [--verbose]
    cmds:
      - echo "I am foobar"
```

## Know issues

### Flag parsing errors when hyphen is used

Failing example:

```yaml
    # [...]

    flags:
      - name: no-fetch
        usage: One example for hyphen usage
        type: bool
        value: false

    cmds:
      - ./my-command {{ if .no-fetch }}--no-fetch{{ end }}
```

Issue is that you can not use the flag name directly.
`if .no-fetch` will not work and end up in a parsing error.

As a workaround you can use the go-template index function: `if index . "no-fetch"`

Working example:

```yaml
    # [...]
    cmds:
      - ./my-command {{ if index . "no-fetch" }}--no-fetch{{ end }}
```
