# Test task - API for managing library of songs

<h3>Export env variables</h3>

```
export DB="postgres:22222@localhost:5435/postgres?sslmode=disable" 
```

<h3>How to run</h3>

```
sudo docker run --name=api -e POSTGRES_PASSWORD=22222 -p 5435:5432 -d postgres:16.2
```

``` make run ```

<h3>Run migrations</h3>

``` make mig-up ```

``` make mig-down ```