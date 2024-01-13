package aoc2016.solutions;

import java.util.ArrayList;
import java.util.PriorityQueue;

import aoc2016.utils.Direction;
import aoc2016.utils.FileLoader;
import aoc2016.utils.MD5;
import aoc2016.utils.Point;

public class Day17 {
    private final static String OPEN_POSITION = "bcdef";

    public static void solve() {
        String gridString = FileLoader.loadFile("day17.txt");
        ArrayList<ArrayList<Character>> grid = parseGrid(gridString);
        String passcode = "rrrbmfta";

        System.out.println("Day 17 part 1: " + part1(passcode, grid));
        System.out.println("Day 17 part 2: " + part2(passcode, grid));
    }

    private static ArrayList<ArrayList<Character>> parseGrid(String gridString) {
        ArrayList<ArrayList<Character>> grid = new ArrayList<>();
        String[] parts = gridString.split("\n");
        for (int i = 0; i < parts.length; i++) {
            ArrayList<Character> row = new ArrayList<>();
            for (int j = 0; j < parts[i].length(); j++) {
                row.add(parts[i].charAt(j));
            }
            grid.add(row);
        }
        return grid;
    }

    private static class State {
        public Point point;
        public String path;
        public int cost;

        public State(Point point, String path, int cost) {
            this.point = point;
            this.path = path;
            this.cost = cost;
        }

    }

    private static ArrayList<Point> getNeighbors(ArrayList<ArrayList<Character>> grid, Point point, String passcode,
            String path) {
        String hash = MD5.getMD5(passcode + path);
        String doors = hash.substring(0, 5);
        ArrayList<Point> neighbors = new ArrayList<>();
        ArrayList<Point> directions = Direction.getNonDiagonalDirections();
        for (int i = 0; i < directions.size(); i++) {
            int dx = point.x + directions.get(i).x;
            int dy = point.y + directions.get(i).y;
            if (dx < 0 || dy < 0 || dx >= grid.size() || dy >= grid.get(0).size()
                    || !OPEN_POSITION.contains(String.valueOf(doors.charAt(i))) || grid.get(dx).get(dy) == '#') {
                continue;
            }

            // If the position is a door, step through it
            if (grid.get(dx).get(dy) == '|' || grid.get(dx).get(dy) == '-') {
                dx += directions.get(i).x;
                dy += directions.get(i).y;
            }

            neighbors.add(new Point(dx, dy));
        }

        return neighbors;
    }

    private static char directionMoved(Point current, Point next) {
        if (current.x == next.x) {
            if (current.y < next.y) {
                return 'R';
            } else {
                return 'L';
            }
        } else {
            if (current.x < next.x) {
                return 'D';
            } else {
                return 'U';
            }
        }
    }

    private static String part1(String passcode, ArrayList<ArrayList<Character>> grid) {
        Point start = new Point(1, 1);
        Point end = new Point(7, 7);

        // Use Dijsktra's algorithm to find the shortest path
        PriorityQueue<State> queue = new PriorityQueue<>((a, b) -> a.cost - b.cost);

        queue.add(new State(start, "", 0));

        while (!queue.isEmpty()) {
            State current = queue.poll();
            if (current.point.equals(end)) {
                return current.path;
            }

            ArrayList<Point> neighbors = getNeighbors(grid, current.point, passcode, current.path);
            for (Point p : neighbors) {
                Point neighborPoint = new Point(p.x, p.y);
                queue.add(new State(neighborPoint, current.path + directionMoved(current.point, neighborPoint),
                        current.cost + 1));
            }
        }

        return "";
    }

    private static int part2(String passcode, ArrayList<ArrayList<Character>> grid) {
        Point start = new Point(1, 1);
        Point[] endRoomPoints = new Point[] { new Point(7, 7), new Point(7, 6), new Point(8, 6) };

        // Find the longest path
        PriorityQueue<State> queue = new PriorityQueue<>((a, b) -> b.cost - a.cost);

        queue.add(new State(start, "", 0));

        int longestPath = 0;

        while (!queue.isEmpty()) {
            State current = queue.poll();
            if (current.point.equals(endRoomPoints[0]) || current.point.equals(endRoomPoints[1])
                    || current.point.equals(endRoomPoints[2])) {
                longestPath = Math.max(longestPath, current.cost);
                continue;
            }

            ArrayList<Point> neighbors = getNeighbors(grid, current.point, passcode, current.path);
            for (Point p : neighbors) {
                Point neighborPoint = new Point(p.x, p.y);
                queue.add(new State(neighborPoint, current.path + directionMoved(current.point, neighborPoint),
                        current.cost + 1));
            }
        }

        return longestPath;
    }
}
