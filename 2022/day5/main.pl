#!/usr/bin/env perl

use strict;
use warnings;

# Read initial stack configuration
my @stacks;
while (<>) {
    my $sidx = 0;
    last if /^\s*[0-9]/;
    for (; length($_) > 1; $_ = substr($_, 4), ++$sidx) {
        unshift @{$stacks[$sidx]}, substr($_,1,1) if substr($_,1,1) ne ' ';
    }
}

# Make a copy for part2
my @p2stacks = map { [ @$_ ] } @stacks;

# Read and execute instructions to move crates
while (<>) {
    next unless /move (\d+) from (\d+) to (\d+)/;
    my ($nr, $from, $to) = ($1, $2, $3);

    push @{$stacks[$to-1]}, reverse splice (@{$stacks[$from-1]}, -$nr);
    push @{$p2stacks[$to-1]}, splice(@{$p2stacks[$from-1]}, -$nr);
}

print pop @$_ for @stacks;   print "\n";
print pop @$_ for @p2stacks; print "\n";
