#!/usr/bin/env perl

use strict;
use warnings;

use Data::Dumper qw(Dumper);
use Array::Set qw(set_diff);

# length => digit
my %known = (2=>1, 4=>4, 3=>7, 7=>8);

my $ret;
while (<>) {
    chomp;
    my ($patterns, $digits) = split /\|/;
    $ret += process([ split /\s/, $patterns ], [ grep { length($_) } split /\s/, $digits ]);
}
print "$ret\n";

sub process {
    my ($patterns, $digits) = @_;
    my %k;

    # Turn them into objects! yay!
    my @p = map { pattern($_) } @$patterns;

    # Let's start by finding our known digits, these are the ones that
    # use a unique number of segments in the 7 segment display: 1, 4, 7, and 8
    # This will be a map of { digit => pattern_object } i.e. { 7 => $obj }
    %k = map { $_->value => $_ } grep { $_->value } @p;

    # Now filter out those that remain to be decoded
    my @left = grep { !defined($_->value) } @p;

    # We can figure out which pattern is 6: it has 6 letters in its pattern
    # and and one of 'c' or 'f'. So we can subtract our known 7 pattern
    # from the possible 6's to find the one that must be 6
    $k{6} = (grep { $_->minus($k{7}) == 4 }
             grep { $_->len == 6 } @left)[0];
    $k{6}->{value} = 6;
    @left = grep { !$_->equal($k{6}) } @left;

    # From 6 we can find 5, 5 has all the same segments as 6, minus one
    $k{5} = (grep { $k{6}->minus($_) == 1 } 
             grep { $_->len == 5 } @left)[0];
    $k{5}->{value} = 5;
    @left = grep { !$_->equal($k{5}) } @left;
   
    # 3 - 5 yields 1
    $k{3} = (grep { $_->minus($k{5}) == 1 }
             grep { $_->len == 5 } @left)[0];
    $k{3}->{value} = 3;
    @left = grep { !$_->equal($k{3}) } @left;

    # 2 - 5 = 2
    $k{2} = (grep { $_->minus($k{5}) == 2 }
             grep { $_->len == 5 } @left)[0];
    $k{2}->{value} = 2;
    @left = grep { !$_->equal($k{2}) } @left;

    # Only 0 and 9 left now.
    # 5 - 9 = 0, 5 - 0 = 1
    $k{9} = (grep { $k{5}->minus($_) == 0 }
             grep { $_->len == 6 } @left)[0];
    $k{9}->{value} = 9;
    @left = grep { !$_->equal($k{9}) } @left;

    $k{0} = shift @left;
    $k{0}->{value} = 0;

    # Now we can decode the digits
    my $num;
    for my $d (@$digits) {
        my $x = (grep { $_->matches($d) } values %k)[0];
        $num .= $x->value;
    }
    $num;
}

sub pattern {
    my @letters = sort split //, shift;
    my $pattern = join('', @letters);
    bless {
        pattern => $pattern,
        letters => \@letters,
        len => length($pattern),
        value => $known{length($pattern)},
    };
}

sub letters  { @{shift->{letters}} }
sub value    { shift->{value} }
sub len      { shift->{len} }
sub minus    { my @r = sort @{set_diff(shift->{letters}, shift->{letters})}; @r }
sub contains { shift->{pattern} =~ /shift/ }
sub equal    { shift->{pattern} eq shift->{pattern} }
sub matches  { shift->{pattern} eq join('', sort split //, shift) }
