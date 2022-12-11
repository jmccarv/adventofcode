#!/usr/bin/env perl

use strict;
use warnings;

# Read initial stack configuration
my @stacks;
while (<>) {
    last if /^\s*[0-9]/;
    for (my $sidx = 0; length($_) > 1; $_ = substr($_, 4), ++$sidx) {
        unshift @{$stacks[$sidx]}, substr($_,1,1) if substr($_,1,1) ne ' ';
    }
}

# Make a copy for part2
my @p2stacks = map { [ @$_ ] } @stacks;

# Read and execute instructions to move crates
while (<>) {
    next unless /move (\d+) from (\d+) to (\d+)/;
    my ($nr, $from, $to) = ($1, $2-1, $3-1);

    push @{$stacks[$to]}, reverse splice (@{$stacks[$from]}, -$nr);
    push @{$p2stacks[$to]}, splice(@{$p2stacks[$from]}, -$nr);
}

print pop @$_ for @stacks;   print "\n";
print pop @$_ for @p2stacks; print "\n";
