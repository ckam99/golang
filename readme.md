# Web application with golang


## commands

Help
```bash
./console help
```

Migrate database
```bash
./console db:migrate
```

Create superuser
```bash
./console createsuperadmin --name [yourname] --email [email address] --password [yourpassword]
```

Seeding database
```bash
./console db:seed --table [table name] 
or 
./console db:seed --all
```

Create fake database
```bash
./console db:fake --table [table name]  --count [number of entrees, default 1]
```

