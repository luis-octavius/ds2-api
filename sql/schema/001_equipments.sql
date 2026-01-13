-- +goose Up 
CREATE TABLE miracles (
  id UUID PRIMARY KEY NOT NULL, 
  name TEXT UNIQUE NOT NULL, 
  uses TEXT NOT NULL, 
  attunement INTEGER NOT NULL,
  faith INTEGER NOT NULL, 
  description TEXT NOT NULL, 
  acquired TEXT NOT NULL, 
  miracle_type TEXT NOT NULL
);

CREATE TABLE sorceries (
  id UUID PRIMARY KEY NOT NULL, 
  name TEXT UNIQUE NOT NULL, 
  uses TEXT NOT NULL, 
  spellslot INTEGER NOT NULL,
  intelligence INTEGER NOT NULL, 
  description TEXT NOT NULL, 
  acquired TEXT NOT NULL, 
  cost TEXT NOT NULL,
  sorcery_type TEXT NOT NULL
);

CREATE TABLE hexes (
  id UUID PRIMARY KEY NOT NULL, 
  name TEXT UNIQUE NOT NULL, 
  uses TEXT NOT NULL, 
  attunement INTEGER NOT NULL, 
  intelligence INTEGER NOT NULL, 
  faith INTEGER NOT NULL, 
  description TEXT NOT NULL, 
  acquired TEXT NOT NULL, 
  cost TEXT NOT NULL, 
  hex_type TEXT NOT NULL
);

CREATE TABLE pyromancies (
  id UUID PRIMARY KEY NOT NULL,
  name TEXT UNIQUE NOT NULL, 
  uses TEXT NOT NULL,
  attunement INTEGER NOT NULL, 
  description TEXT NOT NULL, 
  acquired TEXT NOT NULL, 
  cost TEXT NOT NULL,
  pyromancy_type TEXT NOT NULL
);

-- +goose Down 
DROP TABLE pyromancies; 
DROP TABLE hexes; 
DROP TABLE miracles; 
DROP TABLE sorceries;
