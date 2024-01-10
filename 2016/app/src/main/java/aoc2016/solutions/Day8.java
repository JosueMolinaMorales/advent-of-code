package aoc2016.solutions;

import java.io.File;
import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;

public class Day8 {
    private static enum Operation {
        RECT, ROTATE_ROW, ROTATE_COLUMN
    }

    private static class Instruction {
        private Operation operation;
        private int a;
        private int b;

        public Instruction(Operation operation, int a, int b) {
            this.operation = operation;
            this.a = a;
            this.b = b;
        }

        public void applyOperation(boolean[][] grid) {
            switch (operation) {
                case RECT:
                    for (int i = 0; i < b; i++) {
                        for (int j = 0; j < a; j++) {
                            grid[i][j] = true;
                        }
                    }
                    break;
                case ROTATE_ROW:
                    boolean[] row = grid[a];
                    boolean[] newRow = new boolean[row.length];
                    for (int i = 0; i < row.length; i++) {
                        newRow[(i + b) % row.length] = row[i];
                    }
                    grid[a] = newRow;
                    break;

                case ROTATE_COLUMN:
                    boolean[] column = new boolean[grid.length];
                    for (int i = 0; i < grid.length; i++) {
                        column[(i + b) % grid.length] = grid[i][a];
                    }
                    for (int i = 0; i < grid.length; i++) {
                        grid[i][a] = column[i];
                    }
                    break;
            }
        }
    }

    public static void solve() {
        ClassLoader classLoader = Day8.class.getClassLoader();
        File file = new File(classLoader.getResource("day8.txt").getFile());

        System.out.println("Day 8 part 1: " + part1(file));
        System.out.println("Day 8 part 2: ");
        part2(file);
    }

    private static boolean[][] createGrid(int width, int height) {
        boolean[][] grid = new boolean[height][width];
        return grid;
    }

    private static void printGrid(boolean[][] grid) {
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[i].length; j++) {
                if (grid[i][j]) {
                    System.out.print("#");
                } else {
                    System.out.print(".");
                }
            }
            System.out.println();
        }
        System.out.println();
    }

    private static ArrayList<Instruction> parse(File file) {
        ArrayList<Instruction> instructions = new ArrayList<Instruction>();

        try (BufferedReader br = new BufferedReader(new FileReader(file))) {
            String line;
            while ((line = br.readLine()) != null) {
                String[] parts = line.split(" ");
                if (parts[0].equals("rect")) {
                    String[] dimensions = parts[1].split("x");
                    instructions.add(new Instruction(Operation.RECT, Integer.parseInt(dimensions[0]),
                            Integer.parseInt(dimensions[1])));
                } else if (parts[0].equals("rotate")) {
                    if (parts[1].equals("row")) {
                        instructions.add(new Instruction(Operation.ROTATE_ROW, Integer.parseInt(parts[2].split("=")[1]),
                                Integer.parseInt(parts[4])));
                    } else if (parts[1].equals("column")) {
                        instructions.add(new Instruction(Operation.ROTATE_COLUMN,
                                Integer.parseInt(parts[2].split("=")[1]), Integer.parseInt(parts[4])));
                    }
                }
            }
        } catch (IOException err) {
            System.out.println("Failed to read file for day 8");
            System.exit(1);
        }

        return instructions;
    }

    private static int solve(File file, boolean printGrid) {
        boolean[][] grid = createGrid(50, 6);
        ArrayList<Instruction> instructions = parse(file);
        for (Instruction inst : instructions) {
            inst.applyOperation(grid);
        }

        if (printGrid) {
            printGrid(grid);
        }

        // Count number of lit pixels
        int count = 0;
        for (boolean[] row : grid) {
            for (boolean pixel : row) {
                if (pixel) {
                    count++;
                }
            }
        }
        return count;

    }

    private static int part1(File file) {
        return solve(file, false);
    }

    private static void part2(File file) {
        solve(file, true);
    }
}
