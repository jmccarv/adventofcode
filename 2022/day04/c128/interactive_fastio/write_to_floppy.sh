#!/bin/bash

petcat -text -w2 -o example.pet -- example_input
petcat -text -w2 -o input.pet -- input
(echo -n '  '; cat example.pet) > example.pet.prg
(echo -n '  '; cat input.pet) > input.pet.prg


dev=${1:-8}

for fn in day04 d4.example d4.input; do
    cbmctrl --petscii command $dev "s:$fn"
done

cbmwrite -d 1541 -f P $dev day04
cbmwrite -d 1541 -o d4.example $dev example.pet.prg
cbmwrite -d 1541 -o d4.input $dev input.pet.prg

cbmctrl dir $dev
