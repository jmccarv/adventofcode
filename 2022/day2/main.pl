#!/usr/bin/env perl

use strict;
use warnings;

my %scores = (
    AX => 3,
    AY => 6,
    AZ => 0,

    BX => 0,
    BY => 3,
    BZ => 6,

    CX => 6,
    CY => 0,
    CZ => 3,
);

my $total = 0;
while (<>) {
    chomp;
    s/\s*//g;

    my $me = 3 - (ord('Z') - ord(substr($_, -1, 1)));
    $total += $me + $scores{$_};
    print "$_  $me + $scores{$_} = ".($me + $scores{$_})." :: $total\n";
}
print $total."\n";
