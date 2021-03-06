# go-apt

Collection of frameworks that will used as helper

Features:
- fazzcloud (HTTP client)
- fazzcommon (Helper method)
- fazzdb (Simple ORM for Postgres)
- fazzkv (Key-Value using Redis)

### Fazzdb
- [ ] Documentation and others
  - [x] Godoc
  - [ ] Unit Test
  - [ ] Jenkins Build Status
  - [ ] Test Coverage
- [x] Prepared Statement & Named Query
- [ ] Where Condition
  - [x] Where (AND & OR)
  - [x] Where In
  - [x] Group Where (Bracket)
  - [ ] Subquery
- [ ] Select
  - [x] Select From Table
  - [x] Aggregate
  - [x] Order By
  - [x] Limit
  - [x] Raw Query
  - [x] Group By
  - [x] Having
  - [ ] Join
- [x] Exec Query
  - [x] Bulk Insert
  - [x] Insert
  - [x] Update
  - [x] Delete
  - [x] Raw Query
- [x] ORM
  - [x] Base Model
  - [x] Uuid Model
  - [x] Timestamp (createdAt, updatedAt)
  - [x] Soft Delete (deletedAt)
  - [ ] Recover Soft Delete [NOT TESTED]
- [x] Handle Context
- [ ] Migration
  - [x] Create Table
    - [x] Create Column
    - [x] Primary Key
    - [x] Nullable
    - [x] Unique
    - [x] Foreign Key
  - [ ] Alter Table
    - [x] Add Column
    - [x] Alter Column
    - [x] Drop Column
    - [x] Rename Column
    - [ ] Foreign Key
  - [x] Drop Table
  - [x] Create Enum
  - [ ] Alter Enum
  - [x] Drop Enum
- [ ] Config
