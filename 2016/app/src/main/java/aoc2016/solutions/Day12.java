package aoc2016.solutions;

import java.util.HashMap;

import aoc2016.utils.FileLoader;

public class Day12 {
    public static void solve() {
        String input = FileLoader.loadFile("day12.txt");
        System.out.println("Day 12 part 1: " + part1(input));
        System.out.println("Day 12 part 2: " + part2(input));
    }

    private static int solve(String input, HashMap<Character, Integer> registers) {
        String[] instructions = input.split("\n");
        int i = 0;
        while (i < instructions.length) {
            String[] tokens = instructions[i].split(" ");
            String action = tokens[0];
            switch (action) {
                case "cpy":
                    if (tokens[1].matches("[a-d]")) {
                        registers.put(tokens[2].charAt(0), registers.get(tokens[1].charAt(0)));
                    } else {
                        registers.put(tokens[2].charAt(0), Integer.parseInt(tokens[1]));
                    }
                    break;
                case "inc":
                    registers.put(tokens[1].charAt(0), registers.get(tokens[1].charAt(0)) + 1);
                    break;
                case "dec":
                    registers.put(tokens[1].charAt(0), registers.get(tokens[1].charAt(0)) - 1);
                    break;
                case "jnz":
                    if (tokens[1].matches("[a-d]")) {
                        if (registers.get(tokens[1].charAt(0)) != 0) {
                            i += Integer.parseInt(tokens[2]);
                            continue;
                        }
                    } else {
                        if (Integer.parseInt(tokens[1]) != 0) {
                            i += Integer.parseInt(tokens[2]);
                            continue;
                        }
                    }
                    break;
                default:
                    break;
            }
            i++;
        }
        return registers.get('a');
    }

    public static int part1(String input) {
        HashMap<Character, Integer> registers = new HashMap<>();
        registers.put('a', 0);
        registers.put('b', 0);
        registers.put('c', 0);
        registers.put('d', 0);
        return solve(input, registers);
    }

    public static int part2(String input) {
        HashMap<Character, Integer> registers = new HashMap<>();
        registers.put('a', 0);
        registers.put('b', 0);
        registers.put('c', 1);
        registers.put('d', 0);
        return solve(input, registers);
    }
}
