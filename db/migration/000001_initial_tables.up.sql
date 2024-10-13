CREATE TABLE users
(
    id         SERIAL PRIMARY KEY NOT NULL UNIQUE,
    username   VARCHAR(255)       NOT NULL,
    password   VARCHAR(255)       NOT NULL,
    email      VARCHAR(255)       NOT NULL UNIQUE,
    created_at TIMESTAMP          NOT NULL DEFAULT now(),
    updated_at TIMESTAMP          NOT NULL DEFAULT now()
);

CREATE TABLE categories
(
    id          SERIAL PRIMARY KEY NOT NULL UNIQUE,
    user_id     SERIAL             NOT NULL,
    title       VARCHAR(255)       NOT NULL,
    type        VARCHAR(255)       NOT NULL,
    description VARCHAR(255)       NOT NULL,
    created_at  TIMESTAMP          NOT NULL DEFAULT now(),
    updated_at  TIMESTAMP          NOT NULL DEFAULT now()

);

ALTER TABLE categories
    ADD FOREIGN KEY (user_id) REFERENCES users (id);


CREATE TABLE bills
(
    id          SERIAL PRIMARY KEY NOT NULL UNIQUE,
    user_id     SERIAL             NOT NULL,
    category_id SERIAL             NOT NULL,
    title       VARCHAR(255)       NOT NULL,
    type        VARCHAR(255)       NOT NULL,
    description VARCHAR(255)       NOT NULL,
    amount      DECIMAL(10, 2)     NOT NULL,
    date        DATE               NOT NULL,
    created_at  TIMESTAMP          NOT NULL DEFAULT now(),
    updated_at  TIMESTAMP          NOT NULL DEFAULT now()
);

ALTER TABLE bills
    ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE bills
    ADD FOREIGN KEY (category_id) REFERENCES categories (id);
