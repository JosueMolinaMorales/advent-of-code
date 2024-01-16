package aoc2016.solutions;

import java.util.HashSet;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

import aoc2016.utils.FileLoader;
import aoc2016.utils.Point;

public class Day22 {
    public static void solve() {
        String input = FileLoader.loadFile("day22.txt");
        System.out.println("Day 22 part 1: " + part1(input));
        System.out.println("Day 22 part 2: " + part2(input));
    }

    private static int part1(String input) {
        HashSet<Node> nodes = parse(input);
        // How many viable pairs of nodes are there?
        int viablePairs = 0;
        for (Node a : nodes) {
            for (Node b : nodes) {
                if (a == b) {
                    continue;
                }
                if (a.used > 0 && a.used <= b.avail) {
                    viablePairs += 1;
                }
            }
        }

        return viablePairs;
    }

    private static int part2(String input) {
        HashSet<Node> nodes = parse(input);
        createGrid(nodes);
        // 5 moves to move the empty node to the top right
        // 45 moves to empty the top right
        // 30 moves to move the empty node to the top left
        return 45 + (30 * 5);
    }

    private static void createGrid(HashSet<Node> nodes) {
        int maxX = 0;
        int maxY = 0;
        for (Node node : nodes) {
            if (node.x > maxX) {
                maxX = node.x;
            }
            if (node.y > maxY) {
                maxY = node.y;
            }
        }
        char[][] grid = new char[maxY + 1][maxX + 1];
        for (Node node : nodes) {
            if (node.used == 0) {
                grid[node.y][node.x] = '_';
            } else if (node.used > 100) {
                grid[node.y][node.x] = '#';
            } else {
                grid[node.y][node.x] = '.';
            }
        }
        grid[0][0] = 'G';
        grid[0][maxX] = 'E';
        // Print grid
        for (int y = 0; y <= maxY; y++) {
            for (int x = 0; x <= maxX; x++) {
                System.out.print(grid[y][x]);
            }
            System.out.println();
        }
    }

    private static class Node extends Point {
        int size;
        int used;
        int avail;

        Node(String line) {
            super();
            String regex = "/dev/grid/node-x(\\d+)-y(\\d+)\\s+(\\d+)T\\s+(\\d+)T\\s+(\\d+)T\\s+(\\d+)%";
            // Create a Pattern object
            Pattern pattern = Pattern.compile(regex);

            // Create a Matcher object
            Matcher matcher = pattern.matcher(line);

            // Check if the input matches the pattern
            if (matcher.matches()) {
                // Extract values using group() method
                x = Integer.parseInt(matcher.group(1));
                y = Integer.parseInt(matcher.group(2));
                size = Integer.parseInt(matcher.group(3));
                used = Integer.parseInt(matcher.group(4));
                avail = Integer.parseInt(matcher.group(5));
            } else {
                System.out.println("Input does not match the pattern");
            }
        }

        @Override
        public String toString() {
            return String.format("Node(%d, %d, %d, %d, %d)", x, y, size, used, avail);
        }
    }

    private static HashSet<Node> parse(String input) {
        HashSet<Node> nodes = new HashSet<>();
        for (String line : input.split("\n")) {
            if (!line.contains("/dev/grid/node")) {
                continue;
            }
            nodes.add(new Node(line));
        }
        return nodes;
    }
}
