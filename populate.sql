
INSERT INTO snippets (title, content, created, expires) VALUES (
                                                                   'An old silent pond',
                                                                   'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.'
                                                                   UTC_TIMESTAMP(),
                                                                   DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
                                                                   );
INSERT INTO snippets (title, content, created, expires) VALUES (
                                            'Over the wintry forest',
                                            'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow'
                                            UTC_TIMESTAMP(),
                                            DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY))

INSERT INTO snippets (title, content, created, expires) VALUES (
'First autumn morning',
'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\'
UTC_TIMESTAMP(),
DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY)
);


use snippetbox; 
create table users (
    id integer not null primary key auto_increment, 
    name varchar(255) not null, 
    email varchar(255) not null, 
    hashed_password char(60) not null, 
    created datetime not null
); 

alter table users add constraint users_uc_email unique(email);

