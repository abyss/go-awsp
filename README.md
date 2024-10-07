# awsp - Go-based AWS Profile Switcher

Easily switch between AWS Profiles using an interactive selector.

This is a rewrite of [awsp by johnnyopao](https://github.com/johnnyopao/awsp) in Go, with minor improvements.

## How It Works

The AWS CLI uses the profile defined in the `AWS_PROFILE` environment variable if no specific profile flag is set.

Using a combination of a sourced script and a Go application, this tool parses the current AWS configuration (typically `~/.aws/config`). It then provides a filterable list and sets the `AWS_PROFILE` environment variable based on your selection.

## Requirements

Set up any number of profiles using the AWS CLI.

```sh
aws configure --profile PROFILE_NAME
```

You can also leave off the `--profile PROFILE_NAME` parameter to set your `default` credentials.

Refer to the [AWS CLI Documentation](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html) for more information.

## Installation
The AWS Profile Switcher can be installed on macOS, Linux, or Windows, and works with most standard shells, including PowerShell.

In non-PowerShell shells, an alias is needed because `_source-awsp.sh` must be sourced to modify the parent shell's environment variables.

### Homebrew (Mac/Linux)
1) `brew install abyss/tools/awsp`

2) Add the following to your shell profile (e.g., `.bashrc` or `.zshrc`):
```sh
alias awsp='source "$(brew --prefix awsp)/_source-awsp.sh"'
```

### Manual (Bash, Zsh, etc.)
1) Download the `go-awsp` binary from the [releases page](https://github.com/abyss/go-awsp/releases) and include it in your PATH.
2) Place the script `_source-awsp.sh` in a known directory, such as `~/bin`.
3) Add the following to your shell profile (e.g., `.bashrc` or `.zshrc`), including the full path to `_source-awsp.sh`:
```sh
alias awsp="source ~/bin/_source-awsp.sh"
```

### Manual (PowerShell)
1) Download the `go-awsp` binary from the [releases page](https://github.com/abyss/go-awsp/releases) and include it in your PATH.
2) Put `awsp.ps1` in your PATH.
> Dot sourcing is not necessary in PowerShell when using this script.

If AWS Tools for PowerShell is installed, it will also attempt to set the default profile using `Set-AWSCredential`.

## Usage
To use the interactive profile switcher, simply run `awsp` and select a profile.
```sh
awsp
```
You can type to filter the list or use the arrow keys to navigate through the options. Press \<Enter\> to select the highlighted profile.

You can also specify a profile with the command to switch immediately:
```sh
awsp development
```
This is equivalent to directly running `export AWS_PROFILE='development'`.

## Recommendation: Show Your AWS Profile in Your Prompt
For better visibility into which AWS Profile is selected, it's helpful to configure your prompt to show the value of the environment variable `AWS_PROFILE`.

### Examples
Here is a simplified example: Add this to your shell profile (e.g., `.bashrc` or `.zshrc`):
```sh
function aws_profile {
  local profile="${AWS_PROFILE:=default}"

  echo "aws:(${profile})"
}

PS1="$PS1 \$(aws_profile)"
```

Here is [@johnnyopao's](https://github.com/johnnyopao) example, which requires Oh My Zsh and includes color customization:

```sh
function aws_prof {
  local profile="${AWS_PROFILE:=default}"

  echo "%{$fg_bold[blue]%}aws:(%{$fg[yellow]%}${profile}%{$fg_bold[blue]%})%{$reset_color%} "
}

PROMPT='$PROMPT $(aws_prof)'
```

A more advanced example for Bash can be found in my [dotfiles on GitHub](https://github.com/abyss/dotfiles).

## Contributing
Contributions in the form of issues and pull requests are welcome. 😄

## License
This project is licensed under the [ISC License](LICENSE.md).

Copyright (c) 2024 Abyss
