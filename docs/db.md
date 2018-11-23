# db

## preparation

```
CREATE user ichat password 'ichat';
create database ichat owner ichat;
```

## account

### users

```
CREATE TABLE users (
	id bigserial PRIMARY KEY,
	email character varying(50) UNIQUE NOT NULL,	password character varying(60) NOT NULL DEFAULT '',
	name character varying(50) NOT NULL
);

```

### user_settings

```

```

### email_confirmations

```

```
