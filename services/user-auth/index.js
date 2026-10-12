const express = require("express");
const { createClient } = require("@supabase/supabase-js");
const { logdata } = require("./supaclient");
const notFound = require("./middleware/notFound");
const errorHandler = require("./middleware/errorhandler");

const UserRouter = require("./routes/UserRouter");

require("dotenv").config();
const app = express();

app.use(express.json());

app.get("/loginUser", async (req, res) => {
  logdata(req, res);
});

app.use("/api/v1/user_auth", UserRouter);
app.use(errorHandler);
app.use(notFound);

app.listen(process.env.PORT, () => {
  console.log(`Server is running at http://localhost:${process.env.PORT}`);
});
