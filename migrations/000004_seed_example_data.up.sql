INSERT INTO players (id, username, description)
VALUES
    (1, 'erik', 'Skirmish player'),
    (2, 'alex', 'Skirmish player'),
    (3, 'emma', 'Skirmish player');


INSERT INTO matches (
    id,
    winner_player_id,
    started_at,
    ended_at
)
VALUES
    (
        1,
        1,
        '2026-07-01 18:00:00',
        '2026-07-01 19:30:00'
    ),
    (
        2,
        1,
        '2026-07-02 18:00:00',
        '2026-07-02 20:00:00'
    );


INSERT INTO match_players (match_id, player_id)
VALUES
    (1, 1), -- Erik played match 1
    (1, 2), -- Alex played match 1
    (2, 1), -- Erik played match 2
    (2, 3); -- Emma played match 2