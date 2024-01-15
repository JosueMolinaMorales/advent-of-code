package aoc2016.solutions;

import java.util.ArrayList;

import aoc2016.utils.FileLoader;

public class Day21 {
    public static void solve() {
        String password = "abcdefgh";
        String scrambled = "fbgdceah";
        String input = FileLoader.loadFile("day21.txt");
        System.out.println("Day 21 part 1: " + part1(password, input));
        System.out.println("Day 21 part 2: " + part2(scrambled, input));
    }

    private static String part1(String password, String input) {
        ArrayList<Instruction> instructions = parse(input);
        for (Instruction inst : instructions) {
            password = inst.apply(password);
        }
        return password;
    }

    private static String part2(String password, String input) {
        ArrayList<Instruction> instructions = parse(input);
        for (int i = instructions.size() - 1; i >= 0; i--) {
            password = instructions.get(i).applyReverse(password);
        }
        return password;
    }

    private static ArrayList<Instruction> parse(String input) {
        ArrayList<Instruction> instructions = new ArrayList<>();
        for (String line : input.split("\n")) {
            instructions.add(new Instruction(line));
        }
        return instructions;
    }

    private static enum Operation {
        SWAP_POSITION, SWAP_LETTER, ROTATE_LEFT, ROTATE_RIGHT, ROTATE_BASED, REVERSE, MOVE;
    }

    private static class Instruction {
        Operation operation;
        int arg1;
        int arg2;

        Instruction(String line) {
            String[] parts = line.split(" ");
            switch (parts[0]) {
                case "swap":
                    if (parts[1].equals("position")) {
                        operation = Operation.SWAP_POSITION;
                        arg1 = Integer.parseInt(parts[2]);
                        arg2 = Integer.parseInt(parts[5]);
                    } else {
                        operation = Operation.SWAP_LETTER;
                        arg1 = parts[2].charAt(0);
                        arg2 = parts[5].charAt(0);
                    }
                    break;
                case "rotate":
                    if (parts[1].equals("left")) {
                        operation = Operation.ROTATE_LEFT;
                        arg1 = Integer.parseInt(parts[2]);
                    } else if (parts[1].equals("right")) {
                        operation = Operation.ROTATE_RIGHT;
                        arg1 = Integer.parseInt(parts[2]);
                    } else {
                        operation = Operation.ROTATE_BASED;
                        arg1 = parts[6].charAt(0);
                    }
                    break;
                case "reverse":
                    operation = Operation.REVERSE;
                    arg1 = Integer.parseInt(parts[2]);
                    arg2 = Integer.parseInt(parts[4]);
                    break;
                case "move":
                    operation = Operation.MOVE;
                    arg1 = Integer.parseInt(parts[2]);
                    arg2 = Integer.parseInt(parts[5]);
                    break;
            }
        }

        String apply(String passcode) {
            switch (operation) {
                case SWAP_POSITION:
                    return swapPosition(passcode, arg1, arg2);
                case SWAP_LETTER:
                    return swapLetter(passcode, arg1, arg2);
                case ROTATE_LEFT:
                    return rotateLeft(passcode, arg1);
                case ROTATE_RIGHT:
                    return rotateRight(passcode, arg1);
                case ROTATE_BASED:
                    return rotateBased(passcode, arg1);
                case REVERSE:
                    return reverse(passcode, arg1, arg2);
                case MOVE:
                    return move(passcode, arg1, arg2);
            }
            return null;
        }

        String applyReverse(String passcode) {
            switch (operation) {
                case SWAP_POSITION:
                    return swapPosition(passcode, arg1, arg2);
                case SWAP_LETTER:
                    return swapLetter(passcode, arg1, arg2);
                case ROTATE_LEFT:
                    return rotateRight(passcode, arg1);
                case ROTATE_RIGHT:
                    return rotateLeft(passcode, arg1);
                case ROTATE_BASED:
                    return rotateBasedReverse(passcode, arg1);
                case REVERSE:
                    return reverse(passcode, arg1, arg2);
                case MOVE:
                    return move(passcode, arg2, arg1);
            }
            return null;
        }

        private String swapPosition(String passcode, int pos1, int pos2) {
            char[] chars = passcode.toCharArray();
            char temp = chars[pos1];
            chars[pos1] = chars[pos2];
            chars[pos2] = temp;
            return new String(chars);
        }

        private String swapLetter(String passcode, int letter1, int letter2) {
            char[] chars = passcode.toCharArray();
            for (int i = 0; i < chars.length; i++) {
                if (chars[i] == letter1) {
                    chars[i] = (char) letter2;
                } else if (chars[i] == letter2) {
                    chars[i] = (char) letter1;
                }
            }
            return new String(chars);
        }

        private String rotateLeft(String passcode, int steps) {
            char[] chars = passcode.toCharArray();
            char[] newChars = new char[chars.length];
            for (int i = 0; i < chars.length; i++) {
                newChars[i] = chars[(i + chars.length + steps) % chars.length];
            }
            return new String(newChars);
        }

        private String rotateRight(String passcode, int steps) {
            char[] chars = passcode.toCharArray();
            char[] newChars = new char[chars.length];
            for (int i = 0; i < chars.length; i++) {
                newChars[(i + steps) % chars.length] = chars[i];
            }
            return new String(newChars);
        }

        private String rotateBased(String passcode, int letter) {
            int pos = passcode.indexOf(letter);
            int steps = 1 + pos;
            if (pos >= 4) {
                steps += 1;
            }
            return rotateRight(passcode, steps);
        }

        private String rotateBasedReverse(String passcode, int letter) {
            int pos = passcode.indexOf(letter);
            int steps = 0;
            switch (pos) {
                case 0:
                    steps = 1;
                    break;
                case 1:
                    steps = 1;
                    break;
                case 2:
                    steps = 6;
                    break;
                case 3:
                    steps = 2;
                    break;
                case 4:
                    steps = 7;
                    break;
                case 5:
                    steps = 3;
                    break;
                case 6:
                    steps = 0;
                    break;
                case 7:
                    steps = 4;
                    break;
            }
            return rotateLeft(passcode, steps);
        }

        private String reverse(String passcode, int pos1, int pos2) {
            char[] chars = passcode.toCharArray();
            char[] newChars = new char[chars.length];
            for (int i = 0; i < chars.length; i++) {
                if (i >= pos1 && i <= pos2) {
                    newChars[i] = chars[pos2 - (i - pos1)];
                } else {
                    newChars[i] = chars[i];
                }
            }
            return new String(newChars);
        }

        private String move(String passcode, int pos1, int pos2) {
            // Convert to arraylist for easier manipulation
            ArrayList<Character> chars = new ArrayList<>();
            for (char ch : passcode.toCharArray()) {
                chars.add(ch);
            }
            chars.remove(pos1);
            chars.add(pos2, passcode.charAt(pos1));

            // Convert back to string
            StringBuilder sb = new StringBuilder();
            for (char ch : chars) {
                sb.append(ch);
            }
            return sb.toString();

        }

    }
}
