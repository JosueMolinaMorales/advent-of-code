package aoc2016.utils;

import java.util.ArrayList;

public enum Direction {
    UP, DOWN, LEFT, RIGHT, UP_LEFT, UP_RIGHT, DOWN_LEFT, DOWN_RIGHT;

    public static ArrayList<Point> getNonDiagonalDirections() {
        ArrayList<Point> points = new ArrayList<>();
        points.add(Direction.UP.toPoint());
        points.add(Direction.DOWN.toPoint());
        points.add(Direction.LEFT.toPoint());
        points.add(Direction.RIGHT.toPoint());
        return points;
    }

    public Point toPoint() {
        switch (this) {
            case Direction.UP:
                return new Point(-1, 0);
            case Direction.DOWN:
                return new Point(1, 0);
            case Direction.LEFT:
                return new Point(0, -1);
            case Direction.RIGHT:
                return new Point(0, 1);
            default:
                break;
        }

        return null;
    }
}
