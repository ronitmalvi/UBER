const express = require("express");
const routes = express.Router();
const { getuser } = require("../controller/User");
const authenticate = require("../middleware/auth");

routes.route("/user").get(authenticate, getuser);
module.exports = routes;
