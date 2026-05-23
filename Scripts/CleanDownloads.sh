#!/usr/bin/env bash

DownloadFolder="/data/data/com.termux/files/home/storage/downloads"

PictureFolder="/data/data/com.termux/files/home/storage/pictures"

videosFolder="/data/data/com.termux/files/home/storage/movies"

audioFolder="/data/data/com.termux/files/home/storage/music"

imgExt=(".png" ".jpg" ".jpeg" ".gif")
vidExt=(".mov" ".mp4")
audExt=(".mp3" ".opus")
DownloadFolder
if [ ! -d "$DownloadFolder" ]; then
  exit 1
fi


for file in "$DownloadFolder"/*; do
  if [ -f "$file" ]; then
    echo "$file"
    if [[ $file =~ $imgExt$]]; then
      cp $file $PictureFolder
      rm $file
    fi
    if [[ $file =~ $vidExt$]]; then
      cp $file $videosfolder
      rm $file
    fi
    if [[ $file =~ $audExt$]]; then
      cp $file $audfolder
      rm $file
    fi
  fi
done
