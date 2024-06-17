#!/bin/sh
export PATH=$PATH:/usr/bin:/nix/store
ffmpeg -version || exit 1
/out