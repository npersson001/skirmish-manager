# skirmish-manager
Learning project on blockchain

# Requirements 
- mySql running server 
- podman (or docker desktop, but setup will show with podman)

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
```


# Decisions

Chi used for middleware because it is lightweight and close to net/http. Encourages good organization. 
