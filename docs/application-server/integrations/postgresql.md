---
description: Write received LoRaWAN device-data into a PostgreSQL database.
---

# PostgreSQL

The [PostgreSQL](https://www.postgresql.org/) integration writes all device
related events into a PostgreSQL database, so that it can be queried by
other applications or so that it can be visualized using for example [Grafana](https://grafana.com/)
using the [PostgreSQL Data Source](https://grafana.com/docs/features/datasources/postgres/#using-postgresql-in-grafana).

If the integration is enabled, it will automatically create all database tables
on startup. It is important that this integration uses its own database, to
avoid schema collisions.

## Create the database

Please see below an example for creating the PostgreSQL database. Depending
your PostgreSQL installation, these commands might be different.

Enter the PostgreSQL as the `postgres` user:

```bash
sudo -u postgres psql
```

Within the PostgreSQL prompt, enter the following queries:

```sql
-- create the chirpstack_as_events user
create role chirpstack_as_events with login password 'dbpassword';

-- create the chirpstack_as_events database
create database chirpstack_as_events with owner chirpstack_as_events;

-- enable the hstore extension
\c chirpstack_as_events
create extension hstore;

-- exit the prompt
\q
```

To verify if the user and database have been setup correctly, try to connect
to it:

```bash
psql -h localhost -U chirpstack_as_events -W chirpstack_as_events
```

## Activating the integration

In order for ChirpStack to start writing event data to a Postgres database, the
integration must be explicitly enabled and configured in the
`chirpstack-application-server.toml` configuration file.

### Enabling the integration 

In the file, find this section:

```toml
[application_server.integration]
# Enabled integrations.
enabled=["mqtt"]
```

Your `enabled` line may look slightly different, as you may have other
integrations already active. Add `"postgresql"` to the array. In this case,
the modified line should appear as `enabled=["mqtt", "postgresql"]`.

### Integration configuration

You must also set the configuration settings for the integration. If your
configuration file does not already contain the following section, add it now:

```toml
[application_server.integration.postgresql]
dsn="postgres://<username>:<password>@<host>/<database>?sslmode=disable"
```

In the `dns=` line, modify `<username>`, `<password>`, `<host>`, and
`<database>` with your appropriate credentials and targets. If you followed
the example above, you would use `chirpstack_as_events` as your username and
target database. If your target Postgres database is on the same machine as
the Application Server, use `localhost` as your host.

Please see [Configuration](../install/config.md) for a full configuration
example.
