if ($args.Count -eq 0) { # No arguments
    $selected_profile = & go-awsp
} else {
    $selected_profile = $args -join ' '
}

# Set-AWSCredential is attempted to also set the profile with AWS Tools for PowerShell
# If the command doesn't exist, or the profile isn't valid, this will squelch the error

if ([string]::IsNullOrEmpty($selected_profile) -or $selected_profile -eq "default") {
    Remove-Item Env:AWS_PROFILE -ErrorAction SilentlyContinue
    try { Set-AWSCredential -ProfileName 'default' -Scope Global } catch {}
} else {
    $env:AWS_PROFILE = $selected_profile
    try { Set-AWSCredential -ProfileName $selected_profile -Scope Global } catch {}
}
