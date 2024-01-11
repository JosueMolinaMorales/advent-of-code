package aoc2016.solutions;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.PriorityQueue;

import aoc2016.utils.FileLoader;

public class Day11 {
    private static enum ItemType {
        GENERATOR, MICROCHIP
    }

    private static class Item {
        ItemType type;
        String name;

        public Item(ItemType type, String name) {
            this.type = type;
            this.name = name;
        }

        @Override
        public String toString() {
            return this.name + " " + this.type;
        }
    }

    private static class Floor {
        ArrayList<Item> items;

        public Floor() {
            this.items = new ArrayList<>();
        }

        public void addItem(Item item) {
            this.items.add(item);
        }

        public void removeItem(Item item) {
            this.items.remove(item);
        }

        /**
         * Returns true if the floor is valid, false otherwise.
         * A floor is valid if it does not contain a microchip without its corresponding
         * generator.
         * 
         * @return true if the floor is valid, false otherwise.
         */
        public boolean isValidFloor() {
            ArrayList<String> generators = new ArrayList<>();
            ArrayList<String> microchips = new ArrayList<>();

            for (int i = 0; i < this.items.size(); i++) {
                if (this.items.get(i).type == ItemType.GENERATOR) {
                    generators.add(this.items.get(i).name);
                } else {
                    microchips.add(this.items.get(i).name);
                }
            }

            for (int i = 0; i < microchips.size(); i++) {
                if (!generators.contains(microchips.get(i))) {
                    return false;
                }
            }

            return true;
        }

        @Override
        public String toString() {
            StringBuilder sb = new StringBuilder();
            for (int i = 0; i < this.items.size(); i++) {
                sb.append(this.items.get(i) + " ");
            }
            return sb.toString();
        }
    }

    public static void solve() {
        String input = FileLoader.loadFile("day11.txt");
        System.out.println("Day 11 part 1: " + part1(input));
        System.out.println("Day 11 part 2: " + part2(input));
    }

    private static HashMap<Integer, Floor> parseInput(String input) {
        HashMap<Integer, Floor> items = new HashMap<>();
        String[] lines = input.replaceAll("-compatible", "").replaceAll(",", "").replace(".", "").split("\n");
        for (int i = 0; i < lines.length; i++) {
            System.out.println(lines[i]);
            String[] words = lines[i].split(" ");
            for (int j = 0; j < words.length; j++) {
                if (!items.containsKey(i)) {
                    items.put(i, new Floor());
                }
                if (words[j].equals("generator")) {
                    items.get(i).addItem(new Item(ItemType.GENERATOR, words[j - 1]));
                } else if (words[j].equals("microchip")) {
                    items.get(i).addItem(new Item(ItemType.MICROCHIP, words[j - 1]));
                }
            }
        }

        // Add empty floor
        items.put(lines.length, new Floor());
        return items;
    }

    private static class State {
        HashMap<Integer, Floor> floors;
        int elevatorFloor;
        int steps;

        public State(HashMap<Integer, Floor> floors, int elevatorFloor, int steps) {
            this.floors = floors;
            this.elevatorFloor = elevatorFloor;
            this.steps = steps;
        }
    }

    public static int part1(String input) {
        HashMap<Integer, Floor> floors = parseInput(input);
        for (int i = 0; i < floors.size(); i++) {
            System.out.println("Floor " + i + ": " + floors.get(i));
        }
        // Use dijkstra's algorithm to find the minimum number of steps to move all
        // items to the top floor.
        PriorityQueue<State> queue = new PriorityQueue<>((a, b) -> a.steps - b.steps);
        queue.add(new State(floors, 0, 0));

        while (!queue.isEmpty()) {
            State state = queue.poll();
            System.out.println("State: " + state.elevatorFloor + " " + state.steps);
            if (state.elevatorFloor == 3) {
                return state.steps;
            }

        }
        return 0;
    }

    public static int part2(String input) {
        return 0;
    }
}
