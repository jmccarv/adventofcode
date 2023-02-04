# Useful Targets:
#   all   - compile the program and generate PETSCII input files and a test .d71 image
#   clean - remove temp files (petscii input files, d71 disk image, compiled executable, etc)
#   build - compile the program
#   test  - compile and run the program in vice
#   write - write program and input files to real floppy drive connected to zoomfloppy

# Device connected to zoomfloppy for use with the 'write' target
DEV = 8

# Name of test disk image to create
D71 = test.d71

CC = cl65
CFLAGS = -t c128 -Osir -I $(INCDIR) -L $(LIBDIR)

DAY = $(shell basename `readlink -f ..`)
D = $(shell echo $(DAY)|sed 's/ay//')

UTILDIR = ../../../util/c128
LIBDIR = $(UTILDIR)/lib
INCDIR = $(UTILDIR)/include

INFILES = example input
INPET = $(addsuffix .pet, $(INFILES))
INPRG = $(addsuffix .prg, $(INPET))

SOURCES = $(wildcard *.c)

.PHONY:	all clean build test inpet write d71

all:	build inpet $(D71)

build:	$(DAY)

test:	$(D71)
	x128 $(D71)

inpet: $(INPET)

.ONESHELL:
write:	build inpet
	cbmctrl --petscii command $(DEV) "s:$(DAY)"
	for fn in $(INFILES); do 
		cbmctrl --petscii command $(DEV) "s:$(D).$$fn"
	done
	cbmwrite -f P $(DEV) $(DAY)
	for fn in $(INFILES); do
		cbmwrite -o $(D).$$fn $(DEV) $$fn.pet.prg
	done
	cbmctrl dir $(DEV)

d71: $(D71)

$(DAY):	$(SOURCES)
	$(CC) -DINFILE_PREFIX='"$(D)"' $(CFLAGS) -o $(DAY) $(SOURCES) util.lib

$(D71):	build inpet
	c1541 -format 'mjm aoc,jm' d71 $(D71) -write $(DAY) $(DAY) \
	$(foreach fn, $(INFILES), -write $(fn).pet.prg $(D).$(fn))

%.pet:	$(INFILES)
	petcat -text -w2 -o $@ -- $(basename $@)
	(echo -n '  '; cat $@) > $@.prg

clean:
	$(RM) $(DAY) $(SOURCES:.c=.o) $(INPET) $(INPRG) $(D71)
