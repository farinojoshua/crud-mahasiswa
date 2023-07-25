
## Run Locally

Clone the project

```bash
  git clone https://github.com/farinojoshua/crud-mahasiswa.git
```

Go to the project directory

```bash
  cd crud-mahasiswa
```

update dependencies

```bash
  go mod tidy
```

Create Database in mysql

```bash
  CREATE DATABASE mahasiswa;
```

Setup configuration for db at db/db.go

Create Table in mysql

```bash
  CREATE TABLE mahasiswa (
    id int,
    nim varchar(255),
    nama varchar(255),
    umur int,
    prodi varchar(255),
    alamat varchar(255)
);
```

Start the server

```bash
  go run main.go
```
