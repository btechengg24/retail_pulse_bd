# **Retail Pulse Backend Development Assignment**


## **Description**
This project is a backend service for managing job submissions and their statuses in a retail environment. The system processes store visit data and provides APIs to submit and retrieve the status of jobs. The data is validated against a master dataset of store information, ensuring that only valid stores are processed.


## **Assumptions**
- **Store Master Data**: The `store_master.csv` file contains valid and complete information for all stores.
- **Unique Store IDs**: Each store in the `store_master.csv` has a unique `store_id`.
- **Job Validation**: If any `store_id` in a job request does not exist in the store master, it will return an error but still process the rest of the valid data.
- **Time Format**: All timestamps in the input follow ISO 8601 format (`YYYY-MM-DDTHH:MM:SSZ`).
- **Job IDs**: Generated job IDs are unique and timestamp-based.


## **Installing and Testing Instructions**

### **Clone the Repository**
```bash
git clone https://github.com/<your-repository-name>.git
cd RETAIL_PULSE_BD
```

### **Setup the Environment**
Ensure you have Go installed (version 1.23.3 or higher). You can download it from Go's official website.

Verify the Go installation:
```bash
go version
```

### **Install Dependencies**
Navigate to the project folder and run:
```bash
go mod tidy
```

### **Run the Server**
Start the server using:
```bash
go run main.go
```

### **Test the APIs**
Use tools like Postman or curl to test the following endpoints:

- **Submit Job**: `POST /api/submit`

  **Example Payload**:
  ```json
  {
    "count": 2,
    "visits": [
      {
        "store_id": "S00339218",
        "image_url": ["http://example.com/image1.jpg"],
        "visit_time": "2024-11-15T10:00:00Z"
      },
      {
        "store_id": "S01408764",
        "image_url": ["http://example.com/image2.jpg"],
        "visit_time": "2024-11-15T11:00:00Z"
      }
    ]
  }
    ```
- **Get Job Status**: `GET /api/status?jobid=<JOB_ID>`


## **Work Environment**
- **Computer**: Intel Core i5, 16GB RAM, Windows 10
- **Text Editor/IDE**: Visual Studio Code
- **Go Version**: 1.23.3

### **Libraries**
- Built-in Go packages like `net/http`, `encoding/json`, and `os`.

### **Data Source**
- `store_master.csv` for store data.


## **Improvements if Given More Time**
- **Enhance Error Handling**: Add more detailed error codes and messages for better API responses.
- **Database Integration**: Replace the `store_master.csv` file with a relational database (e.g., PostgreSQL) for improved scalability and performance.
- **Concurrency Improvements**: Optimize job processing with Goroutines to handle large payloads efficiently.
- **Testing Suite**: Add unit tests and integration tests for all APIs and core functionalities.
