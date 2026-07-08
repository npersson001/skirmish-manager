# skirmish-manager
Learning project on blockchain

# Requirements 
- Podman + podman-compose
- Go
- golang-migrate CLI

# Local Setup 
Install brew:
`/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`

Install podman: 
My machine has docker aliased to podman. I did not want to reset everything so this is just going to have to work with 
podman. 
```
podman machine init
podman machine start
```

Install podman compose: 
`brew install podman-compose`

Start mySql DB via docker compose: 
`podman-compose up -d`

Setup DB: 
```
podman exec -it skirmish-mysql mysql -u root -p
```

Switch to skirmish DB: 
```
USE skirmish;
```

Create table:
``` 
CREATE TABLE players (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    description TEXT
);

CREATE TABLE matches (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    winner_player_id BIGINT NOT NULL,
    started_at DATETIME NOT NULL,
    ended_at DATETIME NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (winner_player_id)
        REFERENCES players(id)
);

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
```

Setup migration: 
`brew install golang-migrate`

Run migration: 
```
migrate \
-path migrations \
-database "mysql://app:password@tcp(localhost:3306)/skirmish" \
up
```

# Decisions

## Technical 
- Chi used for middleware because it is lightweight and close to net/http. Encourages good organization. 
- sqlx used for DB access because it allows sql inline which gives us flexibility for more complex queries than ORMs allow

## Business Logic 
- Matches are not update-able, they are immutable
- Deleting a player will not delete a match/match_player as those are historical records 
- Deleting a match will delete a match_player though, but is not a common pattern we would want to support in reality

# To Implement
- Add Swagger
- Add tests
- Add config/env handling
- Add Dockerfile for API
- Figure out how blockchain works and add it + achievement table / api

# Improvements
There are some things I know I should do but do not have time for / are not high enough value to focus on right now: 
- defined errors throughout 
- structured logging 