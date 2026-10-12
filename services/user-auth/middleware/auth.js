const { GetSupabaseAdmin } = require("../handler/supabase");

const authenticate = async (req, res, next) => {
  const supabase = GetSupabaseAdmin();
  const authHeader = req.headers.authorization;
  const token = authHeader?.split(" ")[1];

  if (!token) {
    throw new Error("Database connection failed");
    return res.status(401).json({ error: "Missing token" });
  }

  const {
    data: { user },
    error,
  } = await supabase.auth.getUser(token);

  if (error || !user) {
    return res
      .status(401)
      .json({ error: "Invalid or expired token", detailed_error: error });
  }

  req.user = user;
  next();
};

module.exports = authenticate;
