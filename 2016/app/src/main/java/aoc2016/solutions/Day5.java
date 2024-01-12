package aoc2016.solutions;

import aoc2016.utils.MD5;

public class Day5 {

    public static void solve() {
        System.out.println("Day 5 part 1: " + part1("cxdnnyjw"));
        System.out.println("Day 5 part 2: " + part2("cxdnnyjw"));
    }

    private static String part1(String doorID) {
        String password = "";
        long index = 0;
        while (password.length() < 8) {
            String hash = MD5.getMD5(doorID + index);
            if (hash.startsWith("00000")) {
                password += hash.charAt(5);
            }
            index += 1;

            if (index % 100000 == 0) {
                // Run garbage collection every 100000 iterations
                System.gc();
            }
        }
        return password;
    }

    private static String part2(String doorID) {
        String password = "________";
        long index = 0;
        while (password.contains("_")) {
            String hash = MD5.getMD5(doorID + index);
            if (hash.startsWith("00000")) {
                int pos = Character.getNumericValue(hash.charAt(5));
                if (pos < 8 && password.charAt(pos) == '_') {
                    password = password.substring(0, pos) + hash.charAt(6) + password.substring(pos + 1);
                }
            }
            index += 1;

            if (index % 100000 == 0) {
                // Run garbage collection every 100000 iterations
                System.gc();
            }
        }

        return password;
    }
}
