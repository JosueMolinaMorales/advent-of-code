package aoc2016.solutions;

import aoc2016.utils.FileLoader;

public class Day18 {
    public static void solve() {
        String input = FileLoader.loadFile("day18.txt");
        System.out.println("Day 18 part 1: " + part1(input));
        System.out.println("Day 18 part 2: " + part2(input));
    }

    private static enum Tile {
        TRAP, SAFE;

        public static Tile toTile(char ch) {
            return ch == '.' ? Tile.SAFE : Tile.TRAP;
        }

        public static Tile getTileFromLastRow(int tile, String lastRow) {
            if (tile < 0 || tile >= lastRow.length()) {
                return Tile.SAFE;
            }
            return Tile.toTile(lastRow.charAt(tile));

        }
    }

    private static int solve(String firstRow, int maxRows) {
        int safeTiles = 0;
        for (int i = 0; i < firstRow.trim().length(); i++) {
            safeTiles += firstRow.charAt(i) == '.' ? 1 : 0;
        }

        int i = 1;
        String lastRow = firstRow.trim();
        while (i < maxRows) {
            StringBuilder nextRow = new StringBuilder();
            for (int j = 0; j < lastRow.length(); j++) {
                Tile left = Tile.getTileFromLastRow(j - 1, lastRow);
                Tile center = Tile.getTileFromLastRow(j, lastRow);
                Tile right = Tile.getTileFromLastRow(j + 1, lastRow);
                if ((left == Tile.TRAP && center == Tile.TRAP && right == Tile.SAFE) ||
                        (center == Tile.TRAP && right == Tile.TRAP && left == Tile.SAFE) ||
                        (left == Tile.TRAP && center == Tile.SAFE && right == Tile.SAFE) ||
                        (right == Tile.TRAP && left == Tile.SAFE && center == Tile.SAFE)) {
                    nextRow.append("^");
                } else {
                    nextRow.append(".");
                    safeTiles += 1;
                }

            }

            lastRow = nextRow.toString();
            i++;
        }

        return safeTiles;

    }

    public static int part1(String firstRow) {
        return solve(firstRow.trim(), 40);
    }

    private static int part2(String firstRow) {
        return solve(firstRow.trim(), 400_000);
    }
}
