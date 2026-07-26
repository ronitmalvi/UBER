const { createClient } = require("@supabase/supabase-js");
require("dotenv").config();
// Create a single supabase client for interacting with your database
const supabase = createClient(
  process.env.SUPABASE_URL,
  process.env.SUPABASE_PUBLISHABLE_KEY,
);

const logdata = async (req, res) => {
  const { data, error } = await supabase.auth.signInWithPassword({
    email: "durgesh3103@gmail.com",
    password: "durgesh03",
  });
  res.send(data);
};

module.exports = { logdata };
