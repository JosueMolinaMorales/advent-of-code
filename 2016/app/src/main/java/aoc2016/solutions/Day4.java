package aoc2016.solutions;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.HashMap;

public class Day4 {
    private static class Room {
        String name;
        int sectorId;
        String checksum;

        public Room(String input) {
            String[] parts = input.split("-");
            String name = "";
            for (int i = 0; i < parts.length - 1; i++) {
                name += parts[i] + "-";
            }
            this.name = name;

            String[] lastParts = parts[parts.length - 1].split("\\[");
            this.sectorId = Integer.parseInt(lastParts[0]);
            this.checksum = lastParts[1].substring(0, lastParts[1].length() - 1);
        }
    }

    public static void solve() {
        ClassLoader cl = Day4.class.getClassLoader();
        File file = new File(cl.getResource("day4.txt").getFile());

        // 8265 too low
        System.out.println("Day 4 part 1: " + part1(file));
        System.out.println("Day 4 part 2: " + part2(file));
    }

    private static ArrayList<Room> parse(File file) {
        ArrayList<Room> rooms = new ArrayList<>();
        try (BufferedReader br = new BufferedReader(new FileReader(file))) {
            String input;
            while ((input = br.readLine()) != null) {
                rooms.add(new Room(input));
            }
        } catch (IOException err) {
            System.out.println("Failed to read day 4 file");
            System.exit(1);
        }

        return rooms;
    }

    private static boolean isRoomValid(Room room) {
        // Key: character, Value: (count, index)
        HashMap<Character, Integer> counts = new HashMap<>();
        for (int i = 0; i < room.name.length(); i++) {
            char ch = room.name.charAt(i);
            if (ch == '-') {
                continue;
            }
            if (counts.containsKey(ch)) {
                counts.put(ch, counts.get(ch) + 1);
            } else {
                counts.put(ch, 1);
            }
        }

        String checksum = "";
        for (int i = 0; i < 5; i++) {
            int max = 0;
            char maxChar = ' ';
            for (char ch : counts.keySet()) {
                if (counts.get(ch) > max) {
                    max = counts.get(ch);
                    maxChar = ch;
                } else if (counts.get(ch) == max && ch < maxChar) {
                    max = counts.get(ch);
                    maxChar = ch;
                }
            }
            checksum += maxChar;
            counts.remove(maxChar);
        }
        return checksum.equals(room.checksum);
    }

    private static int part1(File file) {
        ArrayList<Room> rooms = parse(file);
        return rooms.stream().filter(room -> isRoomValid(room)).mapToInt(room -> room.sectorId).sum();
    }

    private static int part2(File file) {
        ArrayList<Room> rooms = parse(file);
        for (Room room : rooms) {
            String decrypted = decrypt(room.name, room.sectorId);
            if (decrypted.contains("north")) {
                return room.sectorId;
            }
        }
        return -1;
    }

    private static String decrypt(String name, int sectorId) {
        String decrypted = "";
        for (int i = 0; i < name.length(); i++) {
            char ch = name.charAt(i);
            if (ch == '-') {
                decrypted += " ";
                continue;
            }
            int chInt = (int) ch;
            chInt -= 97;
            chInt += sectorId;
            chInt %= 26;
            chInt += 97;
            decrypted += (char) chInt;
        }
        return decrypted;
    }
}
