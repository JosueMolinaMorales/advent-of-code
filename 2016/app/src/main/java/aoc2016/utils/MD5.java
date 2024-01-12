package aoc2016.utils;

import java.math.BigInteger;
import java.security.MessageDigest;

public class MD5 {
    public static String getMD5(String input) {
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
