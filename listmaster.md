# ListMaster

## Идея

это будет бот для [telegram](https://telegram.org/), который будет позволять вести некие древовидные (пока 2 уровня) списки. 
Самый просто способ использования: список покупок. Особенно, если включить этого бота в семейную группу и тогда все могут вносить покупки в список и все могут это список видеть, и работать с ним.

## ToDo
- [x] бот @done()
- [x] анализ команд @done 
- [ ] /add
    - [x] добавление в памяти @done 
    - [x] запись в базу (перечитывать дерево после)  @done(2016-04-23)
    - [ ] округления, чтобы потом точнее искать @na
- [ ] /list
    - [x] показ всего дерева @done
    - [ ] показ только одного списка
- [ ] /done
- [ ] /del 
- [x] хранение списков  @done(2016-04-23)
- [ ] Позже
    - [ ] парсинг слачая, когда "/add 1 что-нибудь к чаю", т.е. товар из нескольких слов, но без кавычек???
    - [ ] не вываливаться при ошибках базы
    - [ ] sql-inject ???
    - [ ] оптимизация AddElement, там можно убрать добавление элемента к дереву в памяти.
  
### хранение данных
отладочное дерево

    DevData = []ListElement{
        {1,"Аптека"},
        {1.001,"Канефрон"},
        {1.002,"Йод"},
        {2,"Зоо магазин"},
        {2.001,"Феликс 10 пакетиков"},
        {3,"Овощи, фрукты"},
        {3.001,"огурцы"},
    }


    

## SQL

    host=localhost
    port=5432
    user=USER
    password=PASSWORD
    dbname=DBNAME

    DROP SEQUENCE s_tt CASCADE;
    DROP TABLE tt CASCADE ;

    CREATE SEQUENCE s_tt;
    CREATE TABLE  tt(
        id integer DEFAULT nextval('s_tt'::regclass) NOT NULL UNIQUE,
        chat_id int NOT NULL,
        idx float,
        item text,
        show bool default 't'
    );

    begin;
    insert into tt (chat_id, idx, item) values (120864, 1, 'Аптека');
    insert into tt (chat_id, idx, item) values (120864, 1.001, 'Канефрон');
    insert into tt (chat_id, idx, item) values (120864, 1.002, 'Йод');
    insert into tt (chat_id, idx, item) values (120864, 2, 'Зоо магазин');
    insert into tt (chat_id, idx, item) values (120864, 2.001, 'Феликс 10 пакетиков');
    insert into tt (chat_id, idx, item) values (120864, 3, 'Овощи, фрукты');
    insert into tt (chat_id, idx, item) values (120864, 3.001, 'огурцы');
    insert into tt (chat_id, idx, item) values (-43927056, 1, 'Drug store');
    insert into tt (chat_id, idx, item) values (-43927056, 1.001, 'Canefrom');
    insert into tt (chat_id, idx, item) values (-43927056, 1.002, 'Yod');
    insert into tt (chat_id, idx, item) values (-43927056, 2, 'Zoo shop');
    insert into tt (chat_id, idx, item) values (-43927056, 2.001, 'Felix 10 packets');
    insert into tt (chat_id, idx, item) values (-43927056, 3, 'Green');
    insert into tt (chat_id, idx, item) values (-43927056, 3.001, 'potatoes');
    commit;



