# Always restore before starting
litestream restore todox.db

# Migrate the database
/bin/tools migrate

# Start sever and backup
litestream replicate -exec "/bin/app"