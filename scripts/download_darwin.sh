#!/bin/bash
set -e

path="$1"
echo "Path used: $path"
name="$2"
echo "Name used: $name"
tracker="$3"
echo "Tracker used: $tracker"

osascript -e '
set x to "a"
  tell app "iTerm"
    activate
    tell the first terminal
      launch session "Default Session"
      tell the last session
        set name to "TvShowDownloader"
        write text "cd '$path'"
        write text "ctorrent -u '$tracker' '$name'"
        write text "exit"
      end tell
    end tell
  end tell'
