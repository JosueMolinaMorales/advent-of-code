package aoc2016.solutions;

import java.util.HashMap;
import java.util.LinkedList;
import java.util.Queue;

import aoc2016.utils.FileLoader;

public class Day10 {
    private static class Bot {
        // The bot's id
        int id;
        // The low and high values
        int low;
        int high;
        // The low and high bot ids
        int lowId;
        int highId;

        boolean lowOutput;
        boolean highOutput;

        public Bot(int id) {
            this.id = id;
            this.low = -1;
            this.high = -1;
            this.lowId = -1;
            this.highId = -1;
            this.lowOutput = false;
            this.highOutput = false;
        }

        public void setChip(int value) {
            if (this.low == -1) {
                this.low = value;
                return;
            }

            if (value < this.low) {
                this.high = this.low;
                this.low = value;
            } else {
                this.high = value;
            }
        }

        public void setLowId(int id) {
            this.lowId = id;
        }

        public void setHighId(int id) {
            this.highId = id;
        }

        @Override
        public String toString() {
            return "Bot " + this.id + " low " + this.low + " high " + this.high + " lowId " + this.lowId + " highId "
                    + this.highId + " lowOutput " + this.lowOutput + " highOutput " + this.highOutput;
        }
    }

    public static void solve() {
        String input = FileLoader.loadFile("day10.txt");
        System.out.println("Day 10 part 1: " + part1(input));
        System.out.println("Day 10 part 2: " + part2(input));
    }

    private static int solve(String input, boolean part2) {
        String[] instructions = input.split("\n");
        HashMap<Integer, Bot> bots = new HashMap<>();

        for (String line : instructions) {
            String[] tokens = line.split(" ");
            if (tokens[0].equals("value")) {
                int value = Integer.parseInt(tokens[1]);
                int botId = Integer.parseInt(tokens[5]);
                if (!bots.containsKey(botId)) {
                    bots.put(botId, new Bot(botId));
                }
                Bot bot = bots.get(botId);
                bot.setChip(value);
            } else {
                int botId = Integer.parseInt(tokens[1]);
                if (!bots.containsKey(botId)) {
                    bots.put(botId, new Bot(botId));
                }
                Bot bot = bots.get(botId);
                int lowId = Integer.parseInt(tokens[6]);
                bot.setLowId(lowId);
                if (!bots.containsKey(lowId)) {
                    bots.put(lowId, new Bot(lowId));
                }

                int highId = Integer.parseInt(tokens[11]);
                bot.setHighId(highId);
                if (!bots.containsKey(highId)) {
                    bots.put(highId, new Bot(highId));
                }

                if (tokens[5].equals("output")) {
                    bot.lowOutput = true;
                }

                if (tokens[10].equals("output")) {
                    bot.highOutput = true;
                }
            }
        }

        // Create a queue of bots with 2 chips
        Queue<Bot> queue = new LinkedList<>();
        boolean found = false;
        HashMap<Integer, Integer> outputs = new HashMap<>();
        while (!found) {
            // Find a bot with 2 chips
            for (Bot bot : bots.values()) {
                if (bot.low != -1 && bot.high != -1) {
                    queue.add(bot);
                }
            }
            // Process the queue
            while (!queue.isEmpty()) {
                Bot bot = queue.poll();
                if (!part2 && bot.low == 17 && bot.high == 61) {
                    return bot.id;
                }
                if (bot.lowOutput) {
                    outputs.put(bot.lowId, bot.low);
                } else {
                    Bot lowBot = bots.get(bot.lowId);
                    lowBot.setChip(bot.low);
                }
                if (bot.highOutput) {
                    outputs.put(bot.highId, bot.high);
                } else {
                    Bot highBot = bots.get(bot.highId);
                    highBot.setChip(bot.high);
                }
                bot.low = -1;
                bot.high = -1;
            }
            queue.clear();

            found = true;
            // Check if outputs 0, 1, and 2 have values
            for (int i = 0; i < 3; i++) {
                if (!outputs.containsKey(i)) {
                    found = false;
                    break;
                }
            }
        }

        int result = 1;
        for (int i = 0; i < 3; i++) {
            result *= outputs.get(i);
        }
        return result;
    }

    private static long part1(String input) {
        return solve(input, false);
    }

    private static long part2(String input) {
        return solve(input, true);
    }
}
