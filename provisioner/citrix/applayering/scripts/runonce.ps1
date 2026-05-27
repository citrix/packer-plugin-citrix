$runOncePaths = @(
    "HKLM:\SOFTWARE\Microsoft\Windows\CurrentVersion\RunOnce",
    "HKLM:\SOFTWARE\Wow6432Node\Microsoft\Windows\CurrentVersion\RunOnce"
)

foreach ($path in $runOncePaths) {
    if (Test-Path $path) {
        $entries = Get-ItemProperty -Path $path

        foreach ($prop in $entries.PSObject.Properties) {
            if ($prop.Name -notin 'PSPath','PSParentPath','PSChildName','PSDrive','PSProvider') {
                $command = $prop.Value
                Write-Host "Running: $command"
                try {
                    Start-Process -FilePath "cmd.exe" -ArgumentList "/c `"$command`"" -Wait -NoNewWindow
                    Remove-ItemProperty -Path $path -Name $prop.Name -ErrorAction SilentlyContinue
                    Write-Host "Executed and removed: $prop.Name"
                } catch {
                    Write-Warning "Failed to run or remove: $prop.Name - $_"
                }
            }
        }
    } else {
        Write-Host "Registry path does not exist: $path"
    }
}

shutdown.exe -f -r -t 0 -c "packer restart"