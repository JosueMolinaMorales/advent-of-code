package aoc2016.solutions;

import java.util.ArrayList;
import aoc2016.utils.FileLoader;

public class Day15 {
    public static void solve() {
        String input = FileLoader.loadFile("day15.txt");
        System.out.println("Day 15 part 1: " + part1(input));
        System.out.println("Day 15 part 2: " + part2(input));
    }

    private static class Disc {
        int positions;
        int initial;
        int position;

        Disc(int positions, int initial) {
            this.positions = positions;
            this.initial = initial;
            this.position = initial;
        }

        Disc(String line) {
            String[] parts = line.split(" ");
            positions = Integer.parseInt(parts[3]);
            initial = Integer.parseInt(parts[11].replace(".", ""));
            position = initial;
        }

        void tick(int time) {
            position = (initial + time) % positions;
        }

        boolean isOpen() {
            return position == 0;
        }
    }

    private static ArrayList<Disc> parse(String input) {
        ArrayList<Disc> discs = new ArrayList<>();
        for (String line : input.split("\n")) {
            discs.add(new Disc(line));
        }
        return discs;
    }

    private static int solve(ArrayList<Disc> discs) {
        int time = 0;
        while (true) {
            if (check(discs, time)) {
                return time;
            }
            time += 1;
        }
    }

    private static int part1(String input) {
        return solve(parse(input));
    }

    private static int part2(String input) {
        ArrayList<Disc> discs = parse(input);
        discs.add(new Disc(11, 0));
        return solve(discs);
    }

    private static boolean check(ArrayList<Disc> discs, int time) {
        for (int i = 0; i < discs.size(); i++) {
            Disc disc = discs.get(i);
            disc.tick(time + i + 1);
            if (!disc.isOpen()) {
                return false;
            }
        }
        return true;
    }
}
