#!/usr/bin/env perl

use strict;
use warnings;
use Data::Dumper;

exit main();

sub main {
    my $preamble_length = shift @ARGV;
    die 'Usage: ',$ARGV[0], 'preamble_length'
        unless $preamble_length > 1;
    
    my @nums = map { chomp; $_ } <>;
    
    my $nr = part1($preamble_length, @nums);
    print "part1: ",$nr,"\n";
    print "part2: ",part2($nr, @nums)."\n";
}

sub part1 {
    my ($preamble, @nums) = @_;
    my @sums;

    # precompute sums for the first $preamble elements
    for my $i (0..$preamble-2) {
        push @sums, { n => $nums[$i], s => [map { $nums[$i]+$_ } @nums[($i+1)..$preamble-1]] };
    }

    # $prev holds next number we'll need to sum...
    my $prev = $nums[$preamble-1];

    for my $n (@nums[$preamble..$#nums]) {
        print pretty(@sums)."\n";
        return $n unless grep{ $_ eq $n } map { @{$_->{s}} } @sums;

        shift @sums;
        push @sums, { n => $prev, s => [] };
        push @{$_->{s}}, $n + $_->{n} for @sums;
        $prev = $n;
    }

    return 0;
}

sub part2 {
    my ($nr, @nums) = @_;

    while (my ($i, $n1) = each @nums) {
        my ($sum, $min, $max) = ($n1, $n1, $n1);

        for my $n2 (@nums[$i+1..$#nums]) {
            last if ($sum += $n2) > $nr;
            $min = $n2 if $n2 < $min;
            $max = $n2 if $n2 > $max;
            return $min+$max if $sum == $nr;
        }
    }
    0;
}

sub pretty {
    print "$_->{n} [".join(' ', @{$_->{s}})."]\n" for @_;
}
