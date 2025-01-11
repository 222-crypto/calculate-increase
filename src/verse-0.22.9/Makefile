# The location of the verses
VERSE_LIB = /usr/share/verse/
VERSE_NAME = daily.verse
BIN_DIR = /usr/bin
CFLAGS = -Wall -g -O2 -ansi $(shell dpkg-buildflags --get CFLAGS)
CPPFLAGS = $(shell dpkg-buildflags --get CPPFLAGS)
LDFLAGS = $(shell dpkg-buildflags --get LDFLAGS)
MAN_DIR = /usr/share/man/man1
CC ?= gcc

all : verse

clean :
	rm -f verse *~

distclean : clean

verse : verse.c
	$(CC) $(CFLAGS) $(CPPFLAGS) $(LDFLAGS) -DVERSE_LIB=\"$(VERSE_LIB)$(VERSE_NAME)\" -o verse verse.c

install : verse $(VERSE_NAME)
	install -m 755 -d debian/verse/$(BIN_DIR)
	install -m 755 verse debian/verse/$(BIN_DIR)
	install -m 755 -d debian/verse/$(VERSE_LIB)
	install -m 644 $(VERSE_NAME) debian/verse/$(VERSE_LIB)$(VERSE_NAME)
	install -m 755 -d debian/verse/$(MAN_DIR)
	install -m 644 verse.1 debian/verse/$(MAN_DIR)

