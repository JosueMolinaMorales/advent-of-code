package aoc2016.solutions;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.PriorityQueue;

import aoc2016.utils.Direction;
import aoc2016.utils.FileLoader;
import aoc2016.utils.Point;

public class Day24 {
    public static void solve() {
        String input = FileLoader.loadFile("day24.txt");
        System.out.println("Day 24 part 1: " + part1(input));
        System.out.println("Day 24 part 2: " + part2(input));
    }

    private static int part1(String input) {
        return solve(input, false);
    }

    private static int part2(String input) {
        return solve(input, true);
    }

    private static int solve(String input, boolean part2) {
        Grid grid = new Grid(input);

        // permute the POIs
        ArrayList<Point> pois = new ArrayList<>(grid.pois);
        ArrayList<ArrayList<Point>> permutations = new ArrayList<>();
        permute(pois, 0, permutations);

        int minCost = Integer.MAX_VALUE;
        for (ArrayList<Point> permutation : permutations) {
            if (!permutation.get(0).equals(grid.start)) {
                continue;
            }
            int cost = 0;
            for (int i = 0; i < permutation.size() - 1; i++) {
                cost += dijkstra(grid, permutation.get(i), permutation.get(i + 1));
            }
            if (part2) {
                // Add the cost to return to the start
                cost += dijkstra(grid, permutation.get(permutation.size() - 1), grid.start);
            }
            if (cost < minCost) {
                minCost = cost;
            }
        }
        return minCost;
    }

    private static void permute(ArrayList<Point> pois, int start,
            ArrayList<ArrayList<Point>> permutations) {
        if (start >= pois.size()) {
            permutations.add(new ArrayList<>(pois));
        }
        for (int i = start; i < pois.size(); i++) {
            swap(pois, start, i);
            permute(pois, start + 1, permutations);
            swap(pois, start, i);
        }
    }

    private static void swap(ArrayList<Point> pois, int i, int j) {
        Point temp = pois.get(i);
        pois.set(i, pois.get(j));
        pois.set(j, temp);
    }

    private static int dijkstra(Grid grid, Point start, Point end) {
        PriorityQueue<State> queue = new PriorityQueue<>((a, b) -> a.distance - b.distance);
        HashSet<String> visited = new HashSet<>();
        queue.add(new State(start, new HashSet<>(), 0));
        visited.add(queue.peek().toKey());
        while (!queue.isEmpty()) {
            State state = queue.poll();
            if (state.position.equals(end)) {
                return state.distance;
            }
            ArrayList<Point> neighbours = grid.getNeighbours(state.position);
            for (Point neighbour : neighbours) {
                // If the neighbour is a POI and we haven't visited it yet, add it to the list
                State newState = new State(neighbour, new HashSet<>(), state.distance + 1);
                if (visited.contains(newState.toKey())) {
                    continue;
                }
                visited.add(newState.toKey());
                queue.add(newState);
            }
        }
        return 0;
    }

    private static class State {
        Point position;
        HashSet<Point> POIsVisited;
        int distance;

        State(Point position, HashSet<Point> POIsVisited, int distance) {
            this.position = position;
            this.POIsVisited = POIsVisited;
            this.distance = distance;
        }

        public String toKey() {
            StringBuilder sb = new StringBuilder();
            sb.append(position);
            for (Point poi : POIsVisited) {
                sb.append(poi);
            }
            return sb.toString();
        }

        @Override
        public String toString() {
            return "State [position=" + position + ", POIsVisited=" + POIsVisited + ", distance="
                    + distance + "]";
        }
    }

    private static class Grid {
        ArrayList<ArrayList<Character>> grid;
        HashSet<Point> pois;
        Point start;

        Grid(String input) {
            grid = new ArrayList<>();
            pois = new HashSet<>();
            for (String line : input.split("\n")) {
                ArrayList<Character> row = new ArrayList<>();
                for (char c : line.toCharArray()) {
                    row.add(c);
                }
                grid.add(row);
            }
            for (int x = 0; x < grid.size(); x++) {
                for (int y = 0; y < grid.get(x).size(); y++) {
                    if (grid.get(x).get(y) == '0') {
                        start = new Point(x, y);
                    }
                    if (grid.get(x).get(y) != '#' && grid.get(x).get(y) != '.') {
                        pois.add(new Point(x, y));
                    }
                }
            }
        }

        ArrayList<Point> getNeighbours(Point p) {
            ArrayList<Point> neighbours = new ArrayList<>();
            ArrayList<Point> directions = Direction.getNonDiagonalDirections();
            for (Point direction : directions) {
                Point neighbour = new Point(p.x + direction.x, p.y + direction.y);
                if (neighbour.x < 0 || neighbour.x >= grid.size() || neighbour.y < 0
                        || neighbour.y >= grid.get(neighbour.x).size()
                        || grid.get(neighbour.x).get(neighbour.y) == '#') {
                    continue;
                }
                neighbours.add(neighbour);
            }
            return neighbours;
        }
    }
}
