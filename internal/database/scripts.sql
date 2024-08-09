-- Purpose: This file contains the SQL script to create the database schema for the application.
-- Also seeds the database with some initial data.

-- Drop tables if they exist
DROP TABLE IF EXISTS JOURNAL_ENTRIES;
DROP TABLE IF EXISTS JOURNALS;
DROP TABLE IF EXISTS SYMPTOMS;
DROP TABLE IF EXISTS EXERCISES;
DROP TABLE IF EXISTS USERS;

-- Create USERS table
CREATE TABLE USERS (
    USER_ID SERIAL PRIMARY KEY,
    FIRST_NAME VARCHAR(100) NOT NULL,
    LAST_NAME VARCHAR(100) NOT NULL,
    EMAIL VARCHAR(1024) NOT NULL,
    OAUTH_ID VARCHAR(1024),
    AVATAR_URL VARCHAR(1024),
    "LOCATION" VARCHAR(1024)
);

-- Create EXERCISES table
CREATE TABLE EXERCISES (
    ID SERIAL PRIMARY KEY,
    EXERCISE_NAME VARCHAR(100) NOT NULL
);

-- Create SYMPTOMS table
CREATE TABLE SYMPTOMS (
    SYMPTOM_ID SERIAL PRIMARY KEY,
    SYMPTOM_NAME VARCHAR(100) NOT NULL
);

-- Create JOURNALS table
CREATE TABLE JOURNALS (
    JOURNALS_ID SERIAL PRIMARY KEY,
    USER_ID INT NOT NULL REFERENCES USERS(USER_ID) ON DELETE CASCADE,
    JOURNAL_NAME VARCHAR(100) NOT NULL
);

-- Create JOURNAL_ENTRIES table
CREATE TABLE JOURNAL_ENTRIES (
    JOURNAL_ENTRY_ID SERIAL PRIMARY KEY,
    USER_ID INT NOT NULL REFERENCES USERS(USER_ID) ON DELETE CASCADE,
    SYMPTOM_ID INT NOT NULL REFERENCES SYMPTOMS(SYMPTOM_ID) ON DELETE CASCADE,
    EXERCISE_ID INT NOT NULL REFERENCES EXERCISES(ID) ON DELETE CASCADE,
    JOURNALS_ID INT NOT NULL REFERENCES JOURNALS(JOURNALS_ID) ON DELETE CASCADE,
    JOURNAL_ENTRY_DATE DATE NOT NULL
);

-- Seed data
-- Get the USER_ID of the newly inserted user
WITH inserted_user AS (
    INSERT INTO USERS 
    (
    FIRST_NAME,
    LAST_NAME,
    EMAIL,
    OAUTH_ID,
    AVATAR_URL,
    "LOCATION")
    VALUES (
    'michael', 
    'duren',
    'michael@michael.com',
    'abc123',
    'https://www.example.com/avatar.jpg',
    'New York, NY'
    )
    RETURNING USER_ID
)
-- Seed JOURNALS with the new user's USER_ID
INSERT INTO JOURNALS (USER_ID, JOURNAL_NAME)
SELECT USER_ID, 'My Journal' FROM inserted_user;

-- Insert exercises
INSERT INTO EXERCISES (EXERCISE_NAME)
VALUES
    ('Running'),
    ('Walking'),
    ('Swimming'),
    ('Cycling'),
    ('Yoga'),
    ('Pilates'),
    ('Weight Lifting'),
    ('Dancing'),
    ('Boxing'),
    ('Martial Arts');

-- Insert symptoms
INSERT INTO SYMPTOMS (SYMPTOM_NAME)
VALUES
    ('Headache'),
    ('Nausea'),
    ('Fatigue'),
    ('Dizziness'),
    ('Shortness of Breath'),
    ('Chest Pain'),
    ('Back Pain'),
    ('Joint Pain'),
    ('Muscle Pain'),
    ('Stomach Pain');

-- Insert journal entries
WITH user_and_journal AS (
    SELECT u.USER_ID, j.JOURNALS_ID
    FROM USERS u
    JOIN JOURNALS j ON u.USER_ID = j.USER_ID
    WHERE u.FIRST_NAME = 'michael'
)
INSERT INTO JOURNAL_ENTRIES (
    USER_ID,
    SYMPTOM_ID,
    EXERCISE_ID,
    JOURNALS_ID,
    JOURNAL_ENTRY_DATE
)
VALUES
    ((SELECT USER_ID FROM user_and_journal), 1, 1, (SELECT JOURNALS_ID FROM user_and_journal), '2021-01-01'),
    ((SELECT USER_ID FROM user_and_journal), 2, 2, (SELECT JOURNALS_ID FROM user_and_journal), '2021-01-02'),
    ((SELECT USER_ID FROM user_and_journal), 3, 3, (SELECT JOURNALS_ID FROM user_and_journal), '2021-01-03'),
    ((SELECT USER_ID FROM user_and_journal), 4, 4, (SELECT JOURNALS_ID FROM user_and_journal), '2021-01-04'),
    ((SELECT USER_ID FROM user_and_journal), 5, 5, (SELECT JOURNALS_ID FROM user_and_journal), '2021-01-05');

