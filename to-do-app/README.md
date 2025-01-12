# Running the Application with Docker

To run the application using Docker, follow these steps:

1. Ensure you have Docker and Docker Compose installed on your machine.

2. Clone the repository to your local machine.

3. Navigate to the directory containing the `docker-compose.yml` file.

4. Run the following command to start all services:

   ```sh
   docker-compose up --build
   ```

5. The following services will be available:
   - To-Do Application: `http://localhost:8080`
   - Grafana Dashboard: `http://localhost:3000`
   - Loki (Log Aggregator): `http://localhost:3100`
   - PostgreSQL Database: `localhost:5432`

6. First-time Grafana setup:
   - Log in to Grafana using the default credentials (admin/admin)
   - Add Loki as a data source:
     - Go to Configuration > Data Sources
     - Add new data source
     - Select Loki
     - Set the URL to `http://loki:3100`
     - Click "Save & Test"

7. To stop the services, run:

   ```sh
   docker-compose down
   ```

8. To remove all data (including database and Grafana data) and clean up, run:

   ```sh
   docker-compose down -v
   ```