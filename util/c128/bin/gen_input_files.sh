#!/bin/bash

me=$(basename $0)

usage() {
    echo "$me - convert AOC input files to Commodore PET format"
    echo "usage: $me infile"
    exit 1
}

in="$1"

[ -z "$in" ] && usage
#out=$(echo "$1" | sed 's/_input//').pet
out="$in.pet"

petcat -text -w2 -o "$out" -- "$in" && \
    (echo -n '  '; cat "$out") > "$out.prg"
