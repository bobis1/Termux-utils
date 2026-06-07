expend(){
  local resource $1
  local amount $2
  $resource=$((resource - amount))
}

gain(){
  local resource=$1
  local amount=$2
  $resource=$((resource + amount))
}


taxes() {
  local taxAmount=23
  echo "The tax inspector has come to inspect your fine establishment. You are ordered to surrender $taxAmount gold"
  if [[ $choice == "yes" ]]; then
    expend $gold taxAmount
  fi
  if [[ $choice == "no" ]]; then
    gain $govSurvail 300
  fi
}


