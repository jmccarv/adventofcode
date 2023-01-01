#!/usr/bin/env perl

use strict;
use warnings;

sub pr_struct {
    my $nr = shift;
    print <<EOT;
struct pairs {
    int nr;
    unsigned char min1[$nr];
    unsigned char max1[$nr];
    unsigned char min2[$nr];
    unsigned char max2[$nr];
};

const struct pairs assn = {
EOT
}

my $nr = 0;
my @assn;
while (<>) {
    chomp;
    s/-/,/g;
    my @vals = split ',';
    next unless @vals == 4;

    for (my $i = 0; $i < 4; $i++) {
        push @{$assn[$i]}, $vals[$i];
    }
    $nr++;
}

pr_struct($nr);
print "$nr,\n";

for (@assn) {
    print "{".join(',',@{$_})."},\n";
}

print "};\n";
