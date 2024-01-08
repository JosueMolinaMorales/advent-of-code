package aoc2016.solutions;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;

public class Day6 {
    public static void solve() {
        ClassLoader classLoader = Day6.class.getClassLoader();
        File file = new File(classLoader.getResource("day6.txt").getFile());

        System.out.println("Day 6 part 1: " + part1(file));
        System.out.println("Day 6 part 2: " + part2(file));
    }

    private static ArrayList<String> parse(File file) {
        try (BufferedReader br = new BufferedReader(new FileReader(file))) {
            ArrayList<String> messages = new ArrayList<String>();
            String line;
            while ((line = br.readLine()) != null) {
                messages.add(line);
            }
            return messages;
        } catch (IOException err) {
            System.out.println("Failed to read file for day 6");
            System.exit(1);
        }
        return null;
    }

    private static String part1(File file) {
        return solve(file, false);
    }

    private static String part2(File file) {
        return solve(file, true);
    }

    private static String solve(File file, boolean getMin) {
        ArrayList<String> messages = parse(file);
        int col = 0;
        String rm = "";
        while (col < messages.get(col).length()) {
            HashMap<Character, Integer> count = new HashMap<>();
            for (String message : messages) {
                char ch = message.charAt(col);
                if (count.containsKey(ch)) {
                    count.put(ch, count.get(ch) + 1);
                } else {
                    count.put(ch, 1);
                }
            }

            if (getMin) {
                rm += getMinChar(count);
            } else {
                rm += getMaxChar(count);
            }
            col += 1;
        }

        return rm;
    }

    private static char getMinChar(HashMap<Character, Integer> count) {
        // Get the max character
        int min = Integer.MAX_VALUE;
        char minChar = ' ';
        for (Map.Entry<Character, Integer> set : count.entrySet()) {
            if (set.getValue() < min) {
                min = set.getValue();
                minChar = set.getKey();
            }
        }
        return minChar;
    }

    private static char getMaxChar(HashMap<Character, Integer> count) {
        // Get the max character
        int max = 0;
        char maxChar = ' ';
        for (Map.Entry<Character, Integer> set : count.entrySet()) {
            if (set.getValue() > max) {
                max = set.getValue();
                maxChar = set.getKey();
            }
        }
        return maxChar;
    }

}
