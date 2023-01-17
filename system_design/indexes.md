- copy of selected columns data
  - searched quick/efficeintly
  - sorted
  - has same number of rows as the table

- filter data
- sort data
- retrieve records
- get smallest info the soonest

- indexes sorted by leftmost column
- if duplicates
  - sort by next column

ex
if index is "owner_id, name, id"
can't resolve index based on a search with name,
not initially sorted on name

Rules
- any columns involved in query should be indexed
  - caveat: unless it is contained in another index
    - redundant index: index contained by another index
  - index prefix: prefix subset of data in column
    - use for longer data type: text, blob, varbinary
    - prefix works: find user names that start with "a"
    - make long enough differiantiate data
- use OR to return records satistying one or more of several conditions
  - unless preventing index from being used and instead doing a full table scan
    - then use a Union
- indexes over all items in query all set
  - except when planner choosing the wrong index
    - FORCE INDEX
    - could use index hint
    - watch out for index changing
- avoid redundant data across tables
  - except when additional joins causing noticeable performance degredation
  - high ratio of reads to writes
    - example checking for spammy users before gettings gists, list of users...
    - denormalize spammy user info so faster lookups
      - requires database changes to add new columns and backfil data
      - data quality: nightly jobs to resolve mismatches between the tables

indexes
- slow down INSERTS, UPDATES, and DELETES
- can use explain on queries to learn more about which indexes are being used by query