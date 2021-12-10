CREATE TABLE IF NOT EXISTS books (
    Id SERIAL UNIQUE ,
    Title varchar NOT NULL ,
    Authors text[] NOT NULL ,
    Year date NOT NULL
);

CREATE INDEX IF NOT EXISTS title_idx ON books (title);
CREATE INDEX IF NOT EXISTS authors_idx ON books (authors);
CREATE INDEX IF NOT EXISTS year_idx ON books (year);