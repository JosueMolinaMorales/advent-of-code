package aoc2016.solutions;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.PriorityQueue;

import aoc2016.utils.Point;

public class Day13 {
    public static void solve() {
        int favNumber = 1350;
        System.out.println("Day 13 part 1: " + part1(favNumber));
        System.out.println("Day 13 part 2: " + part2(favNumber));
    }

    private static ArrayList<Point> getNeighbors(Point point, int number) {
        Point[] directions = {
                new Point(0, 1), // Right
                new Point(-1, 0), // Up
                new Point(1, 0), // Down
                new Point(0, -1), // Left
        };

        ArrayList<Point> neighbors = new ArrayList<>();
        for (Point dir : directions) {
            int dx = point.x + dir.x;
            int dy = point.y + dir.y;

            if (dx < 0 || dy < 0) {
                continue;
            }

            // x*x + 3*x + 2*x*y + y + y*y
            int sum = (dx * dx + 3 * dx + 2 * dx * dy + dy + dy * dy) + number;
            String binary = Integer.toBinaryString(sum);

            // Count 1's
            int bits = 0;
            for (int i = 0; i < binary.length(); i++) {
                if (binary.charAt(i) == '1') {
                    bits += 1;
                }
            }

            if (bits % 2 == 0) {
                // Its open space
                neighbors.add(new Point(dx, dy));
            }
        }

        return neighbors;
    }

    private static class State {
        int steps = 0;
        Point point;

        State(Point p, int steps) {
            this.point = p;
            this.steps = steps;
        }

        @Override
        public String toString() {
            return "State{" +
                    "steps=" + steps +
                    ", point=" + point +
                    '}';
        }
    }

    private static int solve(int number, boolean part2) {
        Point start = new Point(1, 1);
        Point end = new Point(31, 39);
        // Use dijstrkas algorithm
        PriorityQueue<State> queue = new PriorityQueue<>((a, b) -> a.steps - b.steps);
        HashSet<Point> visited = new HashSet<>();
        queue.add(new State(start, 0));
        visited.add(start);

        while (!queue.isEmpty()) {
            State currState = queue.poll();
            if (!part2 && currState.point.equals(end)) {
                return currState.steps;
            }
            if (part2 && currState.steps == 50) {
                break;
            }
            for (Point n : getNeighbors(currState.point, number)) {
                if (visited.contains(n)) {
                    continue;
                }
                visited.add(n);
                queue.add(new State(n, currState.steps + 1));
            }
        }
        return part2 ? visited.size() : -1;
    }

    private static int part1(int number) {
        return solve(number, false);
    }

    private static int part2(int number) {
        return solve(number, true);
    }

}
