package aoc2016.solutions;

public class Day16 {
    public static void solve() {
        String initial = "01000100010010111";
        System.out.println("Day 16 part 1: " + part1(initial));
        System.out.println("Day 16 part 2: " + part2(initial));
    }

    private static String dragonsCurve(String init, int maxSize) {
        String a = init;
        while (a.length() < maxSize) {
            // Copy a to b and reverse
            StringBuilder sb = new StringBuilder();
            for (int i = a.length() - 1; i >= 0; i--) {
                sb.append(a.charAt(i) == '1' ? '0' : '1');
            }
            a = a + "0" + sb.toString();
        }

        return a.substring(0, maxSize);
    }

    private static String getCheckSum(String data) {
        while (true) {
            StringBuilder sb = new StringBuilder();
            for (int i = 0; i < data.length() - 1; i += 2) {
                sb.append(data.charAt(i) == data.charAt(i + 1) ? "1" : "0");
            }
            data = sb.toString();

            if (data.length() % 2 == 1) {
                return data;
            }
        }
    }

    private static String part1(String init) {
        String data = dragonsCurve(init, 272);
        String checksum = getCheckSum(data);

        return checksum;
    }

    private static String part2(String init) {
        String data = dragonsCurve(init, 35_651_584);
        String checksum = getCheckSum(data);

        return checksum;
    }
}
