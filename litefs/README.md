## Run Container

```
make build up
```

## Run LiteFS

console A
```
make console

make run CONFIG=litefs.yml
```

console B
```
make console

make run CONFIG=litefs2.yml
```

## Run SQLite3 in LiteFS

console C
```
make run/sqlite3 FILE=/tmp/x/hoge
sqlite3 

> CREATE TABLE hoge(id integer);
> INSERT INTO hoge VALUES (1);
```

console D
```
make run/sqlite3 FILE=/tmp/y/hoge

> SELECT * FROM hoge;
```

