# Define the template folder path and the destination parent folder
$templatePath = "./template"  # Replace with the full path to your template folder
$destinationParentPath = "./"  # Replace with the parent directory for the new folder

# Get the current day of the month
$currentDay = "day" + (Get-Date).ToString("dd")
$currentDayNumber = (Get-Date).ToString("dd")

# Define the destination folder path
$destinationPath = Join-Path -Path $destinationParentPath -ChildPath $currentDay

# Check if the destination folder already exists
if (Test-Path -Path $destinationPath) {
    Write-Host "Folder for today ($currentDay) already exists: $destinationPath" -ForegroundColor Yellow
    exit
}

# Copy the template folder to the destination
try {
    Copy-Item -Path $templatePath -Destination $destinationPath -Recurse -Force
    Write-Host "Template folder has been copied and renamed to: $destinationPath" -ForegroundColor Green
} catch {
    Write-Host "An error occurred during the copy operation: $_" -ForegroundColor Red
    exit
}

# Replace `XX` in specific files within the destination folder
try {
    # Get all files in the destination folder
    Get-ChildItem -Path $destinationPath -File | ForEach-Object {
        # Check if the file name contains 'dayXX'
        if ($_.Name -match "dayXX") {
            # Create the new file name with the current day number
            $newFileName = $_.Name -replace "XX", $currentDayNumber
            $newFilePath = Join-Path -Path $destinationPath -ChildPath $newFileName

            # Rename the file
            Rename-Item -Path $_.FullName -NewName $newFileName
            Write-Host "Renamed file: $($_.Name) to $newFileName" -ForegroundColor Cyan
        }
    }
} catch {
    Write-Host "An error occurred while renaming files: $_" -ForegroundColor Red
}