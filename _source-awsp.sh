# shellcheck shell=sh
# This script should be sourced, not called
# DEPRECATED: use 'awsp' instead. Update your alias to: alias awsp='. awsp'
echo "awsp: _source-awsp.sh is deprecated and will break in a future release. Update your alias to: alias awsp='. awsp'"

if [ $# -eq 0 ]; then # No arguments
  selected_profile=$(go-awsp)
else
  selected_profile="$*"
fi

# Unset default profile, rather than setting it to "default"
if [ -z "$selected_profile" ] || [ "$selected_profile" = "default" ]; then
  unset AWS_PROFILE
else
  export AWS_PROFILE="$selected_profile"
fi
