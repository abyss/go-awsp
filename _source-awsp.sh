# shellcheck shell=sh
# This script should be sourced, not called

# Run the command, which will output the chosen profile to ~/.awsp
if [ $# -eq 0 ]; then # No arguments
  go-awsp
  selected_profile=$(cat "$HOME/.awsp")
else
  selected_profile="$*"
  echo "$selected_profile" > "$HOME/.awsp"
fi

# Unset default profile, rather than setting it to "default"
if [ -z "$selected_profile" ] || [ "$selected_profile" = "default" ]; then
  unset AWS_PROFILE
else
  export AWS_PROFILE="$selected_profile"
fi
