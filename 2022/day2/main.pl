#!/usr/bin/env perl

use strict;
use warnings;
use Data::Dumper;

my @input;
while (<>) {
    chomp;
    s/\s*//g;
    push @input, $_;
}

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

part1(@input);
part2(@input);

sub part1 {
    #print Dumper(\@_);
    my $total = 0;
    for (@_) {
        my $me = 3 - (ord('Z') - ord(substr($_, -1, 1)));
        $total += $me + $scores{$_};
    }
    print $total."\n";
}

sub part2 {
    my %choice = (
        AX => 'AZ',
        AY => 'AX',
        AZ => 'AY',

        BX => 'BX',
        BY => 'BY',
        BZ => 'BZ',

        CX => 'CY',
        CY => 'CZ',
        CZ => 'CX',
    );

    part1(map { $choice{$_} } @_);
}
