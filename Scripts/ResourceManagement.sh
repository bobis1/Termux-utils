#!/usr/bin/env bash

#diaOpts=("Yes" "No" "Skip")
#scenarioOpts=("Tax inspector:I have come to inspect your fine establishment. You are ordered to surrender $taxAmount gold")
#TODO: make a random function appear on start
#gold=100
#food=150
#morale=200
#govSurvail=50
#currentOpt=$diaOpts{2}
#The plan is to have a randomozer that randomly chooses a function, based on the response that is returned the function will react accordingly. Eg. if the taxes senario shows up depending on what the user inputs the senario would unfold. I might switch to go depending on how bad this gets

#run() {
#  local script=$1
#  local workingDir=$(pwd)/Scripts
#  bash $workingDir/$script|| echo "no script found"
  
#}

#expend(){
#  local resource $1
#  local amount $2
#  $resource=$((resource - amount))
#}

#gain(){
 # local resource=$1
#  local amount=$2
  #$resource=$((resource + amount))
#}

#two options: have an array of strings of all of the responses. then run functions based on thet have a check to responses

#source Scripts/ResourcesFunctions.sh
#taxes
#$((resource - amount))
#select choice in "${diaOpts[@]}"; do
#    echo "You chose: $choice"
#    currentOpt=$choice
#     echo $currentOpt
#done


cd Scripts/GoResourceManagement
./Resources