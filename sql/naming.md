# SQL Naming Conventions

Use these conventions across service repos unless there is a strong reason not to:

- schemas are singular service names, for example `athena` and `apollo`
- tables are plural snake_case nouns
- primary keys are `id`
- foreign keys are `<referenced>_id`
- timestamps use `TIMESTAMPTZ`
- flexible metadata uses `JSONB`
- constrained enum-like fields use `CHECK` constraints in early migrations
