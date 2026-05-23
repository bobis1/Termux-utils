#!/usr/bin/env bash
date
echo "---------Readying System---------"
scriptsOpts=("CleanDownloads" "PhoneStats" "Snake")

run() {
  local script=$1
  local workingDir=$(pwd)/Scripts
  bash $workingDir/$script.sh|| echo "no script found"
  
}


select choice in "${scriptsOpts[@]}"; do
    echo "You chose: $choice"
    run $choice
    break
done