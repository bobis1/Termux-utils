#!/usr/bin/env bash
date
echo "---------Readying System---------"
scriptsOpts=("CleanDownloads.sh" "PhoneStats.sh" "Snake.sh" "ResourceManagement.sh")

run() {
  local script=$1
  local workingDir=$(pwd)/Scripts
  bash $workingDir/$script|| echo "no script found"
  
}


select choice in "${scriptsOpts[@]}"; do
    echo "You chose: $choice"
    run $choice
    break
done