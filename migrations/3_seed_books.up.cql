BEGIN
    BATCH
    INSERT INTO library.books (id, title, author)
    VALUES (uuid(), 'The Go Programming Language', 'Alan Donovan');

    INSERT INTO library.books (id, title, author)
    VALUES (uuid(), 'Clean Code', 'Robert Martin');

    INSERT INTO library.books (id, title, author)
    VALUES (uuid(), 'The Pragmatic Programmer', 'Andy Hunt');
    APPLY BATCH;