#!/bin/sh

# Run database migrations
npx prisma migrate deploy

# Run the main container command
exec "$@"