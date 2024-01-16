package aoc2016.solutions;

import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

import aoc2016.utils.FileLoader;

public class Day25 {
    public static void solve() {
        String input = FileLoader.loadFile("day25.txt");
        System.out.println("Day 25 part 1: " + part1(input));
    }

    private static int part1(String input) {
        State state = new State();
        for (String line : input.split("\n")) {
            String[] tokens = line.split(" ");
            switch (tokens[0]) {
                case "cpy":
                    state.getInstructions().add(new Cpy(readArg(new Scanner(tokens[1])),
                            readRegisterArg(new Scanner(tokens[2]))));
                    break;
                case "inc":
                    state.getInstructions().add(new Inc(readRegisterArg(new Scanner(tokens[1]))));
                    break;
                case "dec":
                    state.getInstructions().add(new Dec(readRegisterArg(new Scanner(tokens[1]))));
                    break;
                case "jnz":
                    state.getInstructions().add(new Jnz(readArg(new Scanner(tokens[1])),
                            readArg(new Scanner(tokens[2]))));
                    break;
                case "out":
                    state.getInstructions().add(new Out(readArg(new Scanner(tokens[1]))));
                    break;
            }
        }

        for (int a = 0;; a++) {
            state.setRegisters(new int[] { a, 0, 0, 0 });
            state.setPc(0);
            state.setOutput(new ArrayList<>());

            while (state.getPc() >= 0 && state.getPc() < state.getInstructions().size()) {
                state.setPc(state.getInstructions().get(state.getPc()).execute(state));
                int n = state.getOutput().size();
                if (n > 100) {
                    return a;
                }
                if (n > 0 && state.getOutput().get(n - 1) != (n - 1) % 2) {
                    break;
                }
            }
        }
    }

    private static RegisterArgument readRegisterArg(Scanner scanner) {
        String arg = scanner.next();
        return new RegisterArgument(arg.charAt(0) - 'a');
    }

    private static Argument readArg(Scanner scanner) {
        String arg = scanner.next();
        if (Character.isLetter(arg.charAt(0))) {
            return new RegisterArgument(arg.charAt(0) - 'a');
        } else {
            long value = Long.parseLong(arg);
            return new IntegerArgument((int) value);
        }
    }

    interface Instruction {
        int execute(State state);
    }

    private static class Cpy implements Instruction {
        private Argument from;
        private RegisterArgument to;

        public Cpy(Argument from, RegisterArgument to) {
            this.from = from;
            this.to = to;
        }

        @Override
        public int execute(State state) {
            state.getRegisters()[to.getRegister()] = from.evaluate(state);
            return state.getPc() + 1;
        }
    }

    private static class Inc implements Instruction {
        private RegisterArgument register;

        public Inc(RegisterArgument register) {
            this.register = register;
        }

        @Override
        public int execute(State state) {
            state.getRegisters()[register.getRegister()]++;
            return state.getPc() + 1;
        }
    }

    private static class Dec implements Instruction {
        private RegisterArgument register;

        public Dec(RegisterArgument register) {
            this.register = register;
        }

        @Override
        public int execute(State state) {
            state.getRegisters()[register.getRegister()]--;
            return state.getPc() + 1;
        }
    }

    private static class Jnz implements Instruction {
        private Argument cond;
        private Argument offset;

        public Jnz(Argument cond, Argument offset) {
            this.cond = cond;
            this.offset = offset;
        }

        @Override
        public int execute(State state) {
            if (cond.evaluate(state) != 0) {
                return state.getPc() + offset.evaluate(state);
            } else {
                return state.getPc() + 1;
            }
        }
    }

    private static class Out implements Instruction {
        private Argument value;

        public Out(Argument value) {
            this.value = value;
        }

        @Override
        public int execute(State state) {
            state.getOutput().add(value.evaluate(state));
            return state.getPc() + 1;
        }
    }

    interface Argument {
        int evaluate(State state);
    }

    private static class IntegerArgument implements Argument {
        private int value;

        public IntegerArgument(int value) {
            this.value = value;
        }

        @Override
        public int evaluate(State state) {
            return value;
        }
    }

    private static class RegisterArgument implements Argument {
        private int register;

        public RegisterArgument(int register) {
            this.register = register;
        }

        @Override
        public int evaluate(State state) {
            return state.getRegisters()[register];
        }

        public int getRegister() {
            return register;
        }
    }

    private static class State {
        private List<Instruction> instructions;
        private int pc;
        private int[] registers;
        private List<Integer> output;

        public State() {
            this.instructions = new ArrayList<>();
            this.pc = 0;
            this.registers = new int[4];
            this.output = new ArrayList<>();
        }

        public List<Instruction> getInstructions() {
            return instructions;
        }

        public int getPc() {
            return pc;
        }

        public int[] getRegisters() {
            return registers;
        }

        public List<Integer> getOutput() {
            return output;
        }

        public void setPc(int pc) {
            this.pc = pc;
        }

        public void setRegisters(int[] registers) {
            this.registers = registers;
        }

        public void setOutput(List<Integer> output) {
            this.output = output;
        }
    }

}
