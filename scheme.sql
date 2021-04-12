CREATE TABLE service
(
    request_id    bigint unsigned not null primary key auto_increment,
    user_id INTEGER       NOT NULL,
    date_  DATETIME    NOT NULL,
    state_request BOOLEAN NOT NULL,
    service_id INTEGER    NOT NULL,
    registry_request  DATETIME         NOT NULL
);