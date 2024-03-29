package aoc2016.utils;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.net.URL;

public class FileLoader {
    public static String loadFile(String filename) {
        ClassLoader classLoader = FileLoader.class.getClassLoader();
        URL url = classLoader.getResource(filename);
        if (url == null) {
            throw new RuntimeException("File not found: " + filename);
        }
        File file = new File(classLoader.getResource(filename).getFile());
        return readFile(file);
    }

    private static String readFile(File file) {
        StringBuilder sb = new StringBuilder();
        try (BufferedReader br = new BufferedReader(new FileReader(file))) {
            String line;
            while ((line = br.readLine()) != null) {
                sb.append(line + "\n");
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
        return sb.toString();
    }
}
