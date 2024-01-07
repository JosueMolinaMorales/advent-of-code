package aoc2016.solutions;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;

public class Day3 {
    public static void solve() {
        ClassLoader cl = Day3.class.getClassLoader();
        File file = new File(cl.getResource("day3.txt").getFile());

        System.out.println("Day 3 part 1: " + part1(file));
        System.out.println("Day 3 part 2: " + part2(file));
    }

    private static ArrayList<int[]> parse(File file) {
        ArrayList<int[]> triangles = new ArrayList<>();
        try (BufferedReader br = new BufferedReader(new FileReader(file))) {
            String input;
            while ((input = br.readLine()) != null) {
                int[] triangle = new int[3];
                String[] sides = input.trim().split(" ");
                int pos = 0;
                for (int i = 0; i < sides.length; i++) {
                    if (sides[i].trim().isEmpty()) {
                        continue;
                    }
                    triangle[pos] = Integer.parseInt(sides[i]);
                    pos += 1;
                }
                triangles.add(triangle);
            }
        } catch (IOException err) {
            System.out.println("Failed to read day 3 file");
            System.exit(1);
        }
        return triangles;
    }

    private static boolean isTriangleValid(int[] triangle) {
        int a = triangle[0];
        int b = triangle[1];
        int c = triangle[2];
        return a + b > c && a + c > b && b + c > a;
    }

    private static int part1(File file) {
        ArrayList<int[]> triangles = parse(file);
        return (int) triangles.stream().filter(triangle -> isTriangleValid(triangle)).count();
    }

    private static int part2(File file) {
        ArrayList<int[]> triangles = parse(file);
        ArrayList<int[]> verticalTriangles = new ArrayList<>();
        for (int i = 0; i < triangles.size(); i += 3) {
            int[] triangle1 = triangles.get(i);
            int[] triangle2 = triangles.get(i + 1);
            int[] triangle3 = triangles.get(i + 2);
            verticalTriangles.add(new int[] { triangle1[0], triangle2[0], triangle3[0] });
            verticalTriangles.add(new int[] { triangle1[1], triangle2[1], triangle3[1] });
            verticalTriangles.add(new int[] { triangle1[2], triangle2[2], triangle3[2] });
        }

        return (int) verticalTriangles.stream().filter(triangle -> isTriangleValid(triangle)).count();
    }
}
