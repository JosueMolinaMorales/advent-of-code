use std::cmp::Ordering;

#[derive(Debug, Clone, PartialEq, Eq, Hash)]
pub struct Point {
    pub x: i32,
    pub y: i32,
}

impl Point {
    pub fn new(x: i32, y: i32) -> Self {
        Point { x, y }
    }

    pub fn distance(&self, other: &Point) -> i32 {
        (self.x - other.x).abs() + (self.y - other.y).abs()
    }
}

impl PartialOrd for Point {
    fn partial_cmp(&self, other: &Point) -> Option<Ordering> {
        Some(self.cmp(&other))
    }
}

impl Ord for Point {
    fn cmp(&self, other: &Point) -> Ordering {
        if self.x < other.x {
            Ordering::Less
        } else if self.x > other.x {
            Ordering::Greater
        } else if self.y < other.y {
            Ordering::Less
        } else if self.y > other.y {
            Ordering::Greater
        } else {
            Ordering::Equal
        }
    }
}
