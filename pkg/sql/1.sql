CREATE TABLE spots (
    name varchar(255) NOT NULL,
    created_at bigint NOT NULL,
    PRIMARY KEY (name)
);

CREATE TABLE photos (
    id varchar(36) NOT NULL,
    src text NOT NULL,
    spot varchar(255) NOT NULL,
    created_at bigint NOT NULL,
    PRIMARY KEY (id)
);
