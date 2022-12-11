#!/usr/bin/env perl 

# https://adventofcode.com/2020/day/2

use strict;
use warnings;

my ($valid1, $valid2);

while (<>) {
    next unless /^([0-9]+)-([0-9]+)\s+(.):\s+(\S*)$/;
    $valid1 += valid1($1, $2, $3, $4);
    $valid2 += valid2($1, $2, $3, $4);
}
print "part 1 -- $valid1 valid passwords\n";
print "part 2 -- $valid2 valid passwords\n";

sub valid1 {
    my ($min, $max, $letter, $password) = @_;
    my $nr = grep { $_ eq $letter } split '', $password;
    $nr >= $min && $nr <= $max;
}

sub valid2 {
    my ($p1, $p2, $l, $pwd) = @_;
    (substr($pwd, $p1-1, 1) eq $l) xor (substr($pwd, $p2-1, 1) eq $l);
}
