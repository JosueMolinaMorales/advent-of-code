package aoc2016.solutions;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;

import aoc2016.utils.Point;

public class Day1 {
    private static enum Direction {
        NORTH, EAST, SOUTH, WEST;

        public Direction turn(String direction) {
            switch (this) {
                case NORTH:
                    if (direction.equals("R")) {
                        return EAST;
                    }
                    return WEST;
                case EAST:
                    if (direction.equals("R")) {
                        return SOUTH;
                    }
                    return NORTH;
                case SOUTH:
                    if (direction.equals("R")) {
                        return WEST;
                    }
                    return EAST;
                case WEST:
                    if (direction.equals("R")) {
                        return NORTH;
                    }
                    return SOUTH;
                default:
                    return null;
            }
        }
    }

    private static class Step {
        String direction;
        int distance;

        public Step(String input) {
            direction = input.substring(0, 1);
            distance = Integer.parseInt(input.substring(1));
        }
    }

    public static void solve() {
        ClassLoader classLoader = Day1.class.getClassLoader();
        File file = new File(classLoader.getResource("day1.txt").getFile());

        System.out.println("Day 1 part 1: " + part1(file));
        System.out.println("Day 1 part 2: " + part2(file));
    }

    private static Step[] parseInput(File file) {
        // Read file
        try (BufferedReader reader = new BufferedReader(new FileReader(file))) {
            String line = reader.readLine();
            String[] split = line.split(", ");
            Step[] input = new Step[split.length];
            for (int i = 0; i < split.length; i++) {
                input[i] = new Step(split[i]);
            }
            return input;
        } catch (IOException err) {
            System.out.println("Failed to read file for day 1");
            System.exit(1);
        }
        return null;
    }

    private static int part1(File file) {
        Step[] input = parseInput(file);
        Point point = findBunnyHQ(input, false);
        return point.manhattanDistance(new Point());
    }

    private static int part2(File file) {
        Step[] input = parseInput(file);
        Point point = findBunnyHQ(input, true);
        return point.manhattanDistance(new Point());
    }

    private static Point findBunnyHQ(Step[] instructions, boolean part2) {
        Point point = new Point();
        Direction facing = Direction.NORTH;
        ArrayList<Point> visited = new ArrayList<>();
        for (Step step : instructions) {
            int dx = 0;
            int dy = 0;
            facing = facing.turn(step.direction);
            switch (facing) {
                case NORTH:
                    dy = step.distance;
                    break;
                case EAST:
                    dx = step.distance;
                    break;
                case SOUTH:
                    dy = -step.distance;
                    break;
                case WEST:
                    dx = -step.distance;
                    break;
            }

            // Add all the points we pass through to the visited list
            for (int i = 0; i < Math.abs(dx); i++) {
                if (dx > 0) {
                    point.x++;
                } else {
                    point.x--;
                }
                if (part2 && visited.contains(point)) {
                    return point;
                }
                visited.add(new Point(point.x, point.y));
            }
            for (int i = 0; i < Math.abs(dy); i++) {
                if (dy > 0) {
                    point.y++;
                } else {
                    point.y--;
                }
                if (part2 && visited.contains(point)) {
                    return point;
                }
                visited.add(new Point(point.x, point.y));
            }

        }
        return point;
    }

}
