package aoc2016.solutions;

import java.util.HashMap;

import aoc2016.utils.FileLoader;

public class Day23 {
    public static void solve() {
        String input = FileLoader.loadFile("day23.txt");
        System.out.println("Day 23 part 1: " + part1(input));
        System.out.println("Day 23 part 2: " + part2(input));
    }

    private static boolean validate(String[] tokens) {
        String action = tokens[0];
        switch (action) {
            case "cpy":
                return tokens.length == 3 && tokens[2].matches("[a-d]");
            case "inc":
                return tokens.length == 2 && tokens[1].matches("[a-d]");
            case "dec":
                return tokens.length == 2 && tokens[1].matches("[a-d]");
            case "jnz":
                return tokens.length == 3;
            case "tgl":
                return tokens[1].matches("[a-d]");
            default:
                return false;
        }
    }

    private static int solve(String input, HashMap<Character, Integer> registers) {
        String[] instructions = input.split("\n");
        int i = 0;
        while (i < instructions.length) {
            String[] tokens = instructions[i].split(" ");
            String action = tokens[0];
            if (!validate(tokens)) {
                i++;
                continue;
            }
            // Optimized version of the assembly code
            if (i == 5) {
                registers.put('a', registers.get('a') + (Math.abs(registers.get('b')) *
                        Math.abs(registers.get('d'))));
                registers.put('c', 0);
                registers.put('d', 0);
                i += 5;
                continue;
            }
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
                            if (tokens[2].matches("[a-d]")) {
                                i += registers.get(tokens[2].charAt(0));
                            } else {
                                i += Integer.parseInt(tokens[2]);
                            }
                            continue;
                        }
                    }
                    break;
                case "tgl":
                    int offset = registers.get(tokens[1].charAt(0));
                    if (i + offset < instructions.length) {
                        String[] toggle = instructions[i + offset].split(" ");
                        if (toggle.length == 2) {
                            if (toggle[0].equals("inc")) {
                                toggle[0] = "dec";
                            } else {
                                toggle[0] = "inc";
                            }
                        } else {
                            if (toggle[0].equals("jnz")) {
                                toggle[0] = "cpy";
                            } else {
                                toggle[0] = "jnz";
                            }
                        }
                        instructions[i + offset] = String.join(" ", toggle);
                    }
                    break;
                default:
                    break;
            }
            i++;
        }
        return registers.get('a');
    }

    private static int part1(String input) {
        HashMap<Character, Integer> registers = new HashMap<>();
        registers.put('a', 7);
        registers.put('b', 0);
        registers.put('c', 0);
        registers.put('d', 0);
        return solve(input, registers);
    }

    private static int part2(String input) {
        HashMap<Character, Integer> registers = new HashMap<>();
        registers.put('a', 12);
        registers.put('b', 0);
        registers.put('c', 0);
        registers.put('d', 0);
        return solve(input, registers);
    }
}
