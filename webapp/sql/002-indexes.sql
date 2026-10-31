-- reset.sql の先頭から読む。2回目以降はすべて既にあるので何もしない
CREATE UNIQUE INDEX IF NOT EXISTS users_username ON users(username);
CREATE INDEX IF NOT EXISTS users_region_status ON users(region, status);
CREATE INDEX IF NOT EXISTS posts_image_id ON posts(image_id);
CREATE INDEX IF NOT EXISTS realmojis_image_id ON realmojis(image_id);
CREATE INDEX IF NOT EXISTS friendships_user_high_id ON friendships(user_high_id);
CREATE INDEX IF NOT EXISTS follows_followee_id ON follows(followee_id);
CREATE INDEX IF NOT EXISTS friend_requests_requester_recipient ON friend_requests(requester_id, recipient_id);
CREATE INDEX IF NOT EXISTS friend_requests_recipient_id ON friend_requests(recipient_id);
CREATE INDEX IF NOT EXISTS login_sessions_user_id ON login_sessions(user_id);
