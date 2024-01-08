package aoc2016.solutions;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;

public class Day7 {
    public static void solve() {
        ClassLoader classLoader = Day7.class.getClassLoader();
        File file = new File(classLoader.getResource("day7.txt").getFile());

        System.out.println("Day 7 part 1: " + part1(file));
        System.out.println("Day 7 part 2: " + part2(file));
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
            System.out.println("Failed to read file for day 7");
            System.exit(1);
        }
        return null;
    }

    private static boolean supportsSSL(String ip) {
        boolean inBrackets = false;
        ArrayList<String> aba = new ArrayList<>();
        ArrayList<String> bab = new ArrayList<>();
        for (int i = 0; i < ip.length() - 2; i++) {
            char ch = ip.charAt(i);
            if (ch == '[') {
                inBrackets = true;
                continue;
            } else if (ch == ']') {
                inBrackets = false;
                continue;
            }

            if (ip.charAt(i) != ip.charAt(i + 1) && ip.charAt(i) == ip.charAt(i + 2)) {
                if (inBrackets) {
                    bab.add(ip.substring(i, i + 3));
                } else {
                    aba.add(ip.substring(i, i + 3));
                }
            }
        }

        for (String a : aba) {
            for (String b : bab) {
                if (a.charAt(0) == b.charAt(1) && a.charAt(1) == b.charAt(0)) {
                    return true;
                }
            }
        }
        return false;
    }

    private static boolean supportsTLS(String ip) {
        boolean inBrackets = false;
        boolean hasABBA = false;
        for (int i = 0; i < ip.length() - 3; i++) {
            char ch = ip.charAt(i);
            if (ch == '[') {
                inBrackets = true;
                continue;
            } else if (ch == ']') {
                inBrackets = false;
                continue;
            }

            if (ip.charAt(i) != ip.charAt(i + 1) && ip.charAt(i) == ip.charAt(i + 3)
                    && ip.charAt(i + 1) == ip.charAt(i + 2)) {
                if (inBrackets) {
                    return false;
                }
                hasABBA = true;

            }
        }
        return hasABBA;
    }

    private static int part1(File file) {
        ArrayList<String> ips = parse(file);
        return (int) ips.stream().filter(ip -> supportsTLS(ip)).count();
    }

    private static int part2(File file) {
        ArrayList<String> ips = parse(file);
        return (int) ips.stream().filter(ip -> supportsSSL(ip)).count();
    }
}
