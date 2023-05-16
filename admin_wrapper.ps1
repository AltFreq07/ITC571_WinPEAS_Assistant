# Get the directory of the current script
$scriptDir = Split-Path -Path $MyInvocation.MyCommand.Definition -Parent

# Check if the current process has administrative privileges
$isElevated = ([Security.Principal.WindowsPrincipal] [Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole] "Administrator")

if ($isElevated) {
    # If the process has administrative privileges, execute the main script
    & (Join-Path $scriptDir "temp_script.ps1")
} else {
    # Relaunch the current script with administrative privileges
    Start-Process powershell.exe -ArgumentList "-ExecutionPolicy Bypass -File `"$($MyInvocation.MyCommand.Path)`"" -Verb RunAs -Wait
}
