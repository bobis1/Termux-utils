#!/usr/bin/env bash

batTemp=$(termux-battery-status | jq .temperature)
batPercent=$(termux-battery-status | jq .percentage)
filesystemStats=$(df -h | grep "/data$")
sysUptime=$(uptime)
echo "Storage: $filesystemStats"
echo "Battery Tempurature: $batTemp°C"
echo "Battery Charge: $batPercent%"
echo "System Uptime: $sysUptime"



