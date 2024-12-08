DROP TABLE IF EXISTS reply_relations;
CREATE TABLE reply_relations (
    post_id INT NOT NULL,
    parent_reply_id INT DEFAULT NULL,
    reply_id INT NOT NULL,
    relation_depth INT NOT NULL,
    PRIMARY KEY (post_id, reply_id, relation_depth),
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (parent_reply_id) REFERENCES replies(id) ON DELETE CASCADE,
    FOREIGN KEY (reply_id) REFERENCES replies(id) ON DELETE CASCADE
);
