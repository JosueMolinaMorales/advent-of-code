package aoc2016.solutions;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.HashSet;
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
        public boolean equals(Object obj) {
            if (!(obj instanceof Item)) {
                return false;
            }
            Item other = (Item) obj;
            return this.type == other.type && this.name.equals(other.name);
        }

        @Override
        public int hashCode() {
            return this.name.hashCode() + this.type.hashCode();
        }
    }

    private static class Floor {
        HashSet<Item> items;

        public Floor() {
            this.items = new HashSet<>();
        }

        public void addItem(Item item) {
            this.items.add(item);
        }

        public void removeItem(Item item) {
            this.items.remove(item);
        }

        @Override
        public int hashCode() {
            return this.items.hashCode();
        }

        public boolean isValidFloor() {
            if (this.items.size() == 0) {
                return true;
            }
            ArrayList<Item> generators = new ArrayList<>();
            ArrayList<Item> microchips = new ArrayList<>();
            for (Item item : this.items) {
                if (item.type == ItemType.GENERATOR) {
                    generators.add(item);
                } else if (item.type == ItemType.MICROCHIP) {
                    microchips.add(item);
                }
            }
            if (generators.size() == 0 || microchips.size() == 0) {
                return true;
            }

            for (Item microchip : microchips) {
                boolean found = false;
                for (Item generator : generators) {
                    if (microchip.name.equals(generator.name)) {
                        found = true;
                        break;
                    }
                }
                if (!found) {
                    return false;
                }
            }

            return true;

        }

        /**
         * Make a deep copy of the floor.
         * 
         * @return a deep copy of the floor.
         */
        public Floor copy() {
            Floor newFloor = new Floor();
            for (Item item : this.items) {
                newFloor.addItem(new Item(item.type, item.name));
            }
            return newFloor;
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

    private static int dijkstra(HashMap<Integer, Floor> floors) {
        // Use dijkstra's algorithm to find the minimum number of steps to move all
        // items to the top floor.
        PriorityQueue<State> queue = new PriorityQueue<>((a, b) -> a.steps - b.steps);
        queue.add(new State(floors, 0, 0));
        HashSet<String> visited = new HashSet<>();
        while (!queue.isEmpty()) {
            State state = queue.poll();
            if (isAllOnTopFloor(state.floors)) {
                return state.steps;
            }

            int currentFloor = state.elevatorFloor;
            HashSet<HashSet<Item>> combinations = generateCombinations(state.floors.get(currentFloor).items);

            int[] directions = { -1, 1 };
            for (HashSet<Item> possibleMove : combinations) {
                for (int direction : directions) {
                    int newFloor = currentFloor + direction;
                    if (newFloor < 0 || newFloor >= floors.size()) {
                        continue;
                    }

                    HashMap<Integer, Floor> newFloors = new HashMap<>();
                    for (int i = 0; i < floors.size(); i++) {
                        newFloors.put(i, state.floors.get(i).copy());
                    }
                    for (Item item : possibleMove) {
                        newFloors.get(currentFloor).removeItem(item);
                        newFloors.get(newFloor).addItem(item);
                    }

                    if (newFloors.get(currentFloor).isValidFloor() && newFloors.get(newFloor).isValidFloor()) {
                        State newState = new State(newFloors, newFloor, state.steps + 1);
                        String key = generateKey(newState);
                        if (visited.contains(key)) {
                            continue;
                        }
                        visited.add(key);
                        queue.add(newState);
                    }
                }
            }
        }
        return -1;
    }

    public static int part1(String input) {
        HashMap<Integer, Floor> floors = parseInput(input);

        return dijkstra(floors);
    }

    private static boolean isAllOnTopFloor(HashMap<Integer, Floor> floors) {
        for (int i = 0; i < floors.size() - 1; i++) {
            if (floors.get(i).items.size() > 0) {
                return false;
            }
        }
        return true;
    }

    private static HashSet<HashSet<Item>> generateCombinations(HashSet<Item> items) {
        HashSet<HashSet<Item>> combinations = new HashSet<>();
        // Add all single item combinations
        for (Item item : items) {
            HashSet<Item> combination = new HashSet<>();
            combination.add(item);
            combinations.add(combination);
        }
        // Add all double item combinations
        for (Item item1 : items) {
            for (Item item2 : items) {
                if (item1.equals(item2)) {
                    continue;
                }
                HashSet<Item> combination = new HashSet<>();
                combination.add(item1);
                combination.add(item2);
                combinations.add(combination);
            }
        }
        return combinations;
    }

    private static String generateKey(State state) {
        StringBuilder sb = new StringBuilder();
        sb.append(state.elevatorFloor);
        for (int i = 0; i < state.floors.size(); i++) {
            int generators = 0;
            int microchips = 0;
            for (Item item : state.floors.get(i).items) {
                if (item.type == ItemType.GENERATOR) {
                    generators++;
                } else if (item.type == ItemType.MICROCHIP) {
                    microchips++;
                }
            }
            sb.append(i + "" + generators + "" + microchips);
        }

        return sb.toString();
    }

    public static int part2(String input) {
        HashMap<Integer, Floor> floors = parseInput(input);
        floors.get(0).addItem(new Item(ItemType.GENERATOR, "elerium"));
        floors.get(0).addItem(new Item(ItemType.MICROCHIP, "elerium"));
        floors.get(0).addItem(new Item(ItemType.GENERATOR, "dilithium"));
        floors.get(0).addItem(new Item(ItemType.MICROCHIP, "dilithium"));

        return dijkstra(floors);
    }
}
