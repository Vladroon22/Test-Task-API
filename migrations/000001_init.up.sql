CREATE TABLE library (
    id SERIAL PRIMARY KEY,
    name_group VARCHAR(30) NOT NULL,
    name_song VARCHAR(25) NOT NULL,
    genre VARCHAR(15) NOT NULL,
    album VARCHAR(20) NOT NULL,
    link VARCHAR(80) NOT NULL,
    lyrics VARCHAR(700) NOT NULL,
    createAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

INSERT INTO library (name_group, name_song, genre, album, link, lyrics) 
VALUES 
('Neon Waves', 'Midnight Glow', 'Synthpop', 'Electric Dreams', 'https://example.com/midnightglow', 
'In the dark, we find the glow\nA light that never fades, it grows\nThrough the storm, we hold our ground\nIn silence, we hear the sound\n\nMidnight calls, the stars align\nWhispers of the endless time\nHold my hand, don''t let it go\nTogether we will find the glow\n\nAnd when the world is turning gray\nWe will find a brighter way today.'),
('Crystal Echo', 'Frozen Sky', 'Indie', 'Echoes in Time', 'https://example.com/frozensky', 
'The frozen sky above us\nA thousand dreams below\nWe walk through ice and shadows\nWhere only silence grows\n\nBut in this frozen moment\nI see the fire in your eyes\nA spark that lights the endless\nAcross these empty skies\n\nHold me close, don''t let it fade\nThrough frozen winds, we''ll find our way.'),
('Lunar Drift', 'Eclipse Heart', 'AltRock', 'Shadowed Moon', 'https://example.com/eclipseheart', 
'The moonlight fades, the shadows grow\nIn the night, we lose control\nAn eclipse of the heart, so cold and far\nWe chase the light, but here we are\n\nYou and I, lost in the dark\nBut I can feel the beating spark\nA heart eclipsed but still it shines\nBetween the lines, between the signs\n\nThe world might fall, but we stand tall\nThrough the eclipse, we''ll have it all.'),
('Echoes of Dawn', 'Silver Horizon', 'Ambient', 'Whispers of Light', 'https://example.com/silverhorizon', 
'On the silver horizon, we chase the sun\nThrough the endless waves, our journey''s begun\nThe sky is wide, the sea so deep\nIn the echoes of dawn, our secrets we keep\n\nSail with me, through wind and rain\nAcross the sky, through joy and pain\nTogether we''ll find what we''re searching for\nBeyond the silver, there''s always more.'),
('Starblaze', 'Galactic Dream', 'EDM', 'Into the Stars', 'https://example.com/galacticdream', 
'In the night, we fly so high\nThrough the stars, beyond the sky\nA galactic dream, it feels so real\nIn this cosmic dance, we break the seal\n\nLights and sounds, they spin around\nIn the vast unknown, we''re never found\nTake my hand, let''s drift away\nInto the night, into the day\n\nGalactic dream, where we are free\nLost in the stars, just you and me.'),
('Velvet Surge', 'Crimson Tide', 'Rock', 'Waves of Fire', 'https://example.com/crimsontide', 
'The crimson tide is rising high\nWaves of fire touch the sky\nWe ride the storm, we fight the sea\nThe winds of fate will set us free\n\nThrough the ashes and the flame\nWe will rise to claim our name\nThe tide is strong, the night is long\nBut together we belong\n\nIn the crimson waves, we''ll find our way\nThrough the storm, we''ll make them pay.'),
('Aurora Bloom', 'Echoes of You', 'Pop', 'Whispers in the Wind', 'https://example.com/echoesofyou', 
'I hear the echoes of you in the night\nA voice that whispers, a fading light\nThrough the stars, I search the sky\nBut in the dark, I still wonder why\n\nWhy did you leave, where did you go?\nThe answers lost in the undertow\nBut in the echoes, I feel you near\nA fleeting touch, a falling tear\n\nEchoes of you, forever stay\nEven though you''ve gone away.'),
('Darkstream', 'Phantom Lights', 'Gothic', 'Shadows and Flame', 'https://example.com/phantomlights', 
'In the night, the phantom lights\nDance with shadows, out of sight\nThrough the mist, they call my name\nA distant echo, a burning flame\n\nLost in time, I walk alone\nAmong the ruins, among the stone\nBut the lights, they guide me through\nTo a place where dreams come true\n\nPhantom lights, they never die\nGuiding me through endless sky.'),
('Solar Veil', 'Astral Pulse', 'Trance', 'Eternal Frequencies', 'https://example.com/astralpulse', 
'Feel the pulse of the astral wave\nA rhythm that the stars once gave\nThrough the night, we move as one\nOur journey starts but never done\n\nThe beat of space, the pulse of time\nA cosmic dance, a sacred rhyme\nLet the astral currents flow\nTo places only we can know\n\nIn the pulse, we find the way\nTo brighter stars and endless day.'),
('Skylight Grove', 'Whispering Trees', 'Folk', 'Nature''s Call', 'https://example.com/whisperingtrees', 
'The trees are whispering your name\nA voice so soft, yet not the same\nIn the woods, we walk so slow\nThrough the paths, where secrets grow\n\nThe leaves, they fall, like autumn rain\nBut the whispers still remain\nHold me close, don''t let them fade\nIn the forest, memories are made\n\nWhispering trees, they tell the tale\nOf love that never could fail.');
