## Run Container

```
make build up
```

## Run LiteFS

console A
```
make console

cd /app
make run CONFIG=litefs.yml
```

console B
```
make console

cd /app
make run CONFIG=litefs2.yml
```

## Run SQLite3 in LiteFS

console C
```
sqlite3 /tmp/x/hoge

> CREATE TABLE hoge(id integer);
> INSERT INTO hoge VALUES (1);
```

console D
```
sqlite3 /tmp/x/hoge

> SELECT * FROM hoge;
```

