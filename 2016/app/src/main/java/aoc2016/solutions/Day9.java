package aoc2016.solutions;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;

public class Day9 {
    public static void solve() {
        ClassLoader classLoader = Day9.class.getClassLoader();
        File file = new File(classLoader.getResource("day9.txt").getFile());
        String input = readFile(file);
        System.out.println("Day 9 part 1: " + part1(input));
        System.out.println("Day 9 part 2: " + part2(input));
    }

    private static String readFile(File file) {
        StringBuilder sb = new StringBuilder();
        try (BufferedReader br = new BufferedReader(new FileReader(file))) {
            String line;
            while ((line = br.readLine()) != null) {
                sb.append(line.trim());
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
        return sb.toString();
    }

    private static long part1(String input) {
        return decompress(input, false);
    }

    private static long part2(String input) {
        return decompress(input, true);
    }

    private static long decompress(String input, boolean v2) {
        long res = 0;
        int i = 0;
        while (i < input.length()) {
            if (input.charAt(i) == '(') {
                int end = input.indexOf(')', i);
                String[] marker = input.substring(i + 1, end).split("x");
                int length = Integer.parseInt(marker[0]);
                int times = Integer.parseInt(marker[1]);
                String repeated = input.substring(end + 1, end + 1 + length);
                if (v2) {
                    res += decompress(repeated, true) * times;
                } else {
                    res += repeated.length() * times;
                }
                i = end + 1 + length;
            } else {
                res += 1;
                i++;
            }
        }

        return res;
    }

}
