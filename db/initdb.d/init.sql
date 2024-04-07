DROP SCHEMA IF EXISTS scsbot;
CREATE SCHEMA scsbot;
USE scsbot;

DROP TABLE IF EXISTS line_group;
CREATE TABLE line_group (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    group_name VARCHAR(255) NULL DEFAULT 'undefined',
    group_id VARCHAR(255) NOT NULL,
    invalid TINYINT UNSIGNED NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE(group_id),
    INDEX(invalid)
);

DROP TABLE IF EXISTS scheduling;
CREATE TABLE scheduling (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    line_group_id BIGINT UNSIGNED NOT NULL,
    cron VARCHAR(255) NOT NULL,
    snooze_interval_minutes INT UNSIGNED NOT NULL DEFAULT 15,
    snooze_limit_minutes INT UNSIGNED NOT NULL DEFAULT 60,
    invalid TINYINT UNSIGNED NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE(line_group_id),
    INDEX(invalid),
    FOREIGN KEY (line_group_id) REFERENCES line_group (id)
);

DROP TABLE IF EXISTS target_user;
CREATE TABLE target_user (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    line_group_id BIGINT UNSIGNED NOT NULL,
    target_user VARCHAR(255) NOT NULL,
    invalid TINYINT UNSIGNED NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX(line_group_id, target_user),
    INDEX(invalid),
    FOREIGN KEY (line_group_id) REFERENCES line_group (id)
);

DROP TABLE IF EXISTS notif_event;
CREATE TABLE notif_event (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    group_id VARCHAR(255) NOT NULL,
    target_user VARCHAR(255) NOT NULL,
    invalid TINYINT UNSIGNED NOT NULL DEFAULT 0,
    last_notified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    replyed_at TIMESTAMP NULL DEFAULT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX(group_id),
    INDEX(last_notified_at),
    INDEX(updated_at),
    INDEX(invalid)
);