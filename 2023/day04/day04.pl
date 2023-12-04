#!/usr/bin/env perl

use v5.10;

while (<>) {
    $id++; $wins=0;
    ($winning, $numbers) = map { [split /\s+/] } /:\s+(.*)\|\s+(.*)$/;
    %winning = map { $_ => 1 } @$winning;
    $wins += $winning{$_} for (@$numbers);
    $p1 += 2**($wins-1) if $wins;
    $p2 += $copies{$id} + 1;
    for ($i = $id+1; $i < $id+1+$wins; $i++) { $copies{$i} += $copies{$id}+1 };
}
say "Part 1 $p1";
say "Part 2 $p2";
