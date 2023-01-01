#!/bin/bash

./gen_input.pl > d4b_data.h

cl65 -t c128 -o d4b d4b.c && x128 d4b
