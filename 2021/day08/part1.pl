#!/usr/bin/env perl

use strict;
use warnings;

use Data::Dumper qw(Dumper);

my %interested_lengths = (
    2 => 1,
    4 => 1,
    3 => 1,
    7 => 1,
);

my $ret;
while (<>) {
    chomp;
    my ($patterns, $digits) = split /\|/;
    $ret += grep { $interested_lengths{length($_)} } grep { length($_) } split /\s/, $digits;
}
print "$ret\n";
