---
description: Write received LoRaWAN device-data into a PostgreSQL database.
---

# PostgreSQL

The [PostgreSQL](https://www.postgresql.org/) integration writes all events
into a PostgreSQL database. This database can then be used by other
applications or visualized using for example [Grafana](https://grafana.com/)
using the [PostgreSQL Data Source](https://grafana.com/docs/features/datasources/postgres/#using-postgresql-in-grafana).

* ChirpStack Application Server will not create these tables for you. Create statements are
  given below.
* You must enable the `hstore` extension for this database table, this can be
  done with the SQL statement: `create extension hstore;`.
* This database does not have to be the same database as used by
  ChirpStack Application Server.
* This may generate a lot of data depending the number of devices and number
  of messages sent per device.

## Create database example

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

## Events

### Uplink data

Uplink data is written into the table `device_up`. The following schema
must exist:

```sql
create table device_up (
	id uuid primary key,
	received_at timestamp with time zone not null,
	dev_eui bytea not null,
	device_name varchar(100) not null,
	application_id bigint not null,
	application_name varchar(100) not null,
	frequency bigint not null,
	dr smallint not null,
	adr boolean not null,
	f_cnt bigint not null,
	f_port smallint not null,
	tags hstore not null,
	data bytea not null,
	rx_info jsonb not null,
	object jsonb not null
);

-- NOTE: These are example indices, depending on how this table is being
-- used, you might want to change these.
create index idx_device_up_received_at on device_up(received_at);
create index idx_device_up_dev_eui on device_up(dev_eui);
create index idx_device_up_application_id on device_up(application_id);
create index idx_device_up_frequency on device_up(frequency);
create index idx_device_up_dr on device_up(dr);
create index idx_device_up_tags on device_up(tags);
```

### Device status

Device-status data is written into the table `device_status`. The following
schema must exist:

```sql
create table device_status (
	id uuid primary key,
	received_at timestamp with time zone not null,
	dev_eui bytea not null,
	device_name varchar(100) not null,
	application_id bigint not null,
	application_name varchar(100) not null,
	margin smallint not null,
	external_power_source boolean not null,
	battery_level_unavailable boolean not null,
	battery_level numeric(5, 2) not null,
	tags hstore not null
);

-- NOTE: These are example indices, depending on how this table is being
-- used, you might want to change these.
create index idx_device_status_received_at on device_status(received_at);
create index idx_device_status_dev_eui on device_status(dev_eui);
create index idx_device_status_application_id on device_status(application_id);
create index idx_device_status_tags on device_status(tags);
```

### Join

Join notifications are written into the table `device_join`. The following
schema must exist:

```sql
create table device_join (
	id uuid primary key,
	received_at timestamp with time zone not null,
	dev_eui bytea not null,
	device_name varchar(100) not null,
	application_id bigint not null,
	application_name varchar(100) not null,
	dev_addr bytea not null,
	tags hstore not null
);

-- NOTE: These are example indices, depending on how this table is being
-- used, you might want to change these.
create index idx_device_join_received_at on device_join(received_at);
create index idx_device_join_dev_eui on device_join(dev_eui);
create index idx_device_join_application_id on device_join(application_id);
create index idx_device_join_tags on device_join(tags);
```

### ACK

ACK notifications are written into the table `device_ack`. The following schema
must exist:

```sql
create table device_ack (
	id uuid primary key,
	received_at timestamp with time zone not null,
	dev_eui bytea not null,
	device_name varchar(100) not null,
	application_id bigint not null,
	application_name varchar(100) not null,
	acknowledged boolean not null,
	f_cnt bigint not null,
	tags hstore not null
);

-- NOTE: These are example indices, depending on how this table is being
-- used, you might want to change these.
create index idx_device_ack_received_at on device_ack(received_at);
create index idx_device_ack_dev_eui on device_ack(dev_eui);
create index idx_device_ack_application_id on device_ack(application_id);
create index idx_device_ack_tags on device_ack(tags);
```

### Error

Error notifications are written into the table `device_error`. The following
schema must exist:

```sql
create table device_error (
	id uuid primary key,
	received_at timestamp with time zone not null,
	dev_eui bytea not null,
	device_name varchar(100) not null,
	application_id bigint not null,
	application_name varchar(100) not null,
	type varchar(100) not null,
	error text not null,
	f_cnt bigint not null,
	tags hstore not null
);

-- NOTE: These are example indices, depending on how this table is being
-- used, you might want to change these.
create index idx_device_error_received_at on device_error(received_at);
create index idx_device_error_dev_eui on device_error(dev_eui);
create index idx_device_error_application_id on device_error(application_id);
create index idx_device_error_tags on device_error(tags);
```

### Location

Location notifications are written into the table `device_location`. The
following schema must exist:

```sql
create table device_location (
	id uuid primary key,
	received_at timestamp with time zone not null,
	dev_eui bytea not null,
	device_name varchar(100) not null,
	application_id bigint not null,
	application_name varchar(100) not null,
	altitude double precision not null,
	latitude double precision not null,
	longitude double precision not null,
	geohash varchar(12) not null,
	tags hstore not null,

	-- this field is currently not populated
	accuracy smallint not null
);

-- NOTE: These are example indices, depending on how this table is being
-- used, you might want to change these.
create index idx_device_location_received_at on device_location(received_at);
create index idx_device_location_dev_eui on device_location(dev_eui);
create index idx_device_location_application_id on device_location(application_id);
create index idx_device_location_tags on device_location(tags);
```

## Activating the Integration

In order for ChirpStack to start writing event data to a Postgres database, the integration must be explicitly enabled and configured in the `chirpstack-application-server.toml` configuration file.

### Enabling the Integration 

Run `sudo nano /etc/chirpstack-application-server/chirpstack-application-server.toml` to open it in an editor.

In the file, find this section:
```toml
[application_server.integration]
  # Enabled integrations.
  enabled=["mqtt"]
```
Your `enabled` line may look slightly different, as you may have other integrations already active. Add `"postgresql"` to the array. In this case, the modified line should appear as `enabled=["mqtt", "postgresql"]`.

### Configuring the Integration

You must also set the configuration settings for the integration. If your configuration file does not already contain the following section, add it now:
```toml
  # PostgreSQL database integration.
  [application_server.integration.postgresql]
  # PostgreSQL dsn (e.g.: postgres://user:password@hostname/database?sslmode=disable).
  dsn="postgres://<username>:<password>@<host>/<database>?sslmode=disable"
  
  # This sets the max. number of open connections that are allowed in the
  # PostgreSQL connection pool (0 = unlimited).
  max_open_connections=0

  # Max idle connections.
  #
  # This sets the max. number of idle connections in the PostgreSQL connection
  # pool (0 = no idle connections are retained).
  max_idle_connections=2
```
In the `dns=` line, modify `<username>`, `<password>`, `<host>`, and `<database>` with your appropriate credentials and targets. If you followed the example above, you would use `chirpstack_as_events` as your username and target database. If your target Postgres database is on the same machine as the Application Server, use `localhost` as your host.
