package aoc2016.solutions;

public class Day19 {
    public static void solve() {
        int numElves = 3_014_603;
        System.out.println("Day 19 part 1: " + part1(numElves));
        System.out.println("Day 19 part 2: " + part2(numElves));
    }

    private static int part1(int numElves) {
        int[] elvesPresents = new int[numElves];
        for (int i = 0; i < numElves; i++) {
            elvesPresents[i] = 1;
        }
        int i = 0;
        while (true) {
            if (elvesPresents[i] == 0) {
                i = (i + 1) % numElves;
                continue;
            }

            int nextElf = (i + 1) % numElves;
            while (elvesPresents[nextElf] == 0) {
                nextElf = (nextElf + 1) % numElves;
            }
            elvesPresents[i] += elvesPresents[nextElf];
            elvesPresents[nextElf] = 0;
            if (elvesPresents[i] == numElves) {
                return i + 1;
            }

            i = nextElf % numElves;
        }
    }

    private static int part2(int numElves) {
        // https://www.reddit.com/r/adventofcode/comments/5j4lp1/comment/dbdf50n/?utm_source=share&utm_medium=web2x&context=3
        int i = 1;
        while (i * 3 < numElves) {
            i *= 3;
        }

        return numElves - i;
    }

}
