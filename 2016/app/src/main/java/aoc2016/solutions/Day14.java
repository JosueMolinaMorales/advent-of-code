package aoc2016.solutions;

import java.util.ArrayList;
import java.util.HashMap;

import aoc2016.utils.MD5;

public class Day14 {
    public static void solve() {
        String salt = "jlmsuwbz";

        System.out.println("Day 14 part 1: " + part1(salt));
        System.out.println("Day 14 part 2: " + part2(salt));
    }

    private static class PotentialKey {
        char repeated;
        int count;

        PotentialKey(char repeated, int count) {
            this.repeated = repeated;
            this.count = count;
        }

        public boolean hasRepeatedCharacter(String hash) {
            int count = 0;
            for (int i = 0; i < hash.length(); i++) {
                char ch = hash.charAt(i);
                if (ch == repeated) {
                    count += 1;
                } else {
                    count = 0;
                }
                if (count == 5) {
                    break;
                }
            }

            return count == 5;
        }
    }

    private static int getThreeInRow(String hash) {
        int count = 1;
        int lastChar = hash.charAt(0);
        for (int i = 1; i < hash.length(); i++) {
            char ch = hash.charAt(i);
            if (ch == lastChar) {
                count += 1;
            } else {
                lastChar = ch;
                count = 1;
            }

            if (count == 3) {
                return ch;
            }
        }

        return -1;
    }

    private static int solve(String salt, int hashCount) {
        HashMap<Integer, Integer> keys = new HashMap<>();
        ArrayList<PotentialKey> pKeys = new ArrayList<>();
        int i = 0;

        while (keys.size() < 64) {
            String hash = MD5.getMD5(salt + i);
            for (int j = 0; j < hashCount; j++) {
                hash = MD5.getMD5(hash);
            }
            ArrayList<PotentialKey> keysToRemove = new ArrayList<>();
            for (PotentialKey pk : pKeys) {
                if ((i - pk.count) >= 1000) {
                    keysToRemove.add(pk);
                    continue;
                }
                // Check for each potential key, if the current hash
                // has its repeated key 5 times
                // If so, add it to keys, else inc count
                if (pk.hasRepeatedCharacter(hash)) {
                    keys.put(keys.size(), pk.count);
                    keysToRemove.add(pk);
                }

            }

            pKeys.removeAll(keysToRemove);

            // Check to see if the current key is a potential key, ie it has the same
            // Character 3 times
            int repeated = getThreeInRow(hash);
            if (repeated != -1) {
                pKeys.add(new PotentialKey((char) repeated, i));
            }
            i++;
        }

        return keys.get(63);

    }

    private static int part1(String salt) {
        return solve(salt, 0);
    }

    private static int part2(String salt) {
        return solve(salt, 2016);
    }
}
