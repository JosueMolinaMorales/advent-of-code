package aoc2016.solutions;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;

import aoc2016.utils.Point;

public class Day2 {
    public static void solve() {
        ClassLoader cl = Day2.class.getClassLoader();
        File file = new File(cl.getResource("day2.txt").getFile());

        System.out.println("Day 2 part 1: " + part1(file));
        System.out.println("Day 2 part 2: " + part2(file));
    }

    final static String[][] keypad = {
            { "1", "2", "3" },
            { "4", "5", "6" },
            { "7", "8", "9" }
    };

    final static String[][] bathroomKeypad = {
            { null, null, "1", null, null },
            { null, "2", "3", "4", null },
            { "5", "6", "7", "8", "9" },
            { null, "A", "B", "C", null },
            { null, null, "D", null, null },
    };

    private static ArrayList<String> parse(File file) {
        try (BufferedReader br = new BufferedReader(new FileReader(file))) {
            ArrayList<String> instructions = new ArrayList<String>();
            String line;
            while ((line = br.readLine()) != null) {
                instructions.add(line);
            }
            return instructions;
        } catch (IOException err) {
            System.out.println("Failed to read file for day 2");
            System.exit(1);
        }
        return null;
    }

    private static String solve(File file, String[][] keypad, Point start) {
        ArrayList<String> instructions = parse(file);
        String code = "";
        Point point = start;
        for (String line : instructions) {
            for (int i = 0; i < line.length(); i++) {
                char ch = line.charAt(i);
                point = move(keypad, point, ch);
            }
            code += keypad[point.x][point.y];

        }
        return code;

    }

    private static String part1(File file) {
        return solve(file, keypad, new Point(1, 1));
    }

    private static String part2(File file) {
        return solve(file, bathroomKeypad, new Point(2, 0));
    }

    private static Point move(String[][] keypad, Point cp, char dir) {
        Point np = new Point(cp.x, cp.y);
        switch (dir) {
            case 'U':
                np.x -= 1;
                break;
            case 'D':
                np.x += 1;
                break;
            case 'L':
                np.y -= 1;
                break;
            case 'R':
                np.y += 1;
                break;
        }
        if (np.x < 0 || np.x >= keypad.length || np.y < 0 || np.y >= keypad[0].length || keypad[np.x][np.y] == null) {
            return cp;
        }
        return np;
    }
}
