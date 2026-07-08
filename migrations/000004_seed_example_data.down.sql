DELETE FROM match_players
WHERE match_id IN (1, 2);

DELETE FROM matches
WHERE id IN (1, 2);

DELETE FROM players
WHERE id IN (1, 2, 3);