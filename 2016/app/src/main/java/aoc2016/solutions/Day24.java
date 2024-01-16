package aoc2016.solutions;

import java.util.ArrayList;

import aoc2016.utils.FileLoader;
import aoc2016.utils.Point;

public class Day24 {
    public static void solve() {
        String input = FileLoader.loadFile("day24.txt");
        System.out.println("Day 24 part 1: " + part1(input));
        System.out.println("Day 24 part 2: " + part2(input));
    }

    private static int part1(String input) {
        Grid grid = new Grid(input);
        // Print POIs
        for (Point poi : grid.pois) {
            System.out.println(poi);
        }
        return 0;
    }

    private static int part2(String input) {
        return 0;
    }

    private static class Grid {
        ArrayList<ArrayList<Character>> grid;
        ArrayList<Point> pois;

        Grid(String input) {
            grid = new ArrayList<>();
            pois = new ArrayList<>();
            for (String line : input.split("\n")) {
                ArrayList<Character> row = new ArrayList<>();
                for (char c : line.toCharArray()) {
                    row.add(c);
                }
                grid.add(row);
            }
            for (int y = 0; y < grid.size(); y++) {
                for (int x = 0; x < grid.get(y).size(); x++) {
                    if (grid.get(y).get(x) != '#' && grid.get(y).get(x) != '.') {
                        pois.add(new Point(x, y));
                    }
                }
            }
        }
    }
}
