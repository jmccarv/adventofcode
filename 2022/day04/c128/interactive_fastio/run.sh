#!/bin/bash

cl65 -Oris -t c128 -o day04 day04.c && {
    petcat -text -w2 -o example.pet -- example_input
    petcat -text -w2 -o input.pet -- input
    (echo -n 'xx'; cat example.pet) > example.pet.prg
    (echo -n 'xx'; cat input.pet) > input.pet.prg

    c1541 -format mjmtest,jm d71 test.d71 \
        -write day04 day04 \
        -write example.pet d4s.example,s \
        -write input.pet d4s.input,s \
        -write example.pet.prg d4.example \
        -write input.pet.prg d4.input

    x128 test.d71
}
