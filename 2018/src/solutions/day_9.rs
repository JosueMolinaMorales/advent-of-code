use std::collections::VecDeque;

pub fn solve_day_nine() {
    let players = 424;
    let last_marble = 71144;

    println!("Day 9 part 1: {}", solve(players, last_marble));
    println!("Day 9 part 2: {}", solve(players, last_marble * 100));
}

fn solve(players: i64, last_marble: i64) -> i64 {
    let mut player_scores = vec![0; players as usize];
    let mut curr_player = 0;
    let mut circle = VecDeque::from([0]);

    (1..=last_marble).for_each(|marble| {
        // Check if marble is divisible by 23
        if marble % 23 == 0 {
            // marble is not added to the circle, player keeps that marble
            player_scores[curr_player] += marble;
            // 7 marbles counter-clockwise from the current marble is removed from the circle and also added to the current player's score.
            circle.rotate_right(7);
            player_scores[curr_player] += circle.pop_back().unwrap();
            circle.rotate_left(1);
        } else {
            circle.rotate_left(1);
            circle.push_back(marble);
        }
        curr_player = (curr_player + 1) % players as usize;
    });

    player_scores.iter().max().unwrap().clone()
}
