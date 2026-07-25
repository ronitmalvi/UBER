const express = require("express");

// Initialize the Express application
const app = express();

// Define the port number
const PORT = 3000;

// Handle HTTP GET requests to the root URL ('/')
app.get("/", (req, res) => {
  res.send("user-auth service");
});

// Start the server and listen on the specified port
app.listen(PORT, () => {
  console.log(`Server is running at http://localhost:${PORT}`);
});
