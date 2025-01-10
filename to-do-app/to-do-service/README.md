# Running the Application with Docker

To run the application using Docker, follow these steps:

1. Ensure you have Docker and Docker Compose installed on your machine.

2. Clone the repository to your local machine.

3. Navigate to the directory containing the `docker-compose.yml` file.

4. Run the following command to start the services:

   ```sh
   docker-compose up --build
   ```

5. The application will be available at `http://localhost:8080`.

6. To stop the services, run:

   ```sh
   docker-compose down
   ```

7. To remove the volumes and clean up, run:

   ```sh
   docker-compose down -v
   ```
