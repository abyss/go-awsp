# AWSP (Go Version) - AWS Profile Switcher

Easily switch between AWS Profiles with an interactive selector.

Rewrite [awsp by johnnyopao](https://github.com/johnnyopao/awsp) in golang with minor improvements.

## How it works

The AWS CLI will use the profile present in the `AWS_PROFILE` environment variable, if no flag is set.

Using a combination of a sourced script and a go application, this app parses the current aws configuration (Typically `~/.aws/config`) and provides a filterable list, and then sets that environment variable based on your selection.

## Prerequisites
Set up any number of profiles using the aws cli.

```sh
aws configure --profile PROFILE_NAME
```

You can also leave off the `--profile PROFILE_NAME` param to set your `default` credentials.

Refer to the AWS CLI Documentation for more information:
https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html

## Manual Installation

1) Download and include the `go-awsp` binary in your PATH.

2) Put `_source-awsp` in a known location (`~/bin` is a good spot).

3) Add the following line to your `.bashrc` or `.zshrc` config, including the full path to `_source-awsp`.
```sh
alias awsp="source ~/bin/_source-awsp.sh"
```

> An alias is used because `_source-awsp` needs to be sourced to be able to modify the calling shell's environment variables.

## Usage
Standard usage is just to call `awsp` and select a profile:
```sh
awsp
```
You can type to filter the list, or arrow through the shown options. Press \<Enter\> to select the highlighted option.

You can also type a profile with the command to immediately switch:
```sh
awsp development
```
This is equivalent to directly running `export AWS_PROFILE='development'`.

## Recommendation: Show your AWS Profile in your shell prompt
For better visibility into which AWS Profile is selected it's helpful to configure your prompt to show the value of the env variable `AWS_PROFILE`.

### Examples
Here is a simplified example, just include this in `~/.bashrc` or similar:
```sh
function aws_profile {
  local profile="${AWS_PROFILE:=default}"

  echo "aws:(${profile})"
}

PS1="$PS1 \$(aws_profile)"
```

Here's [@johnnyopao](https://github.com/johnnyopao)'s example (needs oh-my-zsh), with nice colors:

```sh
function aws_prof {
  local profile="${AWS_PROFILE:=default}"

  echo "%{$fg_bold[blue]%}aws:(%{$fg[yellow]%}${profile}%{$fg_bold[blue]%})%{$reset_color%} "
}

PROMPT='$PROMPT $(aws_prof)'
```

A more advanced example for bash can be found in my dotfiles at [https://github.com/abyss/dotfiles](https://github.com/abyss/dotfiles/blob/main/bin/aws-prompt.sh).

## Contributing
Issues and pull requests are welcome. 😄

## License
This project is licensed under the [ISC License](LICENSE.md).

Copyright (c) 2024 Abyss
