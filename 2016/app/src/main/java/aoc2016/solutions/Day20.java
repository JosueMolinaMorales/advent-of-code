package aoc2016.solutions;

import java.util.ArrayList;

import aoc2016.utils.FileLoader;

public class Day20 {
    public static void solve() {
        String input = FileLoader.loadFile("day20.txt");
        // 22939479 too low
        System.out.println("Day 20 part 1: " + part1(input));
        System.out.println("Day 20 part 2: " + part2(input));
    }

    private static class Range {
        long min;
        long max;

        Range(long min, long max) {
            this.min = min;
            this.max = max;
        }

        boolean contains(long value) {
            return value >= min && value <= max;
        }

        boolean overlaps(Range other) {
            return contains(other.min) || contains(other.max) || other.contains(min) || other.contains(max);
        }

        Range merge(Range other) {
            return new Range(Math.min(min, other.min), Math.max(max, other.max));
        }
    }

    private static ArrayList<Range> parse(String input) {
        ArrayList<Range> ranges = new ArrayList<>();
        for (String line : input.split("\n")) {
            String[] parts = line.split("-");
            ranges.add(new Range(Long.parseLong(parts[0]), Long.parseLong(parts[1]) + 1));
        }
        return ranges;
    }

    private static long solve(String input, boolean countValidIPs) {
        // Find all valid ips
        ArrayList<Range> ranges = parse(input);

        // Sort ranges by min
        ranges.sort((a, b) -> Long.compare(a.min, b.min));

        // Merge overlapping ranges
        ArrayList<Range> merged = new ArrayList<>();
        Range current = ranges.get(0);
        for (int i = 1; i < ranges.size(); i++) {
            Range next = ranges.get(i);
            if (current.overlaps(next)) {
                current = current.merge(next);
            } else {
                merged.add(current);
                current = next;
            }
        }
        merged.add(current);

        // Find the number of valid ips
        if (countValidIPs) {
            long valid = 0;
            for (int i = 0; i < merged.size() - 1; i++) {
                Range currentRange = merged.get(i);
                Range nextRange = merged.get(i + 1);
                valid += nextRange.min - currentRange.max;
            }
            return valid;
        } else {
            return merged.get(0).max;
        }

    }

    private static long part1(String input) {
        return solve(input, false);
    }

    private static long part2(String input) {
        return solve(input, true);
    }
}
