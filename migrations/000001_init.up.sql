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
'In the dark, we find the glow. A light that never fades, it grows.\nThrough the storm, we hold our ground. In silence, we hear the sound. Midnight calls, the stars align. Whispers of the endless time.\nHold my hand, don''t let it go. Together we will find the glow. And when the world is turning gray. We will find a brighter way today.'),
('Crystal Echo', 'Frozen Sky', 'Indie', 'Echoes in Time', 'https://example.com/frozensky', 
'The frozen sky above us. A thousand dreams below.\nWe walk through ice and shadows. Where only silence grows.\nBut in this frozen moment. I see the fire in your eyes.\nA spark that lights the endless. Across these empty skies.\nHold me close, don''t let it fade. Through frozen winds, we''ll find our way.'),
('Lunar Drift', 'Eclipse Heart', 'AltRock', 'Shadowed Moon', 'https://example.com/eclipseheart', 
'The moonlight fades, the shadows grow. In the night, we lose control.\nAn eclipse of the heart, so cold and far. We chase the light, but here we are.\nYou and I, lost in the dark. But I can feel the beating spark.\nA heart eclipsed but still it shines. Between the lines, between the signs.\nThe world might fall, but we stand tall. Through the eclipse, we''ll have it all.'),
('Echoes of Dawn', 'Silver Horizon', 'Ambient', 'Whispers of Light', 'https://example.com/silverhorizon', 
'On the silver horizon, we chase the sun. Through the endless waves, our journey''s begun.\nThe sky is wide, the sea so deep. In the echoes of dawn, our secrets we keep.\nSail with me, through wind and rain. Across the sky, through joy and pain.\nTogether we''ll find what we''re searching for. Beyond the silver, there''s always more.'),
('Starblaze', 'Galactic Dream', 'EDM', 'Into the Stars', 'https://example.com/galacticdream', 
'In the night, we fly so high. Through the stars, beyond the sky.\nA galactic dream, it feels so real. In this cosmic dance, we break the seal.\nLights and sounds, they spin around. In the vast unknown, we''re never found.\nTake my hand, let''s drift away. Into the night, into the day.\nGalactic dream, where we are free. Lost in the stars, just you and me.'),
('Velvet Surge', 'Crimson Tide', 'Rock', 'Waves of Fire', 'https://example.com/crimsontide', 
'The crimson tide is rising high. Waves of fire touch the sky.\nWe ride the storm, we fight the sea. The winds of fate will set us free.\nThrough the ashes and the flame. We will rise to claim our name.\nThe tide is strong, the night is long. But together we belong.\nIn the crimson waves, we''ll find our way. Through the storm, we''ll make them pay.'),
('Aurora Bloom', 'Echoes of You', 'Pop', 'Whispers in the Wind', 'https://example.com/echoesofyou', 
'I hear the echoes of you in the night. A voice that whispers, a fading light.\nThrough the stars, I search the sky. But in the dark, I still wonder why.\nWhy did you leave, where did you go? The answers lost in the undertow.\nBut in the echoes, I feel you near. A fleeting touch, a falling tear.\nEchoes of you, forever stay. Even though you''ve gone away.'),
('Darkstream', 'Phantom Lights', 'Gothic', 'Shadows and Flame', 'https://example.com/phantomlights', 
'In the night, the phantom lights. Dance with shadows, out of sight.\nThrough the mist, they call my name. A distant echo, a burning flame.\nLost in time, I walk alone. Among the ruins, among the stone.\nBut the lights, they guide me through. To a place where dreams come true.\nPhantom lights, they never die. Guiding me through endless sky.'),
('Solar Veil', 'Astral Pulse', 'Trance', 'Eternal Frequencies', 'https://example.com/astralpulse', 
'Feel the pulse of the astral wave. A rhythm that the stars once gave.\nThrough the night, we move as one. Our journey starts but never done.\nThe beat of space, the pulse of time. A cosmic dance, a sacred rhyme.\nLet the astral currents flow. To places only we can know.\nIn the pulse, we find the way. To brighter stars and endless day.'),
('Skylight Grove', 'Whispering Trees', 'Folk', 'Nature''s Call', 'https://example.com/whisperingtrees', 
'The trees are whispering your name. A voice so soft, yet not the same.\nIn the woods, we walk so slow. Through the paths, where secrets grow.\nThe leaves, they fall, like autumn rain. But the whispers still remain.\nHold me close, don''t let them fade. In the forest, memories are made.\nWhispering trees, they tell the tale. Of love that never could fail.');
