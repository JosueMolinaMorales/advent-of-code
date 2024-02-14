use super::point::Point;

#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash, PartialOrd, Ord)]
pub enum Direction {
    North,
    South,
    East,
    West,
}

impl Into<Point> for Direction {
    fn into(self) -> Point {
        match self {
            Direction::East => Point { x: 0, y: 1 },
            Direction::North => Point { x: -1, y: 0 },
            Direction::South => Point { x: 1, y: 0 },
            Direction::West => Point { x: 0, y: -1 },
        }
    }
}

impl From<String> for Direction {
    fn from(value: String) -> Self {
        match value.as_str() {
            ">" => Direction::East,
            "<" => Direction::West,
            "^" => Direction::North,
            "v" => Direction::South,
            _ => panic!("Not a valid direction string"),
        }
    }
}
