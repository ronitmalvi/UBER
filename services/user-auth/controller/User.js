const { GetSupabaseAdmin } = require("../handler/supabase");

const getuser = async (req, res) => {
  const { data, error } = await GetSupabaseAdmin().from("Users").select("*");
  console.log(error, data);
  res.send(data);
};

module.exports = { getuser };
