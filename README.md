# CC-EOY
Crowd Management for UOM EOY Party. This app allows event organisors
to scan Student IDs and mark students as present. The scan will
only be successful if a student registered for the event. In
case they did not, the application will allow the organisor
to manually register the student.

# Server
The server, on the `master` branch, is built in Go. To run, adjust
environment variables and build using:

```bash
go build cmd/main.go
```

This will output `main`, the executable. Afterwards, simply run it. The server
will automatically log events in `logs.txt` in a production environment. In
development, it will print the logs to standard output.

**Note**: The sqlite3 package this repo uses depends on CGO. As such
the above build command will not provide a purely static executable.
It will depend on the sqlite3 library at runtime. On debian systems,
install it as follows:

```bash
sudo apt install -y build-essential libsqlite3-dev
```

# Client
The frontend is a mobile app implemented in flutter. It is located on the
`frontend` branch. To build:

```bash
flutter build apk --release --dart-define-from-file .env
```

Make sure to pass the `--dart-define-from-file .env` so that environment variables
are properly initialised.
