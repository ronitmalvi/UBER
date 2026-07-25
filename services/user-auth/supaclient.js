import { createClient } from "@supabase/supabase-js";

// Create a single supabase client for interacting with your database
const supabase = createClient(
  process.env.SUPABASE_URL,
  process.env.SUPABASE_PUBLISHABLE_KEY,
);

const Logdata = () => {
  const { data, error } = await supabase.from("UB_COM").select();
  console.log(data);
};

module.exports = {Logdata}
