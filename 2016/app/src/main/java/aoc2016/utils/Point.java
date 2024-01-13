package aoc2016.utils;

public class Point {
    public int x;
    public int y;

    public Point() {
        x = 0;
        y = 0;
    }

    public Point(int x, int y) {
        this.x = x;
        this.y = y;
    }

    public int manhattanDistance(Point other) {
        return Math.abs(x - other.x) + Math.abs(y - other.y);
    }

    public Direction toDirection() {
        if (x == -1 && y == 0) {
            return Direction.UP;
        } else if (x == 1 && y == 0) {
            return Direction.DOWN;
        } else if (x == 0 && y == -1) {
            return Direction.LEFT;
        } else if (x == 0 && y == 1) {
            return Direction.RIGHT;
        }
        return null;
    }

    @Override
    public boolean equals(Object other) {
        if (!(other instanceof Point)) {
            return false;
        }
        Point otherPoint = (Point) other;
        return x == otherPoint.x && y == otherPoint.y;
    }

    @Override
    public String toString() {
        return "(" + x + ", " + y + ")";
    }

    @Override
    public int hashCode() {
        // https://stackoverflow.com/questions/113511/hash-code-implementation
        int result = 17;
        result = 31 * result + x;
        result = 31 * result + y;
        return result;
    }
}
