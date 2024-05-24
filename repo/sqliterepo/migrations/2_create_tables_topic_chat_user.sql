CREATE TABLE chat (
    id INTEGER PRIMARY KEY
);


CREATE TABLE user_topic (
    user_id INTEGER,
    chat_id INTEGER,
    topic_name TEXT,
    PRIMARY KEY (user_id, chat_id, topic_name)
)
