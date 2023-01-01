#!/bin/bash

go run gen.go > day05_data.h
cl65 -t c128 -o day05 day05.c
