const { createClient } = require("@supabase/supabase-js");
require("dotenv").config();
// Create a single supabase client for interacting with your database
const supabase = createClient(
  process.env.SUPABASE_URL,
  process.env.SUPABASE_PUBLISHABLE_KEY,
);
// Create a single supabase client for interacting with your database

const GetSupabaseAdmin = () => {
  return supabase;
};
// const GetSupabaseUser = (jwt) => {
//   return (supabase = createClient(
//     process.env.SUPABASE_URL,
//     process.env.SUPABASE_PUBLISHABLE_KEY,
//   ));
// };
module.exports = { GetSupabaseAdmin };
