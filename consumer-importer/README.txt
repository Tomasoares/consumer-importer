CONSUMER-IMPORTER
-------------------

This small service has the goal to import data from a textfile to a postgres database. 
When it's executed, the service keeps checking the file repository folder for each 5 seconds
for a file to load. the default folder is configured as 'consumer-importer/file-repository',
but the folder itself can be changed in the consumer-importer.yml docker-compose file.

When a file is found, the importing algorithm is started, where the service will read the
file line-per-line and convert it into a Consumer database row with its data formatted and
validated. After the load is finished, the service keeps then checking for others files in the file 
repository folder.

-------------------
HOW TO RUN:
-------------------

To run the service, it's necessary to have Docker installed in the computer. Then, in the terminal,
execute 'docker-compose up' in the directory './consumer-importer/docker-compose'. The Docker application
will download the required images (golang and postgres) and dependencies and automatically run both the
database and the service. It might be necessary to run 'docker-compose up' twice. You can also configure
database connection string in both 'postgres.yml' and 'consumer-importer.yml' docker-compose files.

-------------------
DATABASE STRUCTURE:
-------------------

The database structure consists of two tables, CONSUMER and FILE. File contains all the data related to
the imported file, and it certifies that the same file doesn't get imported twice. Consumer contains all
the data related to the consumer in the imported file, with 3 additional columns:

- id_file: id of the source file
- valid: if the register is valid and doesn't have and CPF or CNPJ validation error.
- validation_message: The validation error message, in case the registry isn't valid.

-------------------
PROJECT STRUCTURE:
-------------------

The project is divided in 7 packages:

- docker-compose: contains all the docker-compose files to start the service and the postgres database.

- file-repository: it's the default folder where the service will check for files to import.

- io: contains all the .go files with functions and structs related to file reading and converting operations.

- main: contains the main function which is where the app is started and the consumer-importer algorithm which is
the main flow of the app.

- model: contains the CONSUMER and FILE structs, which communicates with the services structs to save or load data
from the database.

- service: contains the services structures for CONSUMER and FILE, which are responsible to do CRUD operations at
the database.

- validator: contains the validators functions necessary to validate the Consumer CPF and CNPJ data.

-------------------
FUTURE IMPLEMENTATIONS:
-------------------

Due to time constraints, I wasn't able to implement some features I'd like to do:

- Add a better dependency management system.
- Implement some more robust testing layer, using Mocking and increase couverture.
- Use go routines to speed up the importing process.