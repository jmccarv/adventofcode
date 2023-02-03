#!/bin/bash

day=day08
d=d8

petcat -text -w2 -o example.pet -- example_input
petcat -text -w2 -o input.pet -- input
(echo -n '  '; cat example.pet) > example.pet.prg
(echo -n '  '; cat input.pet) > input.pet.prg

dev=${1:-8}

for fn in $day ${d}.example ${d}.input; do
    cbmctrl --petscii command $dev "s:$fn"
done

cbmwrite -d 1541 -f P $dev $day
cbmwrite -d 1541 -o ${d}.example $dev example.pet.prg
cbmwrite -d 1541 -o ${d}.input $dev input.pet.prg

cbmctrl dir $dev
