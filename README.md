## Building the Project

To build the project, you can use `make`. Run the following command:

```bash
make
```
## Running the Application

```bash
./build/app -db "root:examplepassword@tcp(localhost:3306)/exampledb" -ipfs "http://localhost:5001" -table users
```

### Command Parameters

The application supports the following command-line parameters:

- `-db`: Specifies the database connection details in the format `username:password@tcp(hostname:port)/databasename`.
  Example: `-db "root:examplepassword@tcp(localhost:3306)/exampledb"`

- `-ipfs`: Sets the IPFS URL for connecting to an IPFS node.
  Example: `-ipfs "http://localhost:5001"`

- `-table`: Specifies the table name to be accessed within the database.
  Example: `-table users`
