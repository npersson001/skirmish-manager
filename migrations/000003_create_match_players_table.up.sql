CREATE TABLE match_players (
    match_id BIGINT NOT NULL,
    player_id BIGINT NOT NULL,

    PRIMARY KEY(match_id, player_id),

    INDEX idx_match_players_player_id(player_id),

    FOREIGN KEY (match_id)
       REFERENCES matches(id)
       ON DELETE CASCADE,

    FOREIGN KEY (player_id)
        REFERENCES players(id)
);