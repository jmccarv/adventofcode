LIB = ../lib/util.lib
CC = cl65
AR = ar65
CFLAGS = -t c128 -Osir -I ../include
SOURCES = $(wildcard *.c)

.PHONY:	all clean

all:	$(LIB)

%.o:	%.c
	$(CC) -c $(CFLAGS) -o $@ $<

$(LIB):	$(SOURCES:.c=.o)
	$(AR) r $(LIB) $(SOURCES:.c=.o)

clean:
	$(RM) $(SOURCES:.c=.o)
