#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash, PartialOrd, Ord)]
pub enum Direction {
    North,
    South,
    East,
    West,
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
