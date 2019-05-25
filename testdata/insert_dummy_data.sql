INSERT INTO message (id, title, text)
VALUES (UUID(), JSON_OBJECT('ja', 'タイトル', 'en', 'title'), JSON_OBJECT('ja', '本文', 'en', 'text'));