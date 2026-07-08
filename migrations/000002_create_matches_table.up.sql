CREATE TABLE matches (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    winner_player_id BIGINT NOT NULL,
    started_at DATETIME NOT NULL,
    ended_at DATETIME NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (winner_player_id)
        REFERENCES players(id)
);