package aoc2016.solutions;

import java.math.BigInteger;
import java.security.MessageDigest;

public class Day5 {
    private static class MD5 {
        private static String getMD5(String input) {
            String hash = "";
            try {
                // Static getInstance method is called with hashing MD5
                MessageDigest md = MessageDigest.getInstance("MD5");

                // digest() method is called to calculate message digest
                // of an input digest() return array of byte
                byte[] messageDigest = md.digest(input.getBytes("UTF-8"));

                // Convert byte array into signum representation
                BigInteger no = new BigInteger(1, messageDigest);

                // Convert message digest into hex value
                String hashtext = no.toString(16);
                while (hashtext.length() < 32) {
                    hashtext = "0" + hashtext;
                }
                hash = hashtext;
            } catch (Exception err) {
                System.out.println("Failed to get MD5 hash");
                System.exit(1);
            }
            return hash;
        }
    }

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
