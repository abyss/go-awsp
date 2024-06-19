#!/bin/sh

if [ $# -eq 0 ]; then
  _go-awsp
  selected_profile=$(cat "$HOME/.awsp")
else
  selected_profile="$*"
  echo "$selected_profile" > "$HOME/.awsp"
fi

if [ -z "$selected_profile" ] || [ "$selected_profile" = "default" ]; then
  unset AWS_PROFILE
else
  export AWS_PROFILE="$selected_profile"
fi
