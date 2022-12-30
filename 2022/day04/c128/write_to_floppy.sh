#!/bin/bash

petcat -text -w2 -o example.pet -- example_input
petcat -text -w2 -o input.pet -- input

dev=${1:-8}
cbmwrite -d 1541 -f P $dev day04
cbmwrite -d 1541 -o example -f S $dev example.pet 
cbmwrite -d 1541 -o input -f S $dev input.pet
