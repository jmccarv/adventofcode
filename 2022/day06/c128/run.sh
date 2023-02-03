#!/bin/bash

day=day06
d=d6

cl65 -Oris -t c128 -o $day *.c && {
    petcat -text -w2 -o example.pet -- example_input
    petcat -text -w2 -o input.pet -- input
    (echo -n '  '; cat example.pet) > example.pet.prg
    (echo -n '  '; cat input.pet) > input.pet.prg

    c1541 -format mjmtest,jm d71 test.d71 \
        -write $day $day \
        -write example.pet ${d}s.example,s \
        -write input.pet ${d}s.input,s \
        -write example.pet.prg ${d}.example \
        -write input.pet.prg ${d}.input

    x128 test.d71
}
